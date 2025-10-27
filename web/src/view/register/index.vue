<template>
  <div id="userRegister" class="w-full h-screen relative overflow-hidden bg-black">
    <!-- 左侧星空银河系动画 -->
    <div class="absolute left-0 top-0 w-1/2 h-full overflow-hidden">
      <div class="relative w-full h-full">
        <!-- 星空背景 -->
        <div class="absolute inset-0 bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900">
          <!-- 星星 -->
          <div v-for="i in 200" :key="`star-${i}`"
               class="absolute rounded-full bg-white"
               :style="{
                 width: Math.random() * 3 + 'px',
                 height: Math.random() * 3 + 'px',
                 top: Math.random() * 100 + '%',
                 left: Math.random() * 100 + '%',
                 opacity: Math.random() * 0.8 + 0.2,
                 animation: `twinkle ${Math.random() * 5 + 3}s infinite`
               }">
          </div>
        </div>
        
        <!-- 银河系中心 -->
        <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <div class="relative w-80 h-80">
            <!-- 中心黑洞 -->
            <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-20 h-20 bg-black rounded-full shadow-2xl">
              <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-cyan-500 rounded-full animate-pulse blur-xl"></div>
            </div>
            
            <!-- 旋转的星球轨道 -->
            <div v-for="(planet, index) in planets" :key="index"
                 class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-full border border-white border-opacity-20"
                 :style="{
                   width: planet.orbit + 'px',
                   height: planet.orbit + 'px',
                   animation: `rotate ${planet.speed}s linear infinite`
                 }">
              <!-- 星球 -->
              <div class="absolute top-0 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-full shadow-lg"
                   :style="{
                     width: planet.size + 'px',
                     height: planet.size + 'px',
                     background: planet.color,
                     boxShadow: `0 0 ${planet.size/2}px ${planet.color}`
                   }">
                <!-- 星球纹理 -->
                <div class="absolute inset-0 rounded-full opacity-50"
                     :style="{
                       background: `radial-gradient(circle at 30% 30%, ${planet.lightColor}, transparent)`
                     }">
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 流星 -->
        <div v-for="i in 5" :key="`meteor-${i}`"
             class="absolute h-0.5 bg-gradient-to-r from-transparent via-white to-transparent"
             :style="{
               top: Math.random() * 100 + '%',
               left: '-100px',
               width: '100px',
               animation: `meteor ${Math.random() * 10 + 10}s linear infinite`,
               animationDelay: Math.random() * 10 + 's'
             }">
        </div>
      </div>
    </div>
    
    <!-- 右侧Web3 Token Logo动画 -->
    <div class="absolute right-0 top-0 w-1/2 h-full overflow-hidden bg-gradient-to-bl from-slate-900 via-blue-900 to-slate-900">
      <div class="relative w-full h-full">
        <!-- Token logos grid -->
        <div class="grid grid-cols-4 gap-8 p-8 absolute inset-0">
          <div v-for="(token, index) in tokens" :key="index"
               class="flex items-center justify-center transform transition-all duration-1000 hover:scale-125"
               :style="{
                 animation: `float ${3 + Math.random() * 2}s ease-in-out infinite`,
                 animationDelay: index * 0.2 + 's'
               }">
            <div class="relative w-16 h-16 rounded-xl flex items-center justify-center bg-black bg-opacity-50 backdrop-blur-sm border border-white border-opacity-20">
              <!-- Token icon -->
              <div class="text-2xl font-bold" :style="{ color: token.color }">
                {{ token.symbol }}
              </div>
              <!-- 发光效果 -->
              <div class="absolute inset-0 rounded-xl opacity-30"
                   :style="{
                     background: `radial-gradient(circle, ${token.color}, transparent)`,
                     animation: `pulse ${2 + Math.random() * 2}s ease-in-out infinite`
                   }">
              </div>
              <!-- 闪光效果 -->
              <div class="absolute inset-0 rounded-xl overflow-hidden">
                <div class="absolute -top-full left-1/2 transform -translate-x-1/2 w-1 h-full bg-gradient-to-b from-transparent via-white to-transparent opacity-60"
                     :style="{
                       animation: `shine ${3 + Math.random() * 2}s linear infinite`,
                       animationDelay: index * 0.3 + 's'
                     }">
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 连接线动画 -->
        <svg class="absolute inset-0 w-full h-full pointer-events-none">
          <defs>
            <linearGradient id="lineGradient" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" style="stop-color:#3B82F6;stop-opacity:0" />
              <stop offset="50%" style="stop-color:#3B82F6;stop-opacity:0.5" />
              <stop offset="100%" style="stop-color:#06B6D4;stop-opacity:0" />
            </linearGradient>
          </defs>
          <line v-for="i in 20" :key="`line-${i}`"
                x1="0" y1="0" x2="100%" y2="100%"
                stroke="url(#lineGradient)"
                stroke-width="0.5"
                :style="{
                  animation: `moveLine ${5 + Math.random() * 5}s linear infinite`,
                  animationDelay: Math.random() * 5 + 's'
                }" />
        </svg>
      </div>
    </div>
    
    <!-- 中央注册表单 -->
    <div class="absolute inset-0 flex items-center justify-center z-20">
      <div class="backdrop-blur-2xl bg-black bg-opacity-40 rounded-3xl shadow-2xl border border-white border-opacity-20 p-8 w-full max-w-lg transform hover:scale-[1.02] transition-all duration-500">
        <div class="text-center mb-8">
          <div class="flex justify-center mb-6">
            <div class="relative"> 
              <div class="absolute -bottom-2 -right-2 w-8 h-8 bg-gradient-to-r from-blue-400 to-cyan-400 rounded-full border-2 border-white animate-pulse shadow-lg"></div>
              <!-- 光环效果 -->
              <div class="absolute inset-0 rounded-full border-4 border-blue-400 border-opacity-30 animate-ping"></div>
            </div>
          </div>
          <!-- 酷炫的标题 -->
          <h1 class="text-6xl font-bold mb-4 bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-blue-600 animate-gradient-x">
            Web3 搞钱
          </h1>
          <p class="text-white text-opacity-80 text-sm">进入去中心化世界，掌控数字资产</p>
        </div>
        
        <el-form ref="registerForm" :model="form" :rules="rules" :validate-on-rule-change="false" @keyup.enter="submit" class="space-y-4">
          <!-- 验证方式选择器 -->
          <el-form-item class="mb-4">
            <div class="w-full"> 
              <el-radio-group v-model="registerType" @change="onRegisterTypeChange" class="w-full">
                <div class="grid grid-cols-2 gap-3">
                  <el-radio-button label="phone" class="flex-1 !h-12">
                    <div class="flex items-center justify-center space-x-2">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                      </svg>
                      <span>用手机注册</span>
                    </div>
                  </el-radio-button>
                  <el-radio-button label="email" class="flex-1 !h-12">
                    <div class="flex items-center justify-center space-x-2">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                      </svg>
                      <span>用邮箱注册</span>
                    </div>
                  </el-radio-button>
                </div>
              </el-radio-group>
            </div>
          </el-form-item>
          
          <!-- 手机号注册 -->
          <template v-if="registerType === 'phone'">
            <el-form-item prop="phone" class="mb-4">
              <el-input 
                v-model="form.phone" 
                size="large" 
                placeholder="请输入手机号码"
                @input="onPhoneChange"
                @blur="checkPhoneAvailability"
                class="cyber-input"
              />
              <!-- 手机号状态显示 -->
              <div v-if="phoneStatus" class="mt-1 text-sm flex items-center transition-all duration-300" :class="phoneStatus.type === 'success' ? 'text-green-400' : 'text-red-400'">
                <span v-if="phoneStatus.type === 'success'" class="mr-1 animate-pulse"></span>
                <span v-else class="mr-1">⚠</span>
                {{ phoneStatus.message }}
              </div>
            </el-form-item>
            <!-- 手机验证码 -->
            <el-form-item prop="phoneCode" class="mb-4">
              <div class="flex space-x-3">
                <el-input v-model="form.phoneCode" placeholder="请输入验证码" size="large" class="cyber-input flex-1" />
                <el-button 
                  type="primary" 
                  size="large" 
                  class="cyber-button !flex-1 max-w-[120px]"
                  :disabled="phoneCodeCountdown > 0 || !canSendPhone" 
                  @click="onSendPhoneCode"
                >
                  {{ phoneCodeCountdown > 0 ? `${phoneCodeCountdown}s` : '获取验证码' }}
                </el-button>
              </div>
            </el-form-item>
          </template>
          
          <!-- 邮箱注册 -->
          <template v-if="registerType === 'email'">
            <el-form-item prop="email" class="mb-4">
              <el-input 
                v-model="form.email" 
                size="large" 
                placeholder="请输入邮箱地址"
                @input="onEmailChange"
                @blur="checkEmailAvailability"
                class="cyber-input"
              />
              <!-- 邮箱状态显示 -->
              <div v-if="emailStatus" class="mt-1 text-sm flex items-center transition-all duration-300" :class="emailStatus.type === 'success' ? 'text-green-400' : 'text-red-400'">
                <span v-if="emailStatus.type === 'success'" class="mr-1 animate-pulse"></span>
                <span v-else class="mr-1">⚠</span>
                {{ emailStatus.message }}
              </div>
            </el-form-item>
            <!-- 邮箱验证码 -->
            <el-form-item prop="emailCode" class="mb-4">
              <div class="flex space-x-3">
                <el-input v-model="form.emailCode" placeholder="请输入验证码" size="large" class="cyber-input flex-1" />
                <el-button 
                  type="primary" 
                  size="large" 
                  class="cyber-button !flex-1 max-w-[120px]"
                  :disabled="emailCodeCountdown > 0 || !canSendEmail" 
                  @click="onSendEmailCode"
                >
                  {{ emailCodeCountdown > 0 ? `${emailCodeCountdown}s` : '获取验证码' }}
                </el-button>
              </div>
            </el-form-item>
          </template>
          
          <!-- 密码 -->
          <el-form-item prop="password" class="mb-4">
            <el-input v-model="form.password" show-password size="large" type="password" placeholder="设置密码（至少6位）"
                     class="cyber-input" />
          </el-form-item>
          
          <!-- 确认密码 -->
          <el-form-item prop="confirmPassword" class="mb-4">
            <el-input v-model="form.confirmPassword" show-password size="large" type="password" placeholder="确认密码"
                     class="cyber-input" />
          </el-form-item>
          
          <!-- 推荐码（可选） -->
          <el-form-item prop="referrerCode" class="mb-6" style="display: none;">
            <el-input v-model="form.referrerCode" size="large" placeholder="推荐码（可选）"
                     class="cyber-input" />
          </el-form-item>
          
          <!-- 注册按钮 -->
          <el-form-item class="mb-6">
            <el-button class="cyber-submit-button !h-14 w-full" 
                       size="large" 
                       :loading="submitting" 
                       @click="submit">
              <span class="flex items-center justify-center space-x-3">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
                <span class="text-lg font-bold">立即注册</span>
              </span>
            </el-button>
          </el-form-item>
          
          <!-- 去登录入口 -->
          <div class="text-center">
            <span class="text-white text-opacity-60 text-sm">已有Web3身份？</span>
            <el-link type="primary" @click="toLogin" class="!text-blue-300 !font-medium hover:!text-blue-200 ml-1">
              立即连接
            </el-link>
            <span class="text-white text-opacity-60 text-sm" style="position: fixed; right: 30px; ">右下角找客服支持</span>
          </div>
        </el-form>
      </div>
    </div>

    <!-- 图形验证码弹窗 -->
    <el-dialog v-model="captchaDialogVisible" title="安全验证" width="400px" @close="onCaptchaDialogClose" custom-class="cyber-dialog">
      <div class="flex flex-col items-center">
        <div class="mb-4">
          <img v-if="captchaImage" :src="captchaImage" alt="验证码" class="cursor-pointer border rounded-lg" @click="getCaptcha" />
        </div>
        <el-input v-model="captchaInput" placeholder="请输入验证码" class="cyber-input" @keyup.enter="onCaptchaConfirm" />
        <div class="flex justify-between w-full">
          <el-button @click="captchaDialogVisible = false" class="cyber-button-secondary">取消</el-button>
          <el-button type="primary" @click="onCaptchaConfirm" class="cyber-button">确认</el-button>
        </div>
      </div>
    </el-dialog>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto w-full z-30">
      <div class="links items-center justify-center gap-3 hidden md:flex"> 
      </div>
    </BottomInfo>
  </div>
</template>

<style scoped>
/* 星星闪烁动画 */
@keyframes twinkle {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 1; }
}

/* 轨道旋转 */
@keyframes rotate {
  from { transform: translate(-50%, -50%) rotate(0deg); }
  to { transform: translate(-50%, -50%) rotate(360deg); }
}

/* 流星动画 */
@keyframes meteor {
  0% { transform: translateX(0) translateY(0) rotate(-45deg); opacity: 0; }
  10% { opacity: 1; }
  100% { transform: translateX(1000px) translateY(500px) rotate(-45deg); opacity: 0; }
}

/* 浮动动画 */
@keyframes float {
  0%, 100% { transform: translateY(0) rotate(0deg); }
  33% { transform: translateY(-10px) rotate(120deg); }
  66% { transform: translateY(10px) rotate(240deg); }
}

/* 脉冲发光 */
@keyframes pulse {
  0%, 100% { opacity: 0.3; transform: scale(1); }
  50% { opacity: 0.6; transform: scale(1.1); }
}

/* 闪光效果 */
@keyframes shine {
  0% { transform: translateX(-100%) translateY(-100%); }
  100% { transform: translateX(200%) translateY(200%); }
}

/* 线条移动 */
@keyframes moveLine {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

/* 渐变动画 */
@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.animate-gradient-x {
  background-size: 200% 200%;
  animation: gradient 3s ease infinite;
}

/* 赛博朋克输入框 */
.cyber-input :deep(.el-input__wrapper) {
  background: rgba(30, 58, 138, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.1);
  backdrop-filter: blur(10px);
}

.cyber-input :deep(.el-input__inner) {
  color: white;
  font-family: 'Courier New', monospace;
}

.cyber-input :deep(.el-input__inner)::placeholder {
  color: rgba(147, 197, 253, 0.5);
}

.cyber-input :deep(.el-input__wrapper:hover) {
  border-color: rgba(59, 130, 246, 0.5);
  box-shadow: 0 0 30px rgba(59, 130, 246, 0.2);
}

.cyber-input :deep(.el-input__wrapper.is-focus) {
  border-color: rgba(59, 130, 246, 0.8);
  box-shadow: 0 0 40px rgba(59, 130, 246, 0.4);
}

/* 赛博朋克按钮 */
.cyber-button {
  background: linear-gradient(135deg, #3B82F6 0%, #1E40AF 100%);
  border: none;
  color: white;
  font-weight: bold;
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.4);
  transition: all 0.3s ease;
}

.cyber-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(59, 130, 246, 0.6);
}

.cyber-button-secondary {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: white;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.cyber-button-secondary:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
}

/* 提交按钮 */
.cyber-submit-button {
  background: linear-gradient(135deg, #3B82F6 0%, #06B6D4 50%, #1E40AF 100%);
  border: none;
  color: white;
  font-weight: bold;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.5);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.cyber-submit-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.5s ease;
}

.cyber-submit-button:hover::before {
  left: 100%;
}

.cyber-submit-button:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.7);
}

/* 单选按钮组样式 */
.cyber-input :deep(.el-radio-button__inner) {
  background: rgba(30, 58, 138, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
  color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(10px);
}

.cyber-input :deep(.el-radio-button__orig-radio:checked + .el-radio-button__inner) {
  background: linear-gradient(135deg, #3B82F6 0%, #1E40AF 100%);
  border-color: transparent;
  color: white;
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.5);
}

/* 对话框样式 */
.cyber-dialog {
  background: rgba(15, 23, 42, 0.95);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(59, 130, 246, 0.3);
  border-radius: 1rem;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

.cyber-dialog .el-dialog__title {
  color: white;
  font-weight: bold;
}

/* 链接悬停效果 */
.el-link:hover {
  text-shadow: 0 0 20px rgba(59, 130, 246, 0.8);
}

/* 响应式优化 */
@media (max-width: 1024px) {
  .w-1\/2 {
    width: 100%;
  }
  .max-w-lg {
    margin: 1rem;
    width: calc(100% - 2rem);
  }
}
</style>

<script setup>
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { reactive, ref, onMounted } from 'vue'
  import { useRouter, useRoute } from 'vue-router'
  import { ElMessage } from 'element-plus'
  import { publicRegister, sendEmailCode, sendPhoneCode, captcha, checkPhone, checkEmail } from '@/api/user'

  defineOptions({ name: 'Register' })

  const router = useRouter()
  const route = useRoute()
  const registerForm = ref(null)
  const submitting = ref(false)
  
  // 注册方式：默认手机注册
  const registerType = ref('phone')
  
  // 验证码倒计时
  const emailCodeCountdown = ref(0)
  const phoneCodeCountdown = ref(0)
  const canSendEmail = ref(true)
  const canSendPhone = ref(true)
  
  // 手机号和邮箱检查状态
  const phoneStatus = ref(null)
  const emailStatus = ref(null)
  const phoneCheckTimer = ref(null)
  const emailCheckTimer = ref(null)
    
  // 图形验证码相关
  const captchaDialogVisible = ref(false)
  const captchaImage = ref('')
  const captchaInput = ref('')
  const captchaId = ref('')
  const pendingAction = ref(null) // 存储待执行的动作：'sendEmail' 或 'sendPhone'

  // authorityId/authorityIds 设置为默认普通用户 888
  const DEFAULT_ROLE = 888

  const form = reactive({
    email: '',
    emailCode: '',
    phone: '',
    phoneCode: '',
    password: '',
    confirmPassword: '',
    authorityId: DEFAULT_ROLE,
    authorityIds: [DEFAULT_ROLE],
    referrerCode: '' // 推荐码
  })

  // 星球数据
  const planets = [
    { size: 12, orbit: 100, speed: 20, color: '#1E40AF', lightColor: '#60A5FA' },
    { size: 16, orbit: 140, speed: 30, color: '#1E3A8A', lightColor: '#3B82F6' },
    { size: 20, orbit: 180, speed: 40, color: '#0891B2', lightColor: '#06B6D4' },
    { size: 14, orbit: 220, speed: 50, color: '#0E7490', lightColor: '#14B8A6' },
    { size: 18, orbit: 260, speed: 60, color: '#6366F1', lightColor: '#818CF8' }
  ]

  // Token数据
  const tokens = [
    { symbol: '₿', color: '#F7931A', name: 'Bitcoin' },
    { symbol: 'Ξ', color: '#627EEA', name: 'Ethereum' },
    { symbol: 'BNB', color: '#F3BA2F', name: 'Binance' },
    { symbol: '♦', color: '#00D4AA', name: 'Cardano' },
    { symbol: '▲', color: '#2617FB', name: 'Solana' },
    { symbol: '◉', color: '#14F195', name: 'Polkadot' },
    { symbol: '⬡', color: '#000000', name: 'Uniswap' },
    { symbol: '⧗', color: '#1A1A1A', name: 'Chainlink' },
    { symbol: '◈', color: '#FF2D55', name: 'Polygon' },
    { symbol: '⬢', color: '#0033AD', name: 'Avalanche' },
    { symbol: '◐', color: '#008C73', name: 'Cosmos' },
    { symbol: '⬟', color: '#002D74', name: 'Algorand' }
  ]

  // 动态验证规则
  const rules = reactive({
    email: [
      { 
        required: () => registerType.value === 'email', 
        message: '请输入邮箱', 
        trigger: 'blur' 
      },
      { 
        type: 'email', 
        message: '邮箱格式不正确', 
        trigger: 'blur',
        required: () => registerType.value === 'email'
      }
    ],
    emailCode: [
      { 
        required: () => registerType.value === 'email', 
        message: '请输入邮箱验证码', 
        trigger: 'blur' 
      }
    ],
    phone: [
      { 
        required: () => registerType.value === 'phone', 
        message: '请输入手机号', 
        trigger: 'blur' 
      },
      { 
        pattern: /^1[3-9]\d{9}$/, 
        message: '手机号格式不正确', 
        trigger: 'blur',
        required: () => registerType.value === 'phone'
      }
    ],
    phoneCode: [
      { 
        required: () => registerType.value === 'phone', 
        message: '请输入手机验证码', 
        trigger: 'blur' 
      }
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, message: '至少6个字符', trigger: 'blur' }
    ],
    confirmPassword: [
      { required: true, message: '请再次输入密码', trigger: 'blur' },
      {
        validator: (rule, value, callback) => {
          if (value !== form.password) {
            callback(new Error('两次输入的密码不一致'))
          } else {
            callback()
          }
        },
        trigger: 'blur'
      }
    ],
    referrerCode: [
      { 
        required: false, 
        message: '推荐码格式不正确', 
        trigger: 'blur',
        pattern: /^[A-Za-z0-9]{6}$/
      }
    ]
  })

  // 监听注册方式变化，清空相关字段
  const onRegisterTypeChange = () => {
    if (registerType.value === 'phone') {
      form.email = ''
      form.emailCode = ''
      emailStatus.value = null
    } else {
      form.phone = ''
      form.phoneCode = ''
      phoneStatus.value = null
    }
  }

  // 手机号输入变化处理（防抖）
  const onPhoneChange = () => {
    phoneStatus.value = null
    if (phoneCheckTimer.value) {
      clearTimeout(phoneCheckTimer.value)
    }
    phoneCheckTimer.value = setTimeout(() => {
      if (form.phone && form.phone.length === 11) {
        checkPhoneAvailability()
      }
    }, 500) // 500ms 防抖延迟
  }

  // 邮箱输入变化处理（防抖）
  const onEmailChange = () => {
    emailStatus.value = null
    if (emailCheckTimer.value) {
      clearTimeout(emailCheckTimer.value)
    }
    emailCheckTimer.value = setTimeout(() => {
      if (form.email && form.email.includes('@')) {
        checkEmailAvailability()
      }
    }, 500) // 500ms 防抖延迟
  }

  // 检查手机号可用性
  const checkPhoneAvailability = async () => {
    if (!form.phone || form.phone.length !== 11) {
      phoneStatus.value = null
      return
    }
    
    try {
      const res = await checkPhone({ phone: form.phone })
      if (res.code === 0) {
        const isAvailable = res.data.available
        phoneStatus.value = {
          type: isAvailable ? 'success' : 'error',
          message: isAvailable ? '✓ 手机号可用' : '× 手机号已被注册'
        }
      }
    } catch (e) {
      console.error('检查手机号失败:', e)
      phoneStatus.value = {
        type: 'error',
        message: '× 检查手机号失败'
      }
    }
  }

  // 检查邮箱可用性
  const checkEmailAvailability = async () => {
    if (!form.email || !form.email.includes('@')) {
      emailStatus.value = null
      return
    }
    
    try {
      const res = await checkEmail({ email: form.email })
      if (res.code === 0) {
        const isAvailable = res.data.available
        emailStatus.value = {
          type: isAvailable ? 'success' : 'error',
          message: isAvailable ? '✓ 邮箱可用' : '× 邮箱已被注册'
        }
      }
    } catch (e) {
      console.error('检查邮箱失败:', e)
      emailStatus.value = {
        type: 'error',
        message: '× 检查邮箱失败'
      }
    }
  }

  // 获取图形验证码
  const getCaptcha = async () => {
    try {
      const res = await captcha()
      if (res.code === 0) {
        captchaImage.value = res.data.picPath
        captchaId.value = res.data.captchaId
      }
    } catch (e) {
      ElMessage.error('获取验证码失败')
    }
  }

  // 显示图形验证码弹窗
  const showCaptchaDialog = (action) => {
    pendingAction.value = action
    captchaInput.value = ''
    captchaDialogVisible.value = true
    getCaptcha()
  }

  // 图形验证码弹窗关闭
  const onCaptchaDialogClose = () => {
    captchaInput.value = ''
    pendingAction.value = null
  }

  // 图形验证码确认
  const onCaptchaConfirm = () => {
    if (!captchaInput.value) {
      ElMessage.warning('请输入图形验证码')
      return
    }
    
    captchaDialogVisible.value = false
    
    // 执行挂起的操作
    if (pendingAction.value === 'sendEmail') {
      sendEmailCodeWithCaptcha()
    } else if (pendingAction.value === 'sendPhone') {
      sendPhoneCodeWithCaptcha()
    }
  }

  // 发送邮箱验证码
  const onSendEmailCode = () => {
    if (!form.email) {
      ElMessage.warning('请先填写邮箱地址')
      return
    }
    if (!emailStatus.value || emailStatus.value.type !== 'success') {
      ElMessage.warning('请确保邮箱可用')
      return
    }
    showCaptchaDialog('sendEmail')
  }

  // 带图形验证码发送邮箱验证码
  const sendEmailCodeWithCaptcha = async () => {
    try {
      canSendEmail.value = false
      const res = await sendEmailCode({ 
        email: form.email,
        captchaId: captchaId.value,
        captcha: captchaInput.value
      })
      if (res.code === 0) {
        ElMessage.success('验证码已发送，请查收邮箱')
        startCountdown('email')
      }
    } catch (e) {
      ElMessage.error(e?.response?.data?.msg || e?.msg || '发送验证码失败')
      // 重新获取图形验证码
      getCaptcha()
    } finally {
      canSendEmail.value = true
    }
  }

  // 发送手机验证码
  const onSendPhoneCode = () => {
    if (!form.phone) {
      ElMessage.warning('请先填写手机号码')
      return
    }
    if (!phoneStatus.value || phoneStatus.value.type !== 'success') {
      ElMessage.warning('请确保手机号可用')
      return
    }
    showCaptchaDialog('sendPhone')
  }

  // 带图形验证码发送手机验证码
  const sendPhoneCodeWithCaptcha = async () => {
    try {
      canSendPhone.value = false
      const res = await sendPhoneCode({ 
        phone: form.phone,
        captchaId: captchaId.value,
        captcha: captchaInput.value
      })
      if (res.code === 0) {
        ElMessage.success('验证码已发送，请查收短信')
        startCountdown('phone')
      }
    } catch (e) {
      ElMessage.error(e?.response?.data?.msg || e?.msg || '发送验证码失败')
      // 重新获取图形验证码
      getCaptcha()
    } finally {
      canSendPhone.value = true
    }
  }

  // 启动倒计时
  const startCountdown = (type) => {
    const countdownRef = type === 'email' ? emailCodeCountdown : phoneCodeCountdown
    countdownRef.value = 60
    const timer = setInterval(() => {
      countdownRef.value--
      if (countdownRef.value <= 0) {
        clearInterval(timer)
        countdownRef.value = 0
      }
    }, 1000)
  }

  // 提交注册
  const submit = async () => {
    registerForm.value.validate(async (valid) => {
      if (!valid) {
        ElMessage.error('请正确填写注册信息')
        return
      }
      
      // 检查手机号/邮箱可用性
      if (registerType.value === 'phone' && (!phoneStatus.value || phoneStatus.value.type !== 'success')) {
        ElMessage.error('请确保手机号可用')
        return
      }
      if (registerType.value === 'email' && (!emailStatus.value || emailStatus.value.type !== 'success')) {
        ElMessage.error('请确保邮箱可用')
        return
      }
      
      submitting.value = true
      try {
        const payload = {
          password: form.password,
          authorityId: form.authorityId,
          authorityIds: form.authorityIds,
          referrerCode: form.referrerCode // 添加推荐码
        }
        
        // 根据注册方式添加对应字段
        if (registerType.value === 'phone') {
          payload.phone = form.phone
          payload.phoneCode = form.phoneCode
        } else {
          payload.email = form.email
          payload.emailCode = form.emailCode
        }
        
        const res = await publicRegister(payload)
        if (res.code === 0) {
          ElMessage.success('注册成功，欢迎进入Web3世界！')
          router.push({ name: 'Login' })
        }
      } catch (e) {
        ElMessage.error(e?.response?.data?.msg || e?.msg || '注册失败')
      } finally {
        submitting.value = false
      }
    })
  }

  // 跳转登录
  const toLogin = () => router.push({ name: 'Login' })

  // 页面初始化 - 获取URL参数中的推荐码
  onMounted(() => {
    // 从URL参数获取推荐码
    const refCode = route.query.ref
    if (refCode && typeof refCode === 'string') {
      form.referrerCode = refCode
    }
  })

 
</script>