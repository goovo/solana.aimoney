<template>
  <div class="phantom-login">
    <!-- Phantom登录按钮 -->
    <el-button
      v-if="!isConnected"
      type="primary"
      :loading="loading"
      @click="connectWallet"
      class="phantom-btn"
    >
      <img src="@/assets/phantom-icon.svg" alt="Phantom" class="phantom-icon" v-if="!loading" />
      {{ loading ? '连接中...' : 'Phantom 钱包登录' }}
    </el-button>

    <!-- 已连接状态 -->
    <div v-else class="connected-status">
      <el-tag type="success" size="large">
        <span class="wallet-address">{{ formatAddress(walletAddress) }}</span>
      </el-tag>
      <el-button size="small" @click="disconnect">断开</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getNonce, phantomLogin } from '@/api/phantom'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'

const router = useRouter()
const userStore = useUserStore()
const routerStore = useRouterStore()

const loading = ref(false)
const isConnected = ref(false)
const walletAddress = ref('')

// 格式化钱包地址
const formatAddress = (address) => {
  if (!address) return ''
  return `${address.slice(0, 4)}...${address.slice(-4)}`
}

// 检测Phantom钱包是否安装
const isPhantomInstalled = () => {
  return window.solana && window.solana.isPhantom
}

// 连接Phantom钱包
const connectWallet = async () => {
  // 检查是否安装Phantom
  if (!isPhantomInstalled()) {
    ElMessage.error('请先安装 Phantom 钱包插件')
    window.open('https://phantom.app/', '_blank')
    return
  }

  loading.value = true

  try {
    // 1. 连接钱包
    console.log('🔵 步骤1: 开始连接Phantom钱包...')
    const response = await window.solana.connect()
    const publicKey = response.publicKey.toString()
    walletAddress.value = publicKey
    console.log('✅ 钱包已连接:', publicKey)

    // 2. 获取Nonce
    console.log('🔵 步骤2: 获取Nonce...')
    const nonceResponse = await getNonce({
      walletAddress: publicKey
    })
    console.log('📦 完整响应:', nonceResponse)

    // axios响应拦截器已经解包了，直接使用data
    const nonceData = nonceResponse.data || nonceResponse
    console.log('📦 Nonce数据:', nonceData)

    if (!nonceData.nonce || !nonceData.message) {
      console.error('❌ Nonce数据格式错误:', nonceData)
      throw new Error('获取Nonce失败')
    }

    const { message, nonce } = nonceData
    console.log('✅ Nonce获取成功:', nonce)

    // 3. 请求签名
    console.log('🔵 步骤3: 请求用户签名...')
    const encodedMessage = new TextEncoder().encode(message)
    const signedMessage = await window.solana.signMessage(encodedMessage, 'utf8')
    console.log('✅ 用户已签名')
    
    // 将签名转换为Base58字符串
    const signature = encodeBase58(signedMessage.signature)
    console.log('📝 签名(Base58):', signature.substring(0, 20) + '...')

    // 4. 调用登录API
    console.log('🔵 步骤4: 调用登录API...')
    const loginResponse = await phantomLogin({
      walletAddress: publicKey,
      signature: signature,
      message: message
    })
    console.log('📦 完整登录响应:', loginResponse)

    // axios响应拦截器已经解包了，直接使用data
    const loginData = loginResponse.data || loginResponse
    console.log('📦 登录数据:', loginData)

    if (!loginData.user || !loginData.token) {
      console.error('❌ 登录数据格式错误:', loginData)
      throw new Error('登录失败')
    }

    // 5. 保存登录信息
    const { user, token, expiresAt } = loginData
    console.log('🔵 步骤5: 保存登录信息并初始化路由...')
    
    // 设置token
    userStore.setToken(token)
    console.log('✅ Token已保存')
    
    // 获取完整的用户信息（包括权限和路由信息）
    await userStore.GetUserInfo()
    console.log('✅ 用户信息已获取:', userStore.userInfo)
    
    // 初始化路由
    await routerStore.SetAsyncRouter()
    const asyncRouters = routerStore.asyncRouters
    console.log('✅ 路由已初始化，共', asyncRouters.length, '个路由')
    
    // 注册到路由表
    asyncRouters.forEach((asyncRouter) => {
      router.addRoute(asyncRouter)
    })
    
    isConnected.value = true
    ElMessage.success('登录成功！')
    console.log('✅ 登录完成，准备跳转...')

    // 6. 跳转到首页
    const defaultRouter = userStore.userInfo.authority?.defaultRouter || 'dashboard'
    console.log('🔵 跳转到:', defaultRouter)
    
    // 设置操作系统类型
    const isWindows = /windows/i.test(navigator.userAgent)
    window.localStorage.setItem('osType', isWindows ? 'WIN' : 'MAC')
    
    setTimeout(() => {
      router.push({ name: defaultRouter })
    }, 500)

  } catch (error) {
    console.error('❌ 错误详情:', error)
    console.error('错误堆栈:', error.stack)
    
    // 更详细的错误提示
    let errorMsg = error.message || '连接钱包失败，请重试'
    if (error.response) {
      console.error('API错误响应:', error.response)
      errorMsg = error.response.data?.msg || errorMsg
    }
    
    ElMessage.error(errorMsg)
    disconnect()
  } finally {
    loading.value = false
  }
}

// 断开连接
const disconnect = async () => {
  try {
    if (window.solana) {
      await window.solana.disconnect()
    }
    walletAddress.value = ''
    isConnected.value = false
  } catch (error) {
    console.error('断开连接失败:', error)
  }
}

// Base58编码（简化版，用于Solana签名）
const encodeBase58 = (buffer) => {
  const ALPHABET = '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
  const base = ALPHABET.length

  if (buffer.length === 0) return ''

  const digits = [0]
  for (let i = 0; i < buffer.length; i++) {
    let carry = buffer[i]
    for (let j = 0; j < digits.length; j++) {
      carry += digits[j] << 8
      digits[j] = carry % base
      carry = (carry / base) | 0
    }
    while (carry > 0) {
      digits.push(carry % base)
      carry = (carry / base) | 0
    }
  }

  let string = ''
  for (let i = 0; buffer[i] === 0 && i < buffer.length - 1; i++) {
    string += ALPHABET[0]
  }
  for (let i = digits.length - 1; i >= 0; i--) {
    string += ALPHABET[digits[i]]
  }

  return string
}

// 监听钱包状态变化
onMounted(() => {
  if (isPhantomInstalled()) {
    window.solana.on('connect', () => {
      console.log('Phantom钱包已连接')
    })

    window.solana.on('disconnect', () => {
      console.log('Phantom钱包已断开')
      disconnect()
    })
  }
})
</script>

<style scoped>
.phantom-login {
  width: 100%;
}

.phantom-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: linear-gradient(135deg, #AB47BC 0%, #7B1FA2 100%);
  border: none;
  transition: all 0.3s;
}

.phantom-btn:hover {
  background: linear-gradient(135deg, #9C27B0 0%, #6A1B9A 100%);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(171, 71, 188, 0.4);
}

.phantom-icon {
  width: 24px;
  height: 24px;
}

.connected-status {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.wallet-address {
  font-family: monospace;
  font-size: 14px;
}
</style>

