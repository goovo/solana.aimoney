
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="key字段:" prop="key">
    <el-input v-model="formData.key" :clearable="true" placeholder="请输入key字段" />
</el-form-item>
        <el-form-item label="valueType字段:" prop="valueType">
    <el-input v-model="formData.valueType" :clearable="true" placeholder="请输入valueType字段" />
</el-form-item>
        <el-form-item label="stringValue字段:" prop="stringValue">
    <el-input v-model="formData.stringValue" :clearable="true" placeholder="请输入stringValue字段" />
</el-form-item>
        <el-form-item label="datetimeValue字段:" prop="datetimeValue">
    <el-date-picker v-model="formData.datetimeValue" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="floatValue字段:" prop="floatValue">
    <el-input-number v-model="formData.floatValue" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="intValue字段:" prop="intValue">
    <el-input v-model.number="formData.intValue" :clearable="true" placeholder="请输入intValue字段" />
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
  createKeyvaluestore,
  updateKeyvaluestore,
  findKeyvaluestore
} from '@/api/strategy/keyvaluestore'

defineOptions({
    name: 'KeyvaluestoreForm'
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
            key: '',
            valueType: '',
            stringValue: '',
            datetimeValue: new Date(),
            floatValue: 0,
            intValue: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findKeyvaluestore({ ID: route.query.id })
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
               res = await createKeyvaluestore(formData.value)
               break
             case 'update':
               res = await updateKeyvaluestore(formData.value)
               break
             default:
               res = await createKeyvaluestore(formData.value)
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
