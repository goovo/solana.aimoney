
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item style="display: none;" > 
          <el-button  type="primary" icon="plus" @click="openDialog()">启动模拟交易</el-button>
          <el-button icon="refresh" @click="onReset">刷新</el-button>
        </el-form-item>
      </el-form>
      <!-- 新增：机器人进程展示位（显示：进程名、进程ID、进程状态、运行时间），进程ID来自后端 getDryRun 接口返回的 data -->
      <div class="process-panel" style="margin: 10px 0 0;">
        <div style="margin-left: 14px; margin-bottom: 10px; font-size: 14px; color: #0000FF;">
          <div><strong>交易机器人为您服务的前提：</strong></div>
          <div>1. 已经完成风险测评（低风险1倍杠杆，中等风险3倍杠杆，高风险5倍杠杆）；</div>
          <div>2. 配置好API并通过程序检测，确定正常可用,且API管理的资产不低于50美金；</div>
          <div>3. 完成对交易机器人授权；</div>
        </div>
        <el-descriptions title="交易机器人进程" :column="4" size="small" border>
          
          <el-descriptions-item label="进程名">
            {{ processNameText }}
          </el-descriptions-item>
          <el-descriptions-item label="进程ID">
            {{ dryRunPid || 0 }}
          </el-descriptions-item>
          <el-descriptions-item label="进程状态">
            <el-tag :type="dryRunPid > 0 ? 'success' : 'info'">{{ dryRunPid > 0 ? '运行中' : '未运行' }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="运行时间">
            {{ processUptimeText }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
      <!-- 新增结束 -->
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            
            <el-button  style="display: inline-block;"  >交易报表如下</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="序列" prop="id" width="60" />

            <el-table-column align="left" label="交易所" prop="exchange" width="90" />

            <el-table-column align="left" label="市场符号" prop="pair" width="150" />

            <el-table-column align="left" label="状态"  width="120" >
                <template #default="scope"> 
                        {{ scope.row.isOpen === 1 ? '持仓中' : '已平仓' }} 
                </template>
            </el-table-column>

            <el-table-column align="left" label="开仓时间" prop="openDate" width="180">
                <template #default="scope">{{ formatDate(scope.row.openDate) }}</template>
            </el-table-column>
            <el-table-column align="left" label="实际开仓价" prop="openRate" width="120" /> 

            <el-table-column align="left" label="开仓价值" prop="openTradeValue" width="120" />

            <el-table-column align="left" label="平仓时间" prop="closeDate" width="180">
                <template #default="scope">{{ formatDate(scope.row.closeDate) }}</template>
            </el-table-column>
            <el-table-column align="left" label="实际平仓价" prop="closeRate" width="120" /> 

            <el-table-column align="left" label="净利润" prop="realizedProfit" width="120" /> 

            <el-table-column align="left" label="策略名称" prop="strategy" width="150" /> 

            <el-table-column align="left" label="K线周期" prop="timeframe" width="120" /> 

            <el-table-column align="left" label="杠杆倍数" prop="leverage" width="120" />

            <el-table-column align="left" label="方向" width="120" >
            <template #default="scope"> 
                        {{ scope.row.isShort == 1 ? '做空' : '做多' }} 
                </template>
            </el-table-column> 

            <el-table-column align="left" label="累计资金费用" prop="fundingFees" width="120" /> 

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" style="display: none;" @click="updateTradesFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" style="display: none;"  @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'配置机器人':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">启 动</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            
            <el-form-item label="交易所:" prop="exchange">
    <!-- 将原先的输入框改为下拉选择框：从 apisData 列表中遍历出所有可用的交易所，并在选项中展示状态 -->
    <el-select v-model="formData.exchange" placeholder="请选择交易所" style="width:100%" filterable :clearable="true">
      <!-- 说明：apisData 由 getApisData() 从后端接口 getUserApiList 拉取，单条记录包含 exchange、status 等字段 -->
      <!-- 状态映射：1=正常，2=无权限，3=错误，标签文本通过 statusLabel() 生成 -->
      <el-option
        v-for="item in apisData"
        :key="item.id || item.exchange"
        :label="`${item.exchange}（${statusLabel(item.status)}）`"
        :value="item.exchange"
      />
    </el-select>
 </el-form-item>
         
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="序列">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="交易所">
    {{ detailForm.exchange }}
</el-descriptions-item>
                    <el-descriptions-item label="市场符号">
    {{ detailForm.pair }}
</el-descriptions-item>
                    <el-descriptions-item label="左侧代码">
    {{ detailForm.baseCurrency }}
</el-descriptions-item>
                    <el-descriptions-item label="右侧代码">
    {{ detailForm.stakeCurrency }}
</el-descriptions-item>
                    <el-descriptions-item label="状态">
    {{ detailForm.isOpen }}
</el-descriptions-item>
                    <el-descriptions-item label="开仓费率">
    {{ detailForm.feeOpen }}
</el-descriptions-item>
                    <el-descriptions-item label="开仓手续费">
    {{ detailForm.feeOpenCost }}
</el-descriptions-item>
                    <el-descriptions-item label="开仓手续费币种">
    {{ detailForm.feeOpenCurrency }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓费率">
    {{ detailForm.feeClose }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓手续费">
    {{ detailForm.feeCloseCost }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓手续费币种">
    {{ detailForm.feeCloseCurrency }}
</el-descriptions-item>
                    <el-descriptions-item label="实际开仓价">
    {{ detailForm.openRate }}
</el-descriptions-item>
                    <el-descriptions-item label="限价单指定的开仓价">
    {{ detailForm.openRateRequested }}
</el-descriptions-item>
                    <el-descriptions-item label="开仓价值">
    {{ detailForm.openTradeValue }}
</el-descriptions-item>
                    <el-descriptions-item label="实际平仓价">
    {{ detailForm.closeRate }}
</el-descriptions-item>
                    <el-descriptions-item label="限价平仓价">
    {{ detailForm.closeRateRequested }}
</el-descriptions-item>
                    <el-descriptions-item label="净利润">
    {{ detailForm.realizedProfit }}
</el-descriptions-item>
                    <el-descriptions-item label="收益百分比">
    {{ detailForm.closeProfit }}
</el-descriptions-item>
                    <el-descriptions-item label="净利润">
    {{ detailForm.closeProfitAbs }}
</el-descriptions-item>
                    <el-descriptions-item label="建仓本金(x杠杆)">
    {{ detailForm.stakeAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="最大本金">
    {{ detailForm.maxStakeAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="Base数量(x杠杆)">
    {{ detailForm.amount }}
</el-descriptions-item>
                    <el-descriptions-item label="原始下单数量">
    {{ detailForm.amountRequested }}
</el-descriptions-item>
                    <el-descriptions-item label="开仓时间">
    {{ detailForm.openDate }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓时间">
    {{ detailForm.closeDate }}
</el-descriptions-item>
                    <el-descriptions-item label="止损触发价">
    {{ detailForm.stopLoss }}
</el-descriptions-item>
                    <el-descriptions-item label="止损相对开仓价的百分比">
    {{ detailForm.stopLossPct }}
</el-descriptions-item>
                    <el-descriptions-item label="硬止损价">
    {{ detailForm.initialStopLoss }}
</el-descriptions-item>
                    <el-descriptions-item label="硬止损百分比">
    {{ detailForm.initialStopLossPct }}
</el-descriptions-item>
                    <el-descriptions-item label="移动止损">
    {{ detailForm.isStopLossTrailing }}
</el-descriptions-item>
                    <el-descriptions-item label="最高价">
    {{ detailForm.maxRate }}
</el-descriptions-item>
                    <el-descriptions-item label="最低价">
    {{ detailForm.minRate }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓原因">
    {{ detailForm.exitReason }}
</el-descriptions-item>
                    <el-descriptions-item label="平仓状态">
    {{ detailForm.exitOrderStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="策略名称">
    {{ detailForm.strategy }}
</el-descriptions-item>
                    <el-descriptions-item label="信号标签">
    {{ detailForm.enterTag }}
</el-descriptions-item>
                    <el-descriptions-item label=" K 线周期">
    {{ detailForm.timeframe }}
</el-descriptions-item>
                    <el-descriptions-item label="交易模式">
    {{ detailForm.tradingMode }}
</el-descriptions-item>
                    <el-descriptions-item label="最小数量">
    {{ detailForm.amountPrecision }}
</el-descriptions-item>
                    <el-descriptions-item label="价格步长">
    {{ detailForm.pricePrecision }}
</el-descriptions-item>
                    <el-descriptions-item label="精度">
    {{ detailForm.precisionMode }}
</el-descriptions-item>
                    <el-descriptions-item label="精度值">
    {{ detailForm.precisionModePrice }}
</el-descriptions-item>
                    <el-descriptions-item label="合约乘数">
    {{ detailForm.contractSize }}
</el-descriptions-item>
                    <el-descriptions-item label="杠杆倍数">
    {{ detailForm.leverage }}
</el-descriptions-item>
                    <el-descriptions-item label="0多１空">
    {{ detailForm.isShort }}
</el-descriptions-item>
                    <el-descriptions-item label="预估强平价">
    {{ detailForm.liquidationPrice }}
</el-descriptions-item>
                    <el-descriptions-item label="资金费率快照">
    {{ detailForm.interestRate }}
</el-descriptions-item>
                    <el-descriptions-item label="累计资金费用">
    {{ detailForm.fundingFees }}
</el-descriptions-item>
                    <el-descriptions-item label="预估资金费用">
    {{ detailForm.fundingFeeRunning }}
</el-descriptions-item>
                    <el-descriptions-item label="行级乐观锁">
    {{ detailForm.recordVersion }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createTrades,
  updateTrades,
  findTrades,
  getTradesList,
  createDryRun,
  getDryRun
} from '@/api/running/trades'

import {
  getUserApiList
} from '@/api/running/sysUserApi'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted, onActivated } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'TradesUser'
})

const apisData = ref([])
// 查询：用户 API 列表
const getApisData = async() => {
  // 说明：此处依赖 page/pageSize/searchInfo，因此不要在它们声明之前调用
  const apis = await getUserApiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (apis.code === 0) {
    apisData.value = apis.data.list 
  }
}
// 注意：移除此处的立即调用，避免在 page/pageSize/searchInfo 尚未初始化时触发，导致报错
// getApisData()

// 状态码到中文文案的映射函数
// 1=正常，2=无权限，3=错误；其余情况统一显示为“未知”
const statusLabel = (status) => {
  // 保持健壮性：后端若返回字符串类型也能兼容
  const s = Number(status)
  if (s === 1) return 'API可用'
  if (s === 2) return 'API无权限'
  if (s === 3) return 'API错误'
  return 'API状态未知'
}

const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 新增：进程信息相关的响应式变量
const dryRunPid = ref(0)                 // 进程ID，来源于后端 /trades/getDryRun 接口的 data
const processNameText = ref('我的Ai交易机器人') // 进程名，当前后端未返回，前端固定显示
const processUptimeText = ref('-')       // 运行时间，后端未提供，先展示占位符

// 新增：封装拉取 dry-run 进程信息的方法
const loadDryRun = async () => {
  // 中文注释：调用后端接口获取当前用户 dry-run 进程PID
  const res = await getDryRun()
  if (res && res.code === 0) {
    dryRunPid.value = Number(res.data) || 0
  } else {
    dryRunPid.value = 0
  }
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            exchange: '',
            pair: '',
            baseCurrency: '',
            stakeCurrency: '',
            isOpen: undefined,
            feeOpen: 0,
            feeOpenCost: 0,
            feeOpenCurrency: '',
            feeClose: 0,
            feeCloseCost: 0,
            feeCloseCurrency: '',
            openRate: 0,
            openRateRequested: 0,
            openTradeValue: 0,
            closeRate: 0,
            closeRateRequested: 0,
            realizedProfit: 0,
            closeProfit: 0,
            closeProfitAbs: 0,
            stakeAmount: 0,
            maxStakeAmount: 0,
            amount: 0,
            amountRequested: 0,
            openDate: new Date(),
            closeDate: new Date(),
            stopLoss: 0,
            stopLossPct: 0,
            initialStopLoss: 0,
            initialStopLossPct: 0,
            isStopLossTrailing: false,
            maxRate: 0,
            minRate: 0,
            exitReason: '',
            exitOrderStatus: '',
            strategy: '',
            enterTag: '',
            timeframe: undefined,
            tradingMode: null,
            amountPrecision: 0,
            pricePrecision: 0,
            precisionMode: undefined,
            precisionModePrice: undefined,
            contractSize: 0,
            leverage: 0,
            isShort: undefined,
            liquidationPrice: 0,
            interestRate: 0,
            fundingFees: 0,
            fundingFeeRunning: 0,
            recordVersion: undefined,
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ==========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
  // 新增：刷新一次进程展示信息
  loadDryRun()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.isStopLossTrailing === ""){
        searchInfo.value.isStopLossTrailing=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTradesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
// 原先直接调用后端方法 getDryRun()，现改为封装方法，便于统一更新展示位
loadDryRun()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 页面挂载与激活时获取 API 列表（解决首次加载未执行问题，并兼容 keep-alive 返回场景）
onMounted(() => {
  // 中文注释：页面首次挂载时请求一次用户 API 列表，用于交易所下拉选项
  getApisData()
  // 中文注释：同时拉取一次进程信息用于展示
  loadDryRun()
})

onActivated(() => {
  // 中文注释：当页面被 keep-alive 缓存并再次激活时，刷新一次数据，保证下拉选项为最新
  getApisData()
  // 中文注释：同时刷新进程信息
  loadDryRun()
})


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}


// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTradesFunc = async(row) => {
    const res = await findTrades({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}



// 弹窗控制标记
const dialogFormVisible = ref(false)

// 为当前用户创建模拟交易机器人
const openDialog = async () => {
    // 中文注释：点击“启动模拟交易”时，先检查当前是否已有 dry-run 进程
    const res = await getDryRun()
    const currentPid = (res && res.code === 0) ? Number(res.data) || 0 : 0
    if (currentPid > 0) {
      // 已经存在进程，不允许重复创建
      ElMessage.warning('每人只能创建一个模拟交易机器人')
      // 同步展示位
      dryRunPid.value = currentPid
      return
    }
    // 不存在进程，允许创建
    const createRes = await createDryRun()
    if (createRes && createRes.code === 0) {
      ElMessage.success('已为您启动模拟交易机器人')
      // 刷新展示位
      loadDryRun()
    }
    // 旧的弹窗逻辑保留（目前未启用表单）
    // type.value = 'create'
    // dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        exchange: '',
        pair: '',
        baseCurrency: '',
        stakeCurrency: '',
        isOpen: undefined,
        feeOpen: 0,
        feeOpenCost: 0,
        feeOpenCurrency: '',
        feeClose: 0,
        feeCloseCost: 0,
        feeCloseCurrency: '',
        openRate: 0,
        openRateRequested: 0,
        openTradeValue: 0,
        closeRate: 0,
        closeRateRequested: 0,
        realizedProfit: 0,
        closeProfit: 0,
        closeProfitAbs: 0,
        stakeAmount: 0,
        maxStakeAmount: 0,
        amount: 0,
        amountRequested: 0,
        openDate: new Date(),
        closeDate: new Date(),
        stopLoss: 0,
        stopLossPct: 0,
        initialStopLoss: 0,
        initialStopLossPct: 0,
        isStopLossTrailing: false,
        maxRate: 0,
        minRate: 0,
        exitReason: '',
        exitOrderStatus: '',
        strategy: '',
        enterTag: '',
        timeframe: undefined,
        tradingMode: null,
        amountPrecision: 0,
        pricePrecision: 0,
        precisionMode: undefined,
        precisionModePrice: undefined,
        contractSize: 0,
        leverage: 0,
        isShort: undefined,
        liquidationPrice: 0,
        interestRate: 0,
        fundingFees: 0,
        fundingFeeRunning: 0,
        recordVersion: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createTrades(formData.value)
                  break
                case 'update':
                  res = await updateTrades(formData.value)
                  break
                default:
                  res = await createTrades(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTrades({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
