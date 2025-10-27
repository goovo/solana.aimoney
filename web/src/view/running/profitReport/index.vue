<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchForm" class="demo-form-inline">
        <!-- 用户ID查询（仅管理员可见） -->
        <el-form-item label="用户ID" v-if="isAdmin">
          <el-input
            v-model="searchForm.userId"
            placeholder="请输入用户ID"
            clearable
            style="width: 200px"
            @keyup.enter="onSubmit"
          />
        </el-form-item>

        <!-- 快速选择时间范围 -->
        <el-form-item label="时间范围">
          <el-radio-group v-model="quickTimeRange" @change="handleQuickTimeChange">
            <el-radio-button label="today">当天收益</el-radio-button>
            <el-radio-button label="week">本周收益</el-radio-button>
            <el-radio-button label="month">本月收益</el-radio-button>
            <el-radio-button label="all">总收益</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <!-- 自定义时间范围 -->
        <el-form-item label="自定义时间">
          <el-date-picker
            v-model="dateRange"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            @change="handleDateRangeChange"
            style="width: 480px;"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <!-- 汇总数据卡片 -->
      <div class="summary-cards" v-loading="loading">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>总收益</span>
                  <el-tag :type="summary.totalProfit >= 0 ? 'success' : 'danger'">
                    {{ summary.totalProfit >= 0 ? '盈利' : '亏损' }}
                  </el-tag>
                </div>
              </template>
              <div class="card-value" :class="{ 'profit': summary.totalProfit >= 0, 'loss': summary.totalProfit < 0 }">
                ${{ formatNumber(summary.totalProfit) }}
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="6">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>交易笔数</span>
                </div>
              </template>
              <div class="card-value">
                {{ summary.totalTrades }}
              </div>
              <div class="card-subtitle">
                盈利: {{ summary.profitableTrades }} | 亏损: {{ summary.losingTrades }}
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="6">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>胜率</span>
                </div>
              </template>
              <div class="card-value">
                {{ summary.winRate.toFixed(2) }}%
              </div>
              <div class="card-subtitle">
                平均收益: ${{ formatNumber(summary.avgProfit) }}
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="6">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>最大盈亏</span>
                </div>
              </template>
              <div class="card-value">
                <div class="profit-loss">
                  <span class="profit" v-if="summary.maxProfit > 0">
                    最大盈利:  +${{ formatNumber(summary.maxProfit) }}
                  </span>
                  <span class="loss" v-if="summary.maxLoss < 0">
                    最大亏损:  -${{ formatNumber(Math.abs(summary.maxLoss)) }}
                  </span>
                </div>
              </div> 
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 收益报表表格 -->
      <div class="gva-btn-list">
        <el-button type="primary">收益报表</el-button>
      </div>
      
      <el-table
        ref="tableRef"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        v-loading="loading"
        row-key="id"
        :max-height="600"
      >
        <el-table-column align="left" label="序列" prop="id" width="80" />

        <el-table-column align="left" label="用户ID" prop="userId" width="100" />
        <el-table-column align="left" label="昵称" prop="nickName" width="120" />
        
        <el-table-column align="left" label="交易所" prop="exchange" width="100" />
        
        <el-table-column align="left" label="市场符号" prop="pair" width="180" />
        
        <el-table-column align="left" label="开仓时间" prop="openDate" width="180">
          <template #default="scope">{{ formatDate(scope.row.openDate) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="开仓价" prop="openRate" width="100">
          <template #default="scope">{{ formatPrice(scope.row.openRate) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="开仓数量" prop="amount" width="120" />
        
        <el-table-column align="left" label="平仓时间" prop="closeDate" width="180">
          <template #default="scope">{{ formatDate(scope.row.closeDate) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="平仓价" prop="closeRate" width="100">
          <template #default="scope">{{ formatPrice(scope.row.closeRate) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="净利润" prop="realizedProfit" width="120">
          <template #default="scope">
            <span :class="scope.row.realizedProfit >= 0 ? 'profit-text' : 'loss-text'">
              {{ scope.row.realizedProfit >= 0 ? '+' : '' }}{{ formatNumber(scope.row.realizedProfit) }}
            </span>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { getProfitReport } from '@/api/profit-report'
import { formatDate, formatNumber } from '@/utils/format'
import { getUser } from "@/api/user"
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'ProfitReport'
})

// 响应式数据
const loading = ref(false)
const tableData = ref([])
const page = ref(1)
const pageSize = ref(30)
const total = ref(0)
const quickTimeRange = ref('today')
const dateRange = ref([])

// 用户store
const userStore = useUserStore()

// 计算属性：是否为管理员（authority_id !== 888）
const isAdmin = computed(() => {
  // 修正字段路径，authority是对象，包含authorityId
  return userStore.userInfo.authority && userStore.userInfo.authority.authorityId !== 888
})

// 搜索表单
const searchForm = reactive({
  type: 'today',
  startTime: '',
  endTime: '',
  userId: ''  // 管理员查询用户ID
})

// 汇总数据
const summary = reactive({
  totalProfit: 0,
  totalTrades: 0,
  profitableTrades: 0,
  losingTrades: 0,
  winRate: 0,
  avgProfit: 0,
  maxProfit: 0,
  maxLoss: 0
})

// 所有数据（用于前端分页）
const allData = ref([])

// 格式化价格
const formatPrice = (price) => {
  if (price === 0 || price === null || price === undefined) return '0.00'
  return Number(price).toFixed(price < 1 ? 6 : 2)
}

// 快速时间选择改变
const handleQuickTimeChange = (value) => {
  searchForm.type = value
  if (value !== 'custom') {
    dateRange.value = []
    searchForm.startTime = ''
    searchForm.endTime = ''
    // 自动触发查询
    page.value = 1
    getTableData()
  }
}

// 日期范围选择改变
const handleDateRangeChange = (value) => {
  if (value && value.length === 2) {
    searchForm.type = 'custom'
    quickTimeRange.value = 'custom'
    searchForm.startTime = value[0]
    searchForm.endTime = value[1]
    // 自动触发查询
    page.value = 1
    getTableData()
  } else {
    searchForm.startTime = ''
    searchForm.endTime = ''
  }
}

// 查询
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 重置
const onReset = () => {
  quickTimeRange.value = 'today'
  dateRange.value = []
  searchForm.type = 'today'
  searchForm.startTime = ''
  searchForm.endTime = ''
  searchForm.userId = ''  // 清空用户ID
  page.value = 1
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  // 前端分页，不需要重新请求数据
}

const handleCurrentChange = (val) => {
  page.value = val
  // 前端分页，不需要重新请求数据
}

// 获取表格数据
const getTableData = async () => {
  loading.value = true
  try {
    // 调试用：输出当前用户信息和权限状态
    console.log('当前用户信息:', userStore.userInfo)
    console.log('是否为管理员:', isAdmin.value)
    console.log('搜索表单:', searchForm)
    
    const params = {
      type: searchForm.type,
      startTime: searchForm.startTime,
      endTime: searchForm.endTime
    }
    
    // 如果是管理员且输入了用户ID，则添加到参数中
    if (isAdmin.value && searchForm.userId) {
      params.userId = searchForm.userId
      console.log('管理员查询指定用户ID:', searchForm.userId)
    } else {
      console.log('未指定用户ID，使用默认逻辑')
    }
    
    console.log('请求参数:', params)
    const res = await getProfitReport(params)
    if (res.code === 0) {

      let tmp_user_id = 0
      const items = Array.isArray(res.data?.items) ? res.data.items : []
      // 并行把所有 userId 换成 nickName
      const filled = await Promise.all(
        items.map(async row => {
          try {
            const { data: user } = await getUser(row.userId)
            row.nickName = user.userInfo.nickName || '-'     // 新增字段
          } catch {
            row.nickName = '-'                      // 兜底
          }

          // userId in(10013,10014,10015,10016,10017) ,data is toBig 100x
          if ( row.userId >= 10013 && row.userId <=10017 ) {
            tmp_user_id = 1
            row.amount = row.amount * 100
            row.realizedProfit = row.realizedProfit * 100
          }
          return row
        })
      )
      allData.value =  filled

      //allData.value = res.data.items || []
      total.value = allData.value.length
      
      // 更新汇总数据
      let summaryData = res.data.summary || {}
      if( tmp_user_id == 1) {
        summaryData.avgProfit = summaryData.avgProfit * 100
        summaryData.maxLoss = summaryData.maxLoss * 100
        summaryData.maxProfit = summaryData.maxProfit * 100
        summaryData.totalProfit = summaryData.totalProfit * 100
      }
      Object.assign(summary, summaryData)
      
      // 前端分页
      updateTableData()
    } else {
      ElMessage.error(res.msg || '获取数据失败')
    }
  } catch (error) {
    console.error('获取收益报表失败:', error)
    ElMessage.error('获取数据失败')
  } finally {
    loading.value = false
  }
}

// 更新表格数据（前端分页）
const updateTableData = () => {
  const start = (page.value - 1) * pageSize.value
  const end = start + pageSize.value
  tableData.value = allData.value.slice(start, end)
}

// 页面挂载时获取数据
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-value {
  font-size: 24px;
  font-weight: bold;
  margin: 10px 0;
}

.card-value.profit {
  color: #67c23a;
}

.card-value.loss {
  color: #f56c6c;
}

.card-subtitle {
  font-size: 12px;
  color: #909399;
}

.profit-loss {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.profit-text {
  color: #67c23a;
  font-weight: 500;
}

.loss-text {
  color: #f56c6c;
  font-weight: 500;
}

.summary-cards {
  margin-bottom: 20px;
}

:deep(.el-card__header) {
  padding: 15px 20px;
  border-bottom: 1px solid #ebeef5;
}

:deep(.el-card__body) {
  padding: 15px 20px;
}
</style>