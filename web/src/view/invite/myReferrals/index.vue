<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="搜索">
          <el-input
            v-model="searchForm.keyword"
            placeholder="请输入用户昵称或用户名"
            clearable
            @keyup.enter="handleSearch"
            style="width: 300px"
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" icon="search" @click="handleSearch">查询</el-button>
          <el-button icon="refresh" @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <!-- 汇总数据卡片 -->
      <el-row :gutter="20" class="summary-row">
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>总推荐人数</span>
              </div>
            </template>
            <div class="summary-content">
              <span class="summary-number">{{ totalReferral }}</span>
              <span class="summary-unit">人</span>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>已投资人数</span>
              </div>
            </template>
            <div class="summary-content">
              <span class="summary-number">{{ investedCount }}</span>
              <span class="summary-unit">人</span>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>转化率</span>
              </div>
            </template>
            <div class="summary-content">
              <span class="summary-number">{{ conversionRate }}</span>
              <span class="summary-unit">%</span>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 推荐列表表格 -->
      <el-table
        :data="referralList"
        v-loading="loading"
        border
        stripe
        style="width: 100%"
      >
        <el-table-column prop="id" label="用户ID" width="80" align="center" />
        
        <el-table-column prop="nickName" label="用户昵称" min-width="120">
          <template #default="{ row }">
            <span v-if="row.nickName">{{ row.nickName }}</span>
            <span v-else class="text-gray">未设置</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="username" label="用户名" min-width="120">
          <template #default="{ row }">
            <span class="username">{{ row.username }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="phone" label="手机号" min-width="120">
          <template #default="{ row }">
            <span v-if="row.phone">{{ row.phone }}</span>
            <span v-else class="text-gray">未绑定</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="email" label="邮箱" min-width="160">
          <template #default="{ row }">
            <span v-if="row.email">{{ row.email }}</span>
            <span v-else class="text-gray">未绑定</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="createdAt" label="注册时间" min-width="160" align="center">
          <template #default="{ row }">
            <span>{{ formatDate(row.createdAt) }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="hasInvested" label="是否投资" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.hasInvested ? 'success' : 'info'">
              {{ row.hasInvested ? '已投资' : '未投资' }}
            </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column prop="investTime" label="投资时间" min-width="160" align="center">
          <template #default="{ row }">
            <span v-if="row.investTime">{{ formatDate(row.investTime) }}</span>
            <span v-else class="text-gray">--</span>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="gva-pagination">
        <el-pagination
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: 'MyReferrals'
})

import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getMyReferrals } from '@/api/invite'

const loading = ref(false)
const referralList = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const totalReferral = ref(0)

const searchForm = reactive({
  keyword: ''
})

// 计算已投资人数
const investedCount = computed(() => {
  return referralList.value.filter(item => item.hasInvested).length
})

// 计算转化率
const conversionRate = computed(() => {
  if (totalReferral.value === 0) return 0
  return ((investedCount.value / totalReferral.value) * 100).toFixed(1)
})

// 获取推荐列表
const fetchReferrals = async () => {
  loading.value = true
  try {
    const res = await getMyReferrals({
      page: page.value,
      pageSize: pageSize.value,
      keyword: searchForm.keyword
    })
    if (res.code === 0) {
      referralList.value = res.data.list
      total.value = res.data.total
      totalReferral.value = res.data.total
    }
  } catch (error) {
    ElMessage.error('获取推荐列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  page.value = 1
  fetchReferrals()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  page.value = 1
  fetchReferrals()
}

// 分页大小改变
const handleSizeChange = (val) => {
  pageSize.value = val
  fetchReferrals()
}

// 当前页改变
const handleCurrentChange = (val) => {
  page.value = val
  fetchReferrals()
}

// 格式化日期
const formatDate = (dateStr) => {
  if (!dateStr) return '--'
  const date = new Date(dateStr)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(() => {
  fetchReferrals()
})
</script>

<style scoped lang="scss">
.summary-row {
  margin-bottom: 20px;
  
  .card-header {
    font-weight: bold;
    color: #303133;
  }
  
  .summary-content {
    text-align: center;
    padding: 20px 0;
    
    .summary-number {
      font-size: 32px;
      font-weight: bold;
      color: #409EFF;
      margin-right: 5px;
    }
    
    .summary-unit {
      font-size: 14px;
      color: #909399;
    }
  }
}

.text-gray {
  color: #909399;
}

.username {
  font-family: 'Courier New', monospace;
  color: #606266;
}

.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>