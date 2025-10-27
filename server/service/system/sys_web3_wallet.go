package system

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"github.com/mr-tron/base58"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Web3WalletService struct{}

var Web3WalletServiceApp = new(Web3WalletService)

// VerifyPhantomSignature 验证Phantom钱包签名
// Solana使用Ed25519签名算法
func (w *Web3WalletService) VerifyPhantomSignature(walletAddress, message, signature string) (bool, error) {
	// 1. 解码公钥（钱包地址是Base58编码的公钥）
	publicKeyBytes, err := base58.Decode(walletAddress)
	if err != nil {
		global.GVA_LOG.Error("解码钱包地址失败", zap.Error(err))
		return false, fmt.Errorf("无效的钱包地址: %w", err)
	}

	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("钱包地址长度不正确")
	}

	// 2. 解码签名（Base58编码）
	signatureBytes, err := base58.Decode(signature)
	if err != nil {
		global.GVA_LOG.Error("解码签名失败", zap.Error(err))
		return false, fmt.Errorf("无效的签名格式: %w", err)
	}

	if len(signatureBytes) != ed25519.SignatureSize {
		return false, fmt.Errorf("签名长度不正确")
	}

	// 3. 验证签名
	messageBytes := []byte(message)
	publicKey := ed25519.PublicKey(publicKeyBytes)
	valid := ed25519.Verify(publicKey, messageBytes, signatureBytes)

	return valid, nil
}

// GenerateNonce 生成随机Nonce
func (w *Web3WalletService) GenerateNonce() string {
	return uuid.New().String()
}

// GetOrCreateWeb3User 获取或创建Web3用户（幽灵账号）
func (w *Web3WalletService) GetOrCreateWeb3User(walletAddress string) (*system.SysUser, error) {
	var web3Wallet system.SysWeb3Wallet

	// 1. 查找是否已存在该钱包绑定
	err := global.GVA_DB.Where("wallet_address = ?", walletAddress).
		Preload("User").
		Preload("User.Authorities").
		Preload("User.Authority").
		First(&web3Wallet).Error

	now := time.Now()

	if err == nil {
		// 已存在，检查用户状态
		if web3Wallet.User.Enable != 1 {
			global.GVA_LOG.Warn("Web3用户被禁用",
				zap.String("walletAddress", walletAddress),
				zap.Uint("userId", web3Wallet.User.ID),
				zap.Int("enable", web3Wallet.User.Enable))
			
			// 自动启用用户（Web3登录默认启用）
			global.GVA_DB.Model(&system.SysUser{}).
				Where("id = ?", web3Wallet.User.ID).
				Update("enable", 1)
			
			// 重新加载用户信息
			web3Wallet.User.Enable = 1
			global.GVA_LOG.Info("已自动启用Web3用户", zap.Uint("userId", web3Wallet.User.ID))
		}
		
		// 更新最后登录时间
		web3Wallet.LastLoginAt = &now
		global.GVA_DB.Save(&web3Wallet)

		// 设置默认路由
		MenuServiceApp.UserAuthorityDefaultRouter(&web3Wallet.User)
		return &web3Wallet.User, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 2. 不存在，创建新的幽灵账号
	return w.createWeb3User(walletAddress)
}

// createWeb3User 创建Web3幽灵账号
func (w *Web3WalletService) createWeb3User(walletAddress string) (*system.SysUser, error) {
	// 开启事务
	return nil, global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 1. 创建用户
		// 生成唯一的用户名和昵称
		username := fmt.Sprintf("phantom_%s", walletAddress[:8])
		nickname := fmt.Sprintf("Phantom用户_%s", walletAddress[:6])

		// 生成唯一的邀请码
		inviteCode, err := w.generateUniqueInviteCode()
		if err != nil {
			return fmt.Errorf("生成邀请码失败: %w", err)
		}

		user := &system.SysUser{
			UUID:        uuid.New(),
			Username:    username,
			NickName:    nickname,
			Password:    utils.BcryptHash(uuid.New().String()), // 随机密码，Web3用户不使用密码登录
			AuthorityId: 888,                                    // 默认普通用户角色
			Enable:      1,                                      // 启用状态
			InviteCode:  inviteCode,
		}

		if err := tx.Create(user).Error; err != nil {
			return fmt.Errorf("创建用户失败: %w", err)
		}

		// 2. 创建钱包绑定
		now := time.Now()
		web3Wallet := &system.SysWeb3Wallet{
			WalletAddress: walletAddress,
			WalletType:    "phantom",
			UserId:        user.ID,
			LastLoginAt:   &now,
		}

		if err := tx.Create(web3Wallet).Error; err != nil {
			return fmt.Errorf("创建钱包绑定失败: %w", err)
		}

		// 3. 创建用户权限关联
		userAuthority := &system.SysUserAuthority{
			SysUserId:               user.ID,
			SysAuthorityAuthorityId: 888,
		}

		if err := tx.Create(userAuthority).Error; err != nil {
			return fmt.Errorf("创建用户权限失败: %w", err)
		}

		global.GVA_LOG.Info("成功创建Web3幽灵账号",
			zap.String("walletAddress", walletAddress),
			zap.String("username", username),
			zap.Uint("userId", user.ID))

		return nil
	})
}

// generateUniqueInviteCode 生成唯一的邀请码
func (w *Web3WalletService) generateUniqueInviteCode() (string, error) {
	const maxRetries = 10
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const codeLength = 6

	for i := 0; i < maxRetries; i++ {
		// 生成随机邀请码
		code := make([]byte, codeLength)
		randomBytes := uuid.New()
		for j := 0; j < codeLength; j++ {
			code[j] = charset[int(randomBytes[j])%len(charset)]
		}

		inviteCode := string(code)

		// 检查是否已存在
		var count int64
		err := global.GVA_DB.Model(&system.SysUser{}).
			Where("invite_code = ?", inviteCode).
			Count(&count).Error

		if err != nil {
			return "", err
		}

		if count == 0 {
			return inviteCode, nil
		}
	}

	return "", fmt.Errorf("无法生成唯一的邀请码")
}

// VerifySignatureAndLogin 验证签名并登录
func (w *Web3WalletService) VerifySignatureAndLogin(walletAddress, message, signature string) (*system.SysUser, error) {
	// 1. 验证签名
	valid, err := w.VerifyPhantomSignature(walletAddress, message, signature)
	if err != nil {
		return nil, fmt.Errorf("验证签名失败: %w", err)
	}

	if !valid {
		return nil, errors.New("签名验证失败")
	}

	// 2. 获取或创建用户
	user, err := w.GetOrCreateWeb3User(walletAddress)
	if err != nil {
		return nil, fmt.Errorf("获取或创建用户失败: %w", err)
	}

	return user, nil
}

// GetWeb3WalletByAddress 根据钱包地址获取绑定信息
func (w *Web3WalletService) GetWeb3WalletByAddress(walletAddress string) (*system.SysWeb3Wallet, error) {
	var web3Wallet system.SysWeb3Wallet
	err := global.GVA_DB.Where("wallet_address = ?", walletAddress).
		Preload("User").
		First(&web3Wallet).Error

	if err != nil {
		return nil, err
	}

	return &web3Wallet, nil
}

// BindWalletToUser 绑定钱包到现有用户
func (w *Web3WalletService) BindWalletToUser(walletAddress string, userId uint) error {
	// 检查钱包是否已绑定
	var count int64
	err := global.GVA_DB.Model(&system.SysWeb3Wallet{}).
		Where("wallet_address = ?", walletAddress).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("该钱包已被绑定")
	}

	// 创建绑定
	now := time.Now()
	web3Wallet := &system.SysWeb3Wallet{
		WalletAddress: walletAddress,
		WalletType:    "phantom",
		UserId:        userId,
		LastLoginAt:   &now,
	}

	return global.GVA_DB.Create(web3Wallet).Error
}

// UnbindWallet 解绑钱包
func (w *Web3WalletService) UnbindWallet(walletAddress string, userId uint) error {
	return global.GVA_DB.Where("wallet_address = ? AND user_id = ?", walletAddress, userId).
		Delete(&system.SysWeb3Wallet{}).Error
}

