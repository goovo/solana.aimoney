
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
                        {{ scope.row.isOpen.Int === 1 ? '持仓中' : '已平仓' }}
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
                        {{ scope.row.isShort.Int == 1 ? '做空' : '做多' }}
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
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="序列:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入序列" />
</el-form-item>
            <el-form-item label="交易所:" prop="exchange">
    <el-input v-model="formData.exchange" :clearable="true" placeholder="请输入交易所" />
</el-form-item>
            <el-form-item label="市场符号:" prop="pair">
    <el-input v-model="formData.pair" :clearable="true" placeholder="请输入市场符号" />
</el-form-item>
            <el-form-item label="左侧代码:" prop="baseCurrency">
    <el-input v-model="formData.baseCurrency" :clearable="true" placeholder="请输入左侧代码" />
</el-form-item>
            <el-form-item label="右侧代码:" prop="stakeCurrency">
    <el-input v-model="formData.stakeCurrency" :clearable="true" placeholder="请输入右侧代码" />
</el-form-item>
            <el-form-item label="状态:" prop="isOpen">
    <el-input v-model.number="formData.isOpen" :clearable="true" placeholder="请输入状态" />
</el-form-item>
            <el-form-item label="开仓费率:" prop="feeOpen">
    <el-input-number v-model="formData.feeOpen" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="开仓手续费:" prop="feeOpenCost">
    <el-input-number v-model="formData.feeOpenCost" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="开仓手续费币种:" prop="feeOpenCurrency">
    <el-input v-model="formData.feeOpenCurrency" :clearable="true" placeholder="请输入开仓手续费币种" />
</el-form-item>
            <el-form-item label="平仓费率:" prop="feeClose">
    <el-input-number v-model="formData.feeClose" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="平仓手续费:" prop="feeCloseCost">
    <el-input-number v-model="formData.feeCloseCost" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="平仓手续费币种:" prop="feeCloseCurrency">
    <el-input v-model="formData.feeCloseCurrency" :clearable="true" placeholder="请输入平仓手续费币种" />
</el-form-item>
            <el-form-item label="实际开仓价:" prop="openRate">
    <el-input-number v-model="formData.openRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="限价单指定的开仓价:" prop="openRateRequested">
    <el-input-number v-model="formData.openRateRequested" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="开仓价值:" prop="openTradeValue">
    <el-input-number v-model="formData.openTradeValue" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="实际平仓价:" prop="closeRate">
    <el-input-number v-model="formData.closeRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="限价平仓价:" prop="closeRateRequested">
    <el-input-number v-model="formData.closeRateRequested" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="净利润:" prop="realizedProfit">
    <el-input-number v-model="formData.realizedProfit" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="收益百分比:" prop="closeProfit">
    <el-input-number v-model="formData.closeProfit" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="净利润:" prop="closeProfitAbs">
    <el-input-number v-model="formData.closeProfitAbs" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="建仓本金(x杠杆):" prop="stakeAmount">
    <el-input-number v-model="formData.stakeAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="最大本金:" prop="maxStakeAmount">
    <el-input-number v-model="formData.maxStakeAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="Base数量(x杠杆):" prop="amount">
    <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="原始下单数量:" prop="amountRequested">
    <el-input-number v-model="formData.amountRequested" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="开仓时间:" prop="openDate">
    <el-date-picker v-model="formData.openDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="平仓时间:" prop="closeDate">
    <el-date-picker v-model="formData.closeDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="止损触发价:" prop="stopLoss">
    <el-input-number v-model="formData.stopLoss" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="止损相对开仓价的百分比:" prop="stopLossPct">
    <el-input-number v-model="formData.stopLossPct" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="硬止损价:" prop="initialStopLoss">
    <el-input-number v-model="formData.initialStopLoss" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="硬止损百分比:" prop="initialStopLossPct">
    <el-input-number v-model="formData.initialStopLossPct" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="移动止损:" prop="isStopLossTrailing">
    <el-switch v-model="formData.isStopLossTrailing" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="最高价:" prop="maxRate">
    <el-input-number v-model="formData.maxRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="最低价:" prop="minRate">
    <el-input-number v-model="formData.minRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="平仓原因:" prop="exitReason">
    <el-input v-model="formData.exitReason" :clearable="true" placeholder="请输入平仓原因" />
</el-form-item>
            <el-form-item label="平仓状态:" prop="exitOrderStatus">
    <el-input v-model="formData.exitOrderStatus" :clearable="true" placeholder="请输入平仓状态" />
</el-form-item>
            <el-form-item label="策略名称:" prop="strategy">
    <el-input v-model="formData.strategy" :clearable="true" placeholder="请输入策略名称" />
</el-form-item>
            <el-form-item label="信号标签:" prop="enterTag">
    <el-input v-model="formData.enterTag" :clearable="true" placeholder="请输入信号标签" />
</el-form-item>
            <el-form-item label=" K 线周期:" prop="timeframe">
    <el-input v-model.number="formData.timeframe" :clearable="true" placeholder="请输入 K 线周期" />
</el-form-item>
            <el-form-item label="交易模式:" prop="tradingMode">
    <el-select v-model="formData.tradingMode" placeholder="请选择交易模式" style="width:100%" filterable :clearable="true">
       <el-option v-for="item in ['SPOT', 'MARGIN', 'FUTURES' ]" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
            <el-form-item label="最小数量:" prop="amountPrecision">
    <el-input-number v-model="formData.amountPrecision" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="价格步长:" prop="pricePrecision">
    <el-input-number v-model="formData.pricePrecision" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="精度:" prop="precisionMode">
    <el-input v-model.number="formData.precisionMode" :clearable="true" placeholder="请输入精度" />
</el-form-item>
            <el-form-item label="精度值:" prop="precisionModePrice">
    <el-input v-model.number="formData.precisionModePrice" :clearable="true" placeholder="请输入精度值" />
</el-form-item>
            <el-form-item label="合约乘数:" prop="contractSize">
    <el-input-number v-model="formData.contractSize" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="杠杆倍数:" prop="leverage">
    <el-input-number v-model="formData.leverage" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="0多１空:" prop="isShort">
    <el-input v-model.number="formData.isShort" :clearable="true" placeholder="请输入0多１空" />
</el-form-item>
            <el-form-item label="预估强平价:" prop="liquidationPrice">
    <el-input-number v-model="formData.liquidationPrice" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="资金费率快照:" prop="interestRate">
    <el-input-number v-model="formData.interestRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="累计资金费用:" prop="fundingFees">
    <el-input-number v-model="formData.fundingFees" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="预估资金费用:" prop="fundingFeeRunning">
    <el-input-number v-model="formData.fundingFeeRunning" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
            <el-form-item label="行级乐观锁:" prop="recordVersion">
    <el-input v-model.number="formData.recordVersion" :clearable="true" placeholder="请输入行级乐观锁" />
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
  deleteTrades,
  deleteTradesByIds,
  updateTrades,
  findTrades,
  getTradesList
} from '@/api/running/trades'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Trades'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
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

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTradesFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteTradesByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
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


// 删除行
const deleteTradesFunc = async (row) => {
    const res = await deleteTrades({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
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
