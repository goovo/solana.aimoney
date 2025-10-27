<template>
  <div class="gva-container2">
    <!-- 统计概览卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
      <gva-card custom-class="col-span-1">
        <div class="stat-card">
          <div class="stat-header">
            <div class="stat-icon market-cap">
              <el-icon><TrendCharts /></el-icon>
            </div>
            <div class="stat-title">总市值</div>
          </div>
          <div class="stat-value">${{ formatNumber(totalMarketCap) }}</div>
          <div class="stat-change" :class="getChangeClass(totalMarketCapChange)">
            {{ totalMarketCapChange > 0 ? '+' : '' }}{{ totalMarketCapChange.toFixed(2) }}%
          </div>
        </div>
      </gva-card>
      
      <gva-card custom-class="col-span-1">
        <div class="stat-card">
          <div class="stat-header">
            <div class="stat-icon volume">
              <el-icon><DataLine /></el-icon>
            </div>
            <div class="stat-title">24H总交易额</div>
          </div>
          <div class="stat-value">${{ formatNumber(totalVolume) }}</div>
          <div class="stat-change positive">
            <el-icon><ArrowUp /></el-icon>
            {{ formatNumber(totalVolumeChange) }}
          </div>
        </div>
      </gva-card>

      <gva-card custom-class="col-span-1">
        <div class="stat-card">
          <div class="stat-header">
            <div class="stat-icon btc">
              <el-icon><Coin /></el-icon>
            </div>
            <div class="stat-title">BTC主导地位</div>
          </div>
          <div class="stat-value">{{ btcDominance.toFixed(2) }}%</div>
          <div class="stat-change" :class="getChangeClass(btcDominanceChange)">
            {{ btcDominanceChange > 0 ? '+' : '' }}{{ btcDominanceChange.toFixed(2) }}%
          </div>
        </div>
      </gva-card>

      <gva-card custom-class="col-span-1">
        <div class="stat-card">
          <div class="stat-header">
            <div class="stat-icon active">
              <el-icon><Grid /></el-icon>
            </div>
            <div class="stat-title">活跃币种</div>
          </div>
          <div class="stat-value">{{ cryptoList.length }}</div>
          <div class="stat-change neutral">
            Top 20
          </div>
        </div>
      </gva-card>
    </div>

    <!-- 加密货币表格 -->
    <gva-card title="加密货币排行榜" show-action class="crypto-table-card">
      <template #action>
        <div class="table-actions">
          <div class="update-info">
            <el-icon><Timer /></el-icon>
            <span>更新: {{ lastUpdateTime }}</span>
          </div>
          <el-button 
            type="primary" 
            :loading="isRefreshing" 
            @click="refreshData"
            size="small"
          >
            <el-icon><Refresh /></el-icon>
            刷新
          </el-button>
        </div>
      </template>
      
      <div class="crypto-table">
        <el-table 
          :data="cryptoList" 
          style="width: 100%"
          :max-height="600"
          stripe
          v-loading="isRefreshing"
        >
          <el-table-column prop="rank" label="排名" width="80" align="center">
            <template #default="{ $index }">
              <span class="rank-badge" :class="`rank-${$index + 1}`">{{ $index + 1 }}</span>
            </template>
          </el-table-column>
          
          <el-table-column label="名称" width="180">
            <template #default="{ row }">
              <div class="crypto-info">
                <img :src="row.icon" :alt="row.name" class="crypto-icon">
                <div class="crypto-details">
                  <div class="crypto-symbol">{{ row.symbol }}</div>
                  <div class="crypto-name">{{ row.name }}</div>
                </div>
              </div>
            </template>
          </el-table-column>
          
          <el-table-column prop="price" label="价格" width="120" align="right">
            <template #default="{ row }">
              <span class="price-value">${{ formatPrice(row.price) }}</span>
            </template>
          </el-table-column>
          
          <el-table-column prop="change24h" label="24H涨跌" width="120" align="right">
            <template #default="{ row }">
              <span class="change-value" :class="getChangeClass(row.change24h)">
                {{ row.change24h > 0 ? '+' : '' }}{{ row.change24h.toFixed(2) }}%
              </span>
            </template>
          </el-table-column>
          
          <el-table-column prop="marketCap" label="市值" width="160" align="right">
            <template #default="{ row }">
              <span class="market-cap-value">${{ formatNumber(row.marketCap) }}</span>
            </template>
          </el-table-column>
          
          <el-table-column prop="volume24h" label="24H交易额" width="160" align="right">
            <template #default="{ row }">
              <span class="volume-value">${{ formatNumber(row.volume24h) }}</span>
            </template>
          </el-table-column>
          
          <el-table-column prop="circulatingSupply" label="流通量" width="160" align="right">
            <template #default="{ row }">
              <span class="supply-value">{{ formatNumber(row.circulatingSupply) }} {{ row.symbol }}</span>
            </template>
          </el-table-column>
          
          <el-table-column label="趋势" width="150">
            <template #default="{ row }">
              <div class="mini-chart" :id="`chart-${row.id}`"></div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </gva-card>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { Timer, Refresh, Loading, TrendCharts, DataLine, Coin, Grid, ArrowUp } from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import { getCryptoData } from '@/api/crypto'

// 响应式数据
const cryptoList = ref([])
const isRefreshing = ref(false)
const lastUpdateTime = ref('')
let refreshTimer = null

// 统计数据
const totalMarketCap = ref(0)
const totalMarketCapChange = ref(0)
const totalVolume = ref(0)
const totalVolumeChange = ref(0)
const btcDominance = ref(0)
const btcDominanceChange = ref(0)

// 从API获取加密货币数据
const fetchCryptoData = async () => {
  try {
    const response = await getCryptoData()
    if (response.code === 0) {
      return response.data.cryptoList
    }
    throw new Error(response.msg || '获取数据失败')
  } catch (error) {
    console.error('获取加密货币数据失败:', error)
    // 返回空数组，前端会显示加载失败状态
    return []
  }
}

// 生成趋势数据
const generateTrendData = () => {
  const data = []
  let value = Math.random() * 100
  for (let i = 0; i < 24; i++) {
    value += (Math.random() - 0.5) * 10
    data.push(Math.max(0, value))
  }
  return data
}

// 格式化数字
const formatNumber = (num) => {
  if (num >= 1e12) return (num / 1e12).toFixed(2) + 'T'
  if (num >= 1e9) return (num / 1e9).toFixed(2) + 'B'
  if (num >= 1e6) return (num / 1e6).toFixed(2) + 'M'
  if (num >= 1e3) return (num / 1e3).toFixed(2) + 'K'
  return num.toFixed(2)
}

const formatPrice = (price) => {
  if (price >= 1000) return price.toFixed(0)
  if (price >= 1) return price.toFixed(2)
  return price.toFixed(4)
}

// 获取涨跌样式类
const getChangeClass = (change) => {
  if (change > 0) return 'positive'
  if (change < 0) return 'negative'
  return 'neutral'
}

// 获取粒子样式
const getParticleStyle = (index) => {
  return {
    left: Math.random() * 100 + '%',
    top: Math.random() * 100 + '%',
    animationDelay: Math.random() * 20 + 's',
    animationDuration: (Math.random() * 20 + 10) + 's'
  }
}

// 刷新数据
const refreshData = async () => {
  if (isRefreshing.value) return
  
  isRefreshing.value = true
  try {
    const response = await getCryptoData()
    if (response.code === 0) {
      cryptoList.value = response.data.cryptoList
      
      // 更新统计数据
      totalMarketCap.value = response.data.statistics.totalMarketCap
      totalMarketCapChange.value = response.data.statistics.totalMarketCapChange
      totalVolume.value = response.data.statistics.totalVolume
      totalVolumeChange.value = response.data.statistics.totalVolumeChange
      btcDominance.value = response.data.statistics.btcDominance
      btcDominanceChange.value = response.data.statistics.btcDominanceChange
      
      updateLastUpdateTime()
      
      // 渲染趋势图表
      await nextTick()
      renderTrendCharts()
    } else {
      throw new Error(response.msg || '获取数据失败')
    }
  } catch (error) {
    console.error('刷新数据失败:', error)
    ElMessage.error('刷新数据失败，请稍后重试')
  } finally {
    isRefreshing.value = false
  }
}


// 更新最后更新时间
const updateLastUpdateTime = () => {
  const now = new Date()
  lastUpdateTime.value = now.toLocaleTimeString('zh-CN')
}

// 渲染趋势图表
const renderTrendCharts = () => {
  cryptoList.value.forEach((crypto, index) => {
    const chartDom = document.getElementById(`chart-${crypto.id}`)
    if (chartDom) {
      const chart = echarts.init(chartDom)
      const option = {
        grid: {
          left: 0,
          right: 0,
          top: 0,
          bottom: 0
        },
        xAxis: {
          type: 'category',
          show: false,
          data: Array(24).fill('')
        },
        yAxis: {
          type: 'value',
          show: false
        },
        series: [{
          data: crypto.trendData,
          type: 'line',
          smooth: true,
          showSymbol: false,
          lineStyle: {
            width: 1,
            color: crypto.change24h >= 0 ? '#10b981' : '#ef4444'
          },
          areaStyle: {
            color: {
              type: 'linear',
              x: 0,
              y: 0,
              x2: 0,
              y2: 1,
              colorStops: [{
                offset: 0,
                color: crypto.change24h >= 0 ? 'rgba(16, 185, 129, 0.3)' : 'rgba(239, 68, 68, 0.3)'
              }, {
                offset: 1,
                color: crypto.change24h >= 0 ? 'rgba(16, 185, 129, 0)' : 'rgba(239, 68, 68, 0)'
              }]
            }
          }
        }]
      }
      chart.setOption(option)
      
      // 响应式调整
      const resizeHandler = () => chart.resize()
      window.addEventListener('resize', resizeHandler)
      
      // 组件卸载时移除监听器
      onUnmounted(() => {
        window.removeEventListener('resize', resizeHandler)
        chart.dispose()
      })
    }
  })
}

// 启动自动刷新
const startAutoRefresh = () => {
  refreshTimer = setInterval(() => {
    refreshData()
  }, 5 * 60 * 1000) // 5分钟
}

// 停止自动刷新
const stopAutoRefresh = () => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
    refreshTimer = null
  }
}

// 生命周期钩子
onMounted(() => {
  refreshData()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<style lang="scss" scoped>
// 统计卡片样式
.stat-card {
  padding: 16px;
  
  .stat-header {
    display: flex;
    align-items: center;
    margin-bottom: 12px;
    
    .stat-icon {
      width: 40px;
      height: 40px;
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 12px;
      font-size: 20px;
      
      &.market-cap {
        background: rgba(59, 130, 246, 0.1);
        color: #3b82f6;
      }
      
      &.volume {
        background: rgba(16, 185, 129, 0.1);
        color: #10b981;
      }
      
      &.btc {
        background: rgba(251, 191, 36, 0.1);
        color: #fbbf24;
      }
      
      &.active {
        background: rgba(139, 92, 246, 0.1);
        color: #8b5cf6;
      }
    }
    
    .stat-title {
      font-size: 14px;
      color: #666;
    }
  }
  
  .stat-value {
    font-size: 24px;
    font-weight: bold;
    color: #333;
    margin-bottom: 8px;
  }
  
  .stat-change {
    font-size: 12px;
    
    &.positive {
      color: #10b981;
    }
    
    &.negative {
      color: #ef4444;
    }
    
    &.neutral {
      color: #666;
    }
  }
}

// 表格操作区
.table-actions {
  display: flex;
  align-items: center;
  gap: 16px;
  
  .update-info {
    display: flex;
    align-items: center;
    gap: 4px;
    color: #666;
    font-size: 12px;
  }
}

// 排名徽章
.rank-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  font-weight: bold;
  font-size: 12px;
  
  &.rank-1 {
    background: linear-gradient(135deg, #ffd700, #ffed4e);
    color: #1e293b;
  }
  
  &.rank-2 {
    background: linear-gradient(135deg, #c0c0c0, #e8e8e8);
    color: #1e293b;
  }
  
  &.rank-3 {
    background: linear-gradient(135deg, #cd7f32, #e4a054);
    color: #1e293b;
  }
}

// 加密货币信息
.crypto-info {
  display: flex;
  align-items: center;
  gap: 8px;
  
  .crypto-icon {
    width: 24px;
    height: 24px;
    border-radius: 50%;
  }
  
  .crypto-details {
    .crypto-symbol {
      font-weight: 600;
      font-size: 14px;
      color: #333;
    }
    
    .crypto-name {
      font-size: 12px;
      color: #666;
    }
  }
}

// 价格和数值
.price-value, .market-cap-value, .volume-value, .supply-value {
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 13px;
}

// 涨跌幅
.change-value {
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  
  &.positive {
    color: #10b981;
    background: rgba(16, 185, 129, 0.1);
  }
  
  &.negative {
    color: #ef4444;
    background: rgba(239, 68, 68, 0.1);
  }
}

// 趋势图表
.mini-chart {
  height: 30px;
  width: 100%;
}

// 表格卡片特殊样式
:deep(.crypto-table-card) {
  .el-card__body {
    padding: 0;
  }
}

:deep(.el-table) {
  .el-table__header-wrapper {
    th {
      background: #fafafa;
      font-weight: 600;
      color: #333;
    }
  }
  
  .el-table__row {
    &:hover {
      background: #f5f5f5;
    }
  }
}
</style>