<template>
  <div id="userLayout" class="w-full h-screen relative overflow-hidden bg-black">
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
    
    <!-- 右侧区块链网络动画 -->
    <div class="absolute right-0 top-0 w-1/2 h-full overflow-hidden bg-gradient-to-bl from-slate-900 via-blue-900 to-slate-900">
      <div class="relative w-full h-full">
        <!-- 网络节点动画 -->
        <svg class="absolute inset-0 w-full h-full">
          <!-- 连接线 -->
          <g v-for="i in 15" :key="`line-${i}`">
            <line 
              :x1="Math.random() * 100 + '%'" 
              :y1="Math.random() * 100 + '%'" 
              :x2="Math.random() * 100 + '%'" 
              :y2="Math.random() * 100 + '%'"
              stroke="url(#lineGradient)"
              stroke-width="1"
              stroke-opacity="0.3"
              :style="{
                animation: `pulseLine ${3 + Math.random() * 2}s ease-in-out infinite`,
                animationDelay: Math.random() * 3 + 's'
              }"
            />
          </g>
          
          <!-- 节点 -->
          <g v-for="i in 20" :key="`node-${i}`">
            <circle 
              :cx="Math.random() * 100 + '%'" 
              :cy="Math.random() * 100 + '%'" 
              r="4"
              fill="#3B82F6"
              :style="{
                animation: `nodePulse ${2 + Math.random() * 2}s ease-in-out infinite`,
                animationDelay: Math.random() * 2 + 's'
              }"
            />
            <circle 
              :cx="Math.random() * 100 + '%'" 
              :cy="Math.random() * 100 + '%'" 
              r="6"
              fill="#06B6D4"
              opacity="0.6"
              :style="{
                animation: `nodeGlow ${3 + Math.random() * 2}s ease-in-out infinite`,
                animationDelay: Math.random() * 3 + 's'
              }"
            />
          </g>
          
          <defs>
            <linearGradient id="lineGradient" x1="0%" y1="0%" x2="100%" y2="0%">
              <stop offset="0%" style="stop-color:#3B82F6;stop-opacity:0" />
              <stop offset="50%" style="stop-color:#06B6D4;stop-opacity:0.8" />
              <stop offset="100%" style="stop-color:#3B82F6;stop-opacity:0" />
            </linearGradient>
          </defs>
        </svg>
        
        <!-- 数据流粒子 -->
        <div v-for="i in 30" :key="`particle-${i}`"
             class="absolute w-2 h-2 bg-blue-400 rounded-full"
             :style="{
               top: Math.random() * 100 + '%',
               left: Math.random() * 100 + '%',
               animation: `dataFlow ${5 + Math.random() * 5}s linear infinite`,
               animationDelay: Math.random() * 5 + 's',
               boxShadow: '0 0 10px #3B82F6'
             }">
        </div>
      </div>
    </div>
    
    <!-- 中央登录表单 -->
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
            AiMoney 登录
          </h1>
          <p class="text-white text-opacity-80 text-sm">连接去中心化网络，掌控数字未来</p>
        </div>
        
        <el-form
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          :validate-on-rule-change="false"
          @keyup.enter="submitForm"
          class="space-y-4"
        >
          <el-form-item prop="username" class="mb-4">
            <el-input
              v-model="loginFormData.username"
              size="large"
              placeholder="用户名/手机号/邮箱"
              class="cyber-input"
            >
              <template #prefix>
                <svg class="w-5 h-5 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item prop="password" class="mb-4">
            <el-input
              v-model="loginFormData.password"
              show-password
              size="large"
              type="password"
              placeholder="请输入密码"
              class="cyber-input"
            >
              <template #prefix>
                <svg class="w-5 h-5 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                </svg>
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="mb-6">
            <div class="flex space-x-3">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="请输入验证码"
                size="large"
                class="cyber-input flex-1"
              />
              <div class="w-32 h-12 bg-slate-800 rounded-lg overflow-hidden border border-blue-500 border-opacity-30 cursor-pointer"
                   @click="loginVerify()">
                <img
                  v-if="picPath"
                  class="w-full h-full object-cover"
                  :src="picPath"
                  alt="验证码"
                />
              </div>
            </div>
          </el-form-item>
          
          <el-form-item class="mb-6">
            <el-button
              class="cyber-submit-button !h-14 w-full"
              size="large"
              type="primary"
              @click="submitForm"
            >
              <span class="flex items-center justify-center space-x-3">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"></path>
                </svg>
                <span class="text-lg font-bold">立即登录</span>
              </span>
            </el-button>
          </el-form-item>
          
          <!-- Web3钱包登录分隔线 -->
          <div class="relative my-6">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-blue-500 border-opacity-30"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-4 bg-gradient-to-r from-slate-900 via-slate-800 to-slate-900 text-blue-300">或使用Web3钱包登录</span>
            </div>
          </div>

          <!-- Phantom钱包登录 -->
          <el-form-item class="mb-6">
            <PhantomLogin />
          </el-form-item>
          
          <el-form-item class="mb-3">
            <div class="flex justify-between items-center text-sm">
              <div class="text-white text-opacity-60">
                <span>还没有AiMoney身份？</span>
                <el-link type="primary" @click="toRegister" class="!text-blue-300 !font-medium hover:!text-blue-200 ml-1">
                  立即创建
                </el-link>
                <span style="position: fixed; right: 25px;">右下角找客服支持</span>
              </div>
            </div>
          </el-form-item>
          
          <el-form-item class="mb-6" style="display:none">
            <el-button
              class="cyber-button-secondary !h-12 w-full"
              size="large"
              @click="checkInit"
            >
              前往初始化
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>

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

/* 线条脉冲 */
@keyframes pulseLine {
  0%, 100% { stroke-opacity: 0.1; }
  50% { stroke-opacity: 0.5; }
}

/* 节点脉冲 */
@keyframes nodePulse {
  0%, 100% { r: 4; opacity: 0.6; }
  50% { r: 6; opacity: 1; }
}

/* 节点发光 */
@keyframes nodeGlow {
  0%, 100% { opacity: 0.3; }
  50% { opacity: 0.8; }
}

/* 数据流动 */
@keyframes dataFlow {
  0% { transform: translate(0, 0); opacity: 0; }
  10% { opacity: 1; }
  90% { opacity: 1; }
  100% { transform: translate(500px, 300px); opacity: 0; }
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
  background: rgba(30, 58, 138, 0.2);
  border: 1px solid rgba(59, 130, 246, 0.3);
  color: white;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.cyber-button-secondary:hover {
  background: rgba(30, 58, 138, 0.3);
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
  import { captcha } from '@/api/user'
  import { checkDB } from '@/api/initdb'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import PhantomLogin from '@/components/phantom/PhantomLogin.vue'
  import { reactive, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  
  // 星球数据
  const planets = [
    { size: 12, orbit: 100, speed: 20, color: '#1E40AF', lightColor: '#60A5FA' },
    { size: 16, orbit: 140, speed: 30, color: '#1E3A8A', lightColor: '#3B82F6' },
    { size: 20, orbit: 180, speed: 40, color: '#0891B2', lightColor: '#06B6D4' },
    { size: 14, orbit: 220, speed: 50, color: '#0E7490', lightColor: '#14B8A6' },
    { size: 18, orbit: 260, speed: 60, color: '#6366F1', lightColor: '#818CF8' }
  ]

  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  }
  loginVerify()

  // 登录相关操作
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: '',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })
  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [
      {
        message: '验证码格式不正确',
        trigger: 'blur'
      }
    ]
  })

  const userStore = useUserStore()
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        await loginVerify()
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
  const checkInit = async () => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }

  // 跳转到注册页（公开注册）
  const toRegister = () => {
    // 直接跳转到注册页面
    router.push({ name: 'Register' })
  }

</script> 