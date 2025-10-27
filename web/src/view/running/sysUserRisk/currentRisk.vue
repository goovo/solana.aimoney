<template>
  <div class="risk-assessment-container">

    <!-- 风险等级信息展示 -->
    <div class="risk-info-section">
      <div class="section-header">
        <h2>用户风险等级评估</h2>
        <p>基于您的投资偏好和风险承受能力，系统为您提供专业的风险评估服务</p> 
        <el-button type="primary" @click="handleReassess">重新评估</el-button>
      </div>

      <div v-if="hasRiskAssessment" class="risk-result">
        <el-card class="risk-card">
          <div class="risk-content">
            <div class="risk-level">
              <h3 :class="`risk-${currentRisk.risk}`">
                {{ riskLevelText[currentRisk.risk] }}
              </h3>
              <p class="assessment-time">评估时间：{{ formatDate(currentRisk.created_at) }}</p>
            </div>
            <div class="risk-details">
              <p>用户ID：{{ currentRisk.userId }}</p>
              <p>风险偏好：{{ riskDescription[currentRisk.risk] }}</p>
              <p>建议投资策略：{{ investmentStrategy[currentRisk.risk] }}</p>
            </div>
          </div>
        </el-card>
      </div>

      
    </div>

    
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router' // 引入路由用于页面跳转
import { getRisk } from "@/api/user"
import { useUserStore } from '@/pinia/modules/user' // 引入用户状态管理

// 响应式数据
const currentRisk = ref(null)
const assessmentForm = ref({
  experience: '',
  tolerance: '',
  period: ''
})

// 使用路由实例和用户状态
const router = useRouter() // 使用 vue-router 实例
const userStore = useUserStore() // 获取用户状态

// 点击"重新评估"按钮时，跳转到风险评估页面
// 说明：在项目使用的哈希路由模式下，最终地址会呈现为 /#/layout/routerHolder2/runRisk
// 这里通过 router.push 传入不含 # 的路径，保持与 SPA 内部路由一致
const handleReassess = () => {
  router.push('/layout/routerHolder2/runRisk') // 跳转到评估页
}

// 获取用户风险等级数据
const fetchUserRisk = async () => {
  try {
    // 获取当前用户ID
    const userId = userStore.userInfo.ID
    if (!userId) {
      console.warn('用户ID不存在，跳转到评估页面')
      router.push('/layout/routerHolder2/runRisk')
      return
    }

    // 调用API获取用户风险等级，传递userId参数
    const userRisk = await getRisk()
    
    // 检查返回数据是否为空或失败
    if (!userRisk || userRisk.code !== 0 || !userRisk.data) {
      console.warn('未查询到风险等级数据，跳转到评估页面')
      router.push('/layout/routerHolder2/runRisk')
      return
    }

    // 成功获取数据，更新 currentRisk 对象
    const { risk, updatedAt, userId: responseUserId } = userRisk.data
    currentRisk.value = {
      risk: risk,                    // 风险等级: low/medium/high
      created_at: updatedAt,         // 评估时间
      userId: responseUserId         // 用户ID
    }

    console.log('风险等级数据获取成功:', currentRisk.value)
    
  } catch (error) {
    console.error('获取风险等级失败:', error)
    ElMessage.warning('获取风险等级失败，请重新进行评估')
    // 发生错误时跳转到评估页面
    router.push('/layout/routerHolder2/runRisk')
  }
}

// 示例数据
const sampleRisks = [
  {
    created_at: '2024-01-15T10:30:00Z',
    userId: 'USER001',
    risk: 'low'
  },
  {
    created_at: '2024-01-20T14:45:00Z',
    userId: 'USER002',
    risk: 'medium'
  },
  {
    created_at: '2024-01-25T09:15:00Z',
    userId: 'USER003',
    risk: 'high'
  }
]

const services = ref([
  {
    title: '智能投顾服务',
    description: '基于人工智能算法的个性化投资建议，帮助您实现资产优化配置'
  },
  {
    title: '财富管理方案',
    description: '专业的财富管理团队为您提供全方位的资产管理和增值服务'
  },
  {
    title: '风险控制体系',
    description: '完善的风险评估和控制机制，保障您的投资安全'
  }
])

const riskLevelText = {
  low: '低风险',
  medium: '中等风险',
  high: '高风险'
}

const riskDescription = {
  low: '保守型投资者，偏好稳定收益',
  medium: '稳健型投资者，平衡风险与收益',
  high: '积极型投资者，追求高额回报'
}

const investmentStrategy = {
  low: '不加杠杆，稳健收益',
  medium: '2倍杠杆，积极获取高收益',
  high: '3倍杠杆，对标的波动疯狂套利'
}

// 计算属性
const hasRiskAssessment = computed(() => currentRisk.value !== null)

// 方法
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString('zh-CN')
}

const submitAssessment = () => {
  // 模拟评估逻辑
  const randomIndex = Math.floor(Math.random() * sampleRisks.length)
  currentRisk.value = sampleRisks[randomIndex]
  ElMessage.success('风险评估提交成功！')
}

// 生命周期
onMounted(() => {
  // 页面加载时调用获取用户风险等级数据
  fetchUserRisk()
})
</script>

<style scoped>
.risk-assessment-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.banner-section {
  margin-bottom: 40px;
}

.banner-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.section-header {
  text-align: center;
  margin-bottom: 30px;
}

.section-header h2 {
  color: #1890ff;
  margin-bottom: 10px;
}

.risk-result {
  max-width: 600px;
  margin: 0 auto;
}

.risk-card {
  margin-bottom: 30px;
}

.risk-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.risk-level h3 {
  margin: 0;
  font-size: 24px;
  font-weight: bold;
}

.risk-low {
  color: #52c41a;
}

.risk-medium {
  color: #faad14;
}

.risk-high {
  color: #f5222d;
}

.assessment-time {
  color: #666;
  font-size: 14px;
  margin-top: 5px;
}

.risk-details p {
  margin: 5px 0;
  color: #333;
}



.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  margin-top: 30px;
}

.service-card {
  transition: transform 0.3s ease;
}

.service-card:hover {
  transform: translateY(-5px);
}

.service-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.service-content {
  padding: 15px;
}

.service-content h4 {
  color: #1890ff;
  margin-bottom: 10px;
}

.service-content p {
  color: #666;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .risk-content {
    flex-direction: column;
    text-align: center;
  }
  
  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>