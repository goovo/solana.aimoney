<template>
  <div class="about-container">
    <!-- Banner轮播图 -->
    <section class="banner-section">
      <div class="banner-slider">
        <img :src="bannerImages[currentBannerIndex]" alt="AI赚钱平台" class="banner-image">
        <div class="banner-content">
          <h1>智能量化投资 引领数字资产新纪元</h1>
          <p>基于深度学习的AI策略系统，为您的加密资产投资保驾护航</p>
          <button class="cta-button">立即体验</button>
        </div>
        <div class="banner-controls">
          <button v-for="(_, index) in bannerImages" 
                  :key="index"
                  :class="['banner-dot', { active: currentBannerIndex === index }]"
                  @click="currentBannerIndex = index">
          </button>
        </div>
      </div>
    </section>

    <!-- 核心优势展示 -->
    <section class="advantages-section">
      <h2>我们的核心优势</h2>
      <div class="advantages-grid">
        <div class="advantage-item" v-for="(item, index) in advantages" :key="index">
          <img :src="item.image" :alt="item.title">
          <h3>{{ item.title }}</h3>
          <p>{{ item.description }}</p>
        </div>
      </div>
    </section>

    <!-- 产品服务介绍 -->
    <section class="services-section">
      <h2>产品服务体系</h2>
      <div class="services-container">
        <div class="service-card" v-for="(service, index) in services" :key="index">
          <div class="service-icon">
            <img :src="service.icon" :alt="service.title">
          </div>
          <h3>{{ service.title }}</h3>
          <p>{{ service.description }}</p>
          <ul>
            <li v-for="(feature, featureIndex) in service.features" :key="featureIndex">{{ feature }}</li>
          </ul>
        </div>
      </div>
    </section>

    <!-- 数据统计展示 -->
    <section class="stats-section">
      <h2>平台数据表现</h2>
      <div class="stats-grid">
        <div class="stat-item" v-for="stat in statistics" :key="stat.label">
          <div class="stat-value">{{ stat.value }}</div>
          <div class="stat-label">{{ stat.label }}</div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const currentBannerIndex = ref(0)

const bannerImages = [
  'https://picsum.photos/1200/600?random=1',
  'https://picsum.photos/1200/600?random=2',
  'https://picsum.photos/1200/600?random=3'
]

const advantages = [
  {
    image: 'https://picsum.photos/300/200?random=4',
    title: 'AI量化策略',
    description: '基于深度神经网络的多因子量化模型，实现智能投资决策'
  },
  {
    image: 'https://picsum.photos/300/200?random=5',
    title: '风险控制体系',
    description: '多层次风险管理系统，确保投资组合的稳健性和安全性'
  },
  {
    image: 'https://picsum.photos/300/200?random=6',
    title: '实时监控',
    description: '7×24小时全天候监控市场动态，及时调整投资策略'
  }
]

const services = [
  {
    icon: 'https://picsum.photos/100/100?random=7',
    title: '量化交易策略',
    description: '专业级的量化交易解决方案',
    features: ['多因子模型', '机器学习算法', '实时调仓']
  },
  {
    icon: 'https://picsum.photos/100/100?random=8',
    title: '资产管理系统',
    description: '全方位的资产配置服务',
    features: ['投资组合优化', '风险评估', '收益分析']
  },
  {
    icon: 'https://picsum.photos/100/100?random=9',
    title: '数据分析平台',
    description: '深度市场数据分析工具',
    features: ['大数据处理', '可视化分析', '预测模型']
  }
]

const statistics = [
  { value: '78.2%', label: '策略胜率' },
  { value: '25.2%', label: '年均收益率' },
  { value: '5000+', label: '活跃用户' },
  { value: '99.9%', label: '系统稳定性' }
]
</script>

<style scoped>
.about-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.banner-section {
  position: relative;
  margin-bottom: 60px;
}

.banner-slider {
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.banner-image {
  width: 100%;
  height: 400px;
  object-fit: cover;
}

.banner-content {
  position: absolute;
  top: 50%;
  left: 50px;
  transform: translateY(-50%);
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.banner-content h1 {
  font-size: 2.5rem;
  margin-bottom: 16px;
  font-weight: 600;
}

.banner-content p {
  font-size: 1.2rem;
  margin-bottom: 24px;
}

.cta-button {
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
  color: white;
  padding: 12px 32px;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(37, 99, 235, 0.3);
}

.banner-controls {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 12px;
}

.banner-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  border: none;
  background: rgba(255, 255, 255, 0.5);
  cursor: pointer;
  transition: all 0.3s ease;
}

.banner-dot.active {
  background: white;
  transform: scale(1.2);
}

.advantages-section,
.services-section,
.stats-section {
  margin-bottom: 80px;
}

.advantages-section h2,
.services-section h2,
.stats-section h2 {
  text-align: center;
  font-size: 2.2rem;
  margin-bottom: 40px;
  color: #1e40af;
}

.advantages-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.advantage-item {
  text-align: center;
  padding: 30px;
  border-radius: 12px;
  background: white;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.advantage-item:hover {
  transform: translateY(-5px);
}

.advantage-item img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
  margin-bottom: 20px;
}

.advantage-item h3 {
  font-size: 1.4rem;
  margin-bottom: 15px;
  color: #2563eb;
}

.advantage-item p {
  color: #64748b;
  line-height: 1.6;
}

.services-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
}

.service-card {
  background: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: all 0.3s ease;
}

.service-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(37, 99, 235, 0.15);
}

.service-icon img {
  width: 80px;
  height: 80px;
  margin-bottom: 20px;
}

.service-card h3 {
  font-size: 1.4rem;
  margin-bottom: 15px;
  color: #2563eb;
}

.service-card p {
  color: #64748b;
  margin-bottom: 20px;
  line-height: 1.6;
}

.service-card ul {
  text-align: left;
  color: #475569;
  padding-left: 20px;
}

.service-card li {
  margin-bottom: 8px;
  line-height: 1.5;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 30px;
}

.stat-item {
  text-align: center;
  padding: 30px;
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
  color: white;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(37, 99, 235, 0.2);
}

.stat-value {
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 10px;
}

.stat-label {
  font-size: 1.1rem;
  opacity: 0.9;
}

@media (max-width: 768px) {
  .banner-content {
    left: 20px;
    right: 20px;
    text-align: center;
  }
  
  .banner-content h1 {
    font-size: 2rem;
  }
  
  .advantages-grid,
  .services-container,
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .advantage-item,
  .service-card,
  .stat-item {
    padding: 20px;
  }
}
</style>
