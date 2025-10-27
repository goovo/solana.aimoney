<template>
  <div class="risk-assessment-container">
    <!-- Banner轮播图 -->
    <div class="banner-section">
      <div class="banner-content">
        <h1>投资风险承受能力评估</h1>
        <p>通过专业的风险评估问卷，帮助您了解自身的投资风险偏好</p>
        <p>科学评估您的风险承受能力，为您提供合适的投资建议</p>
        <p>基于大数据分析和金融工程模型，确保评估结果的准确性</p>
      </div>
      <img :src="bannerImage" alt="风险评估banner" class="banner-image">
    </div>

    <!-- 评估问题表单 -->
    <div class="assessment-section">
      <div class="assessment-card">
        <h2>风险承受能力评估问卷</h2>
        <p class="assessment-description">请根据您的实际情况选择最符合的选项，每个问题必须选择一个答案</p>
        
        <div class="questions-container">
          <div v-for="(question, index) in questions" :key="index" class="question-item">
            <h3>{{ index + 1 }}. {{ question.title }}</h3>
            <div class="options-group">
              <label
                v-for="(option, optIndex) in question.options"
                :key="optIndex"
                class="option-label"
                :class="{ active: question.selected === option.score }"
              >
                <input
                  type="radio"
                  :name="'question' + index"
                  :value="option.score"
                  v-model.number="question.selected"
                  @change="updateSelection"
                >
                <span class="option-text">{{ option.text }}</span>
              </label>
            </div>
          </div>
        </div>

        <!-- 提交按钮 -->
        <button 
          class="submit-btn" 
          :class="{ disabled: !allQuestionsAnswered || submitting }"
          :disabled="!allQuestionsAnswered || submitting"
          @click="submitAssessment"
        >
          提交评估
        </button>

        <!-- 结果显示 -->
        <div v-if="showResult" class="result-section">
          <h3>您的风险评估结果</h3>
          <div :class="['result-card', resultClass]">
            <h4>{{ riskResult.title }}</h4>
            <p>{{ riskResult.description }}</p>
            <p>建议投资类型: {{ riskResult.suggestion }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 产品服务介绍 -->
    <div class="services-section">
      <h2>我们的投资服务</h2>
      <div class="services-grid">
        <div v-for="(service, index) in services" :key="index" class="service-card">
          <img :src="service.image" :alt="service.title" class="service-image">
          <h3>{{ service.title }}</h3>
          <p>{{ service.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, reactive } from 'vue'
import { useRouter } from 'vue-router' // 新增：从 vue-router 获取路由实例（中文注释）
import { ElMessage } from 'element-plus'
import { setUserRisk } from '@/api/running/sysUserRisk'

const router = useRouter() // 新增：获取路由实例，后续用于编程式导航（中文注释）
const bannerImage = ref('https://picsum.photos/1200/400?random=1')

const questions = reactive([
  {
    title: '您的投资经验如何？',
    options: [
      { text: '无投资经验', score: 1 },
      { text: '少于2年', score: 2 },
      { text: '2-5年', score: 3 },
      { text: '5年以上', score: 4 }
    ],
    selected: null
  },
  {
    title: '您能承受的最大投资亏损是多少？',
    options: [
      { text: '不能承受任何亏损', score: 1 },
      { text: '10%以内', score: 2 },
      { text: '10%-30%', score: 3 },
      { text: '30%以上', score: 4 }
    ],
    selected: null
  },
  {
    title: '您的投资期限是多久？',
    options: [
      { text: '1年以内', score: 1 },
      { text: '1-3年', score: 2 },
      { text: '3-5年', score: 3 },
      { text: '5年以上', score: 4 }
    ],
    selected: null
  },
  {
    title: '您的年收入水平？',
    options: [
      { text: '10万以下', score: 1 },
      { text: '10-30万', score: 2 },
      { text: '30-50万', score: 3 },
      { text: '50万以上', score: 4 }
    ],
    selected: null
  },
  {
    title: '您的投资主要目标是？',
    options: [
      { text: '保值增值', score: 1 },
      { text: '稳健收益', score: 2 },
      { text: '较高收益', score: 3 },
      { text: '最大化收益', score: 4 }
    ],
    selected: null
  },
  {
    title: '您对投资产品的了解程度？',
    options: [
      { text: '不了解', score: 1 },
      { text: '基本了解', score: 2 },
      { text: '比较了解', score: 3 },
      { text: '非常了解', score: 4 }
    ],
    selected: null
  },
  {
    title: '您的家庭财务状况？',
    options: [
      { text: '有负债压力', score: 1 },
      { text: '收支平衡', score: 2 },
      { text: '略有结余', score: 3 },
      { text: '财务自由', score: 4 }
    ],
    selected: null
  },
  {
    title: '您能投入多少时间关注投资？',
    options: [
      { text: '几乎没时间', score: 1 },
      { text: '偶尔关注', score: 2 },
      { text: '经常关注', score: 3 },
      { text: '全职关注', score: 4 }
    ],
    selected: null
  },
  {
    title: '您对市场波动的反应？',
    options: [
      { text: '非常担忧', score: 1 },
      { text: '有些担忧', score: 2 },
      { text: '相对平静', score: 3 },
      { text: '视为机会', score: 4 }
    ],
    selected: null
  }
])

const showResult = ref(false)
const riskLevel = ref('')
// 提交按钮的loading状态，防止重复提交（中文注释）
const submitting = ref(false)

const allQuestionsAnswered = computed(() => {
  return questions.every(q => q.selected !== null)
})

const riskResult = computed(() => {
  const results = {
    low: {
      title: '低风险承受能力',
      description: '您适合稳健型投资策略',
      suggestion: '货币基金、国债、银行理财等低风险产品'
    },
    medium: {
      title: '中等风险承受能力',
      description: '您适合平衡型投资策略',
      suggestion: '混合基金、债券基金、指数基金等中等风险产品'
    },
    high: {
      title: '高风险承受能力',
      description: '您适合进取型投资策略',
      suggestion: '股票基金、股权投资、衍生品等高收益产品'
    }
  }
  return results[riskLevel.value] || {}
})

const resultClass = computed(() => {
  return `risk-${riskLevel.value}`
})

const updateSelection = () => {
  // 任何一次选择发生变化，都隐藏上一次的结果，避免误导
  showResult.value = false
}

const calculateRiskLevel = () => {
  // 将 selected 强制转换为数字，避免字符串参与求和造成类型问题
  const totalScore = questions.reduce((sum, q) => sum + Number(q.selected || 0), 0)
  const averageScore = totalScore / questions.length

  if (averageScore <= 2) {
    riskLevel.value = 'low'
  } else if (averageScore <= 3) {
    riskLevel.value = 'medium'
  } else {
    riskLevel.value = 'high'
  }
  
  showResult.value = true
}

// 提交评估：先计算风险等级，再通过 API 持久化到后端（中文注释）
const submitAssessment = async () => {
  // 前置校验：必须答完所有题目（中文注释）
  if (!allQuestionsAnswered.value) {
    ElMessage({ type: 'warning', message: '请先完成所有题目再提交' })
    return
  }

  // 计算风险等级并展示本地结果（中文注释）
  calculateRiskLevel()

  // 构建提交数据，仅需提交风险等级，用户ID由后端从token中获取（中文注释）
  const payload = { risk: riskLevel.value }

  try {
    submitting.value = true
    const res = await setUserRisk(payload)
    if (res && res.code === 0) {
      ElMessage({ type: 'success', message: '评估提交成功' })
      // 延迟 2 秒再跳转，确保用户看到成功提示（中文注释）
      setTimeout(() => {
        router.push({ name: 'currentRisk' }) // 原来是立即跳转，这里延迟 2 秒执行（中文注释）
      }, 2000)
    } else {
      ElMessage({ type: 'error', message: res?.msg || '提交失败，请稍后重试' })
    }
  } catch (e) {
    ElMessage({ type: 'error', message: e?.message || '提交异常，请检查网络后重试' })
  } finally {
    submitting.value = false
  }
}

const services = [
  {
    title: '智能投顾服务',
    description: '基于人工智能算法的个性化投资建议，为您提供专业的资产配置方案',
    image: 'https://picsum.photos/300/200?random=2'
  },
  {
    title: '财富管理',
    description: '专业的财富管理团队，为您提供全方位的资产增值和风险控制服务',
    image: 'https://picsum.photos/300/200?random=3'
  },
  {
    title: '投资咨询',
    description: '资深投资顾问提供一对一的投资咨询服务，帮助您做出明智的投资决策',
    image: 'https://picsum.photos/300/200?random=4'
  }
]
</script>

<style scoped>
.risk-assessment-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.banner-section {
  display: flex;
  align-items: center;
  gap: 40px;
  margin-bottom: 60px;
  background: linear-gradient(135deg, #1976d2 0%, #42a5f5 100%);
  padding: 40px;
  border-radius: 12px;
  color: white;
}

.banner-content h1 {
  font-size: 2.5em;
  margin-bottom: 20px;
}

.banner-content p {
  font-size: 1.1em;
  margin-bottom: 10px;
  line-height: 1.6;
}

.banner-image {
  width: 50%;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.assessment-section {
  margin-bottom: 60px;
}

.assessment-card {
  background: white;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.assessment-card h2 {
  color: #1976d2;
  margin-bottom: 20px;
  text-align: center;
}

.assessment-description {
  text-align: center;
  color: #666;
  margin-bottom: 30px;
  font-size: 1.1em;
}

.questions-container {
  margin-bottom: 30px;
}

.question-item {
  margin-bottom: 30px;
  padding: 20px;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  background: #fafafa;
}

.question-item h3 {
  color: #333;
  margin-bottom: 15px;
}

.options-group {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 12px;
}

.option-label {
  display: flex;
  align-items: center;
  padding: 12px;
  border: 2px solid #e0e0e0;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  /* 确保可点击 */
  pointer-events: auto;
}

.option-label:hover {
  border-color: #1976d2;
  background: #f5f9ff;
}

/* 选中态：高亮边框与背景，增强可见性 */
.option-label.active {
  border-color: #1976d2;
  background: #f0f7ff;
  box-shadow: 0 0 0 2px rgba(25, 118, 210, 0.08) inset;
}

.option-label input[type="radio"] {
  margin-right: 10px;
  /* 提升原生单选圆点的可见性（现代浏览器支持） */
  accent-color: #1976d2;
}

/* 当原生单选被选中时，文字也高亮 */
.option-label input[type="radio"]:checked + .option-text {
  color: #1976d2;
  font-weight: 600;
}

.option-text {
  font-size: 0.95em;
}

.submit-btn {
  width: 100%;
  padding: 15px;
  background: #1976d2;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1.1em;
  cursor: pointer;
  transition: background 0.3s ease;
}

.submit-btn:hover:not(.disabled) {
  background: #1565c0;
}

.submit-btn.disabled {
  background: #ccc;
  cursor: not-allowed;
}

.result-section {
  margin-top: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.result-section h3 {
  color: #1976d2;
  margin-bottom: 20px;
  text-align: center;
}

.result-card {
  padding: 25px;
  border-radius: 8px;
  text-align: center;
  color: white;
}

.risk-low {
  background: linear-gradient(135deg, #4caf50 0%, #66bb6a 100%);
}

.risk-medium {
  background: linear-gradient(135deg, #ff9800 0%, #ffb74d 100%);
}

.risk-high {
  background: linear-gradient(135deg, #f44336 0%, #ef5350 100%);
}

.result-card h4 {
  font-size: 1.5em;
  margin-bottom: 15px;
}

.result-card p {
  margin-bottom: 10px;
  line-height: 1.6;
}

.services-section {
  margin-bottom: 40px;
}

.services-section h2 {
  color: #1976d2;
  text-align: center;
  margin-bottom: 40px;
  font-size: 2em;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.service-card {
  background: white;
  padding: 25px;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s ease;
}

.service-card:hover {
  transform: translateY(-5px);
}

.service-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
  margin-bottom: 20px;
}

.service-card h3 {
  color: #1976d2;
  margin-bottom: 15px;
  font-size: 1.3em;
}

.service-card p {
  color: #666;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .banner-section {
    flex-direction: column;
    text-align: center;
  }
  
  .banner-image {
    width: 100%;
    margin-top: 20px;
  }
  
  .options-group {
    grid-template-columns: 1fr;
  }
  
  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>

