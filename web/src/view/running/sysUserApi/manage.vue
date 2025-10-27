
<template>
  <div> 
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
        
            <el-table-column align="left" label="序列" prop="id" width="80" />

            <el-table-column align="left" label="添加时间" prop="createdAt" width="160">
              <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
            </el-table-column>
            
            <el-table-column align="left" label="用户UID" prop="userId" width="80" />

            <el-table-column align="left" label="交易所名称" prop="exchange" width="100" />

            <el-table-column align="left" label="api key" prop="key" width="180" />

            <el-table-column align="left" label="api Secret" prop="secret" width="120" />

            <el-table-column align="left" label="密码,如果有" prop="passwd" width="120" />

            <el-table-column align="left" label="api状态"   width="120" >
              <template #default="scope"> 
                        {{ scope.row.status == 1 ? '正常' : '错误' }} 
                </template>
            </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSysUserApiFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination" style="display: none;">
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

          <el-form  :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item style="display:none" label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
            <el-form-item style="display:none" label="添加时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item style="display:none" label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item style="display:none" label="逻辑删除时间:" prop="deletedAt">
    <el-date-picker v-model="formData.deletedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item style="display:none" label="用户UID:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户UID" />
</el-form-item>
            <el-form-item label="交易所名称:" prop="exchange">
    <!-- 中文注释：将交易所名称改为下拉选择，只允许币安与OKX，默认币安 -->
    <el-select v-model="formData.exchange" placeholder="请选择交易所" style="width:100%">
      <el-option label="binance" value="binance" />
      <el-option label="OKX" value="OKX" />
    </el-select>
</el-form-item>
            <el-form-item label="api key:" prop="key">
    <el-input v-model="formData.key" :clearable="true" placeholder="请输入api key" />
</el-form-item>
            <el-form-item label="api Secret:" prop="secret">
    <el-input v-model="formData.secret" :clearable="true" placeholder="请输入api Secret" />
</el-form-item>
            <el-form-item label="api密码(如果有):" prop="passwd">
    <el-input v-model="formData.passwd" :clearable="true" placeholder="请输入api密码，OKx等交易所有设置" />
</el-form-item>
            <el-form-item style="display:none" label="api状态: 1正常，2无权限，3错误:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入api状态: 1正常，2无权限，3错误" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="添加时间">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="更新时间">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="逻辑删除时间">
    {{ detailForm.deletedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="用户UID">
    {{ detailForm.userId }}
</el-descriptions-item>
                    <el-descriptions-item label="交易所名称">
    {{ detailForm.exchange }}
</el-descriptions-item>
                    <el-descriptions-item label="api key">
    {{ detailForm.key }}
</el-descriptions-item>
                    <el-descriptions-item label="api Secret">
    {{ detailForm.secret }}
</el-descriptions-item>
                    <el-descriptions-item label="api密码，OKx等交易所有设置">
    {{ detailForm.passwd }}
</el-descriptions-item>
                    <el-descriptions-item label="api状态: 1正常，2无权限，3错误">
    {{ detailForm.status }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSysUserApi,
  deleteSysUserApi,
  deleteSysUserApiByIds,
  updateSysUserApi,
  findSysUserApi,
  getUserApiList
} from '@/api/running/sysUserApi'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format' // 中文注释：修复导入的函数名错别字，使用正确的 filterDataSource，避免运行时报错
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"
import { useUserStore } from '@/pinia/modules/user' // 中文注释：引入用户状态，以便获取当前登录用户ID




defineOptions({
    // 中文注释：为避免与 sysUserApi.vue 的组件名重复而导致 keep-alive 复用缓存，这里使用唯一组件名
    name: 'SysUserApiManage'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()
const userStore = useUserStore() // 中文注释：获取用户store，用于读取当前登录用户信息

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(), 
            userId: userStore.userInfo.ID, // 中文注释：默认将当前登录用户的ID写入，用于创建时自动绑定
            exchange: 'binance', // 中文注释：默认选择“币安”
            key: '',
            secret: '',
            passwd: '',
            status: undefined,
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
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
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
  const table = await getUserApiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteSysUserApiFunc(row)
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
      const res = await deleteSysUserApiByIds({ ids })
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
const updateSysUserApiFunc = async(row) => {
    const res = await findSysUserApi({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSysUserApiFunc = async (row) => {
    const res = await deleteSysUserApi({ id: row.id })
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
    // 中文注释：打开创建弹窗时，确保userId为当前登录用户的ID
    formData.value.userId = userStore.userInfo.ID
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        createdAt: new Date(),
        updatedAt: new Date(),
        deletedAt: new Date(),
        userId: userStore.userInfo.ID, // 中文注释：重置时也保持为当前登录用户ID，避免提交为空
        exchange: '币安', // 中文注释：重置时默认“币安”
        key: '',
        secret: '',
        passwd: '',
        status: undefined,
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
                  res = await createSysUserApi(formData.value)
                  break
                case 'update':
                  res = await updateSysUserApi(formData.value)
                  break
                default:
                  res = await createSysUserApi(formData.value)
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
  const res = await findSysUserApi({ id: row.id })
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
