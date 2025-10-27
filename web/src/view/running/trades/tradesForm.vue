
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTrades,
  updateTrades,
  findTrades
} from '@/api/running/trades'

defineOptions({
    name: 'TradesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTrades({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
