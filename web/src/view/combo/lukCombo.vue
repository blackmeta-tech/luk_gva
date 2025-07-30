<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
        <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
         —
        <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
<!--          <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        >
        <el-table-column align="left" label="图片" width="150" >
          <template #default="scope">
            <el-image style="width: 100px; height: 100px" :src="scope.row.pic"></el-image>
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="套餐天数" prop="days" width="120">
            <template #default="scope">
            {{ scope.row.days }}天套餐
            </template>
        </el-table-column>
        <el-table-column align="left" label="今日总额度[USDT]" prop="limit"  width="150" />
        <el-table-column align="left" label="达标业绩[USDT]" prop="performance" width="130" />
        <el-table-column align="left" label="最低购入价格[USDT]" prop="priceMin" width="160" />
        <el-table-column align="left" label="最高购入价格[USDT]" prop="priceMax" width="160"/>
        <el-table-column align="left" label="收益率[%]" prop="rate" width="100" />
        <el-table-column align="left" label="上架日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="上架状态" >
          <template #default="scope">
            <el-switch v-model="scope.row.status" :active-value="0" :inactive-value="-1" active-color="#13ce66" inactive-color="#ff4949" @change="switchDialog(scope.row,  scope.row.status)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateLukComboFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" label-width="120px">
        <el-form-item>
          <template #>
            <uploader-pic @updatePic="(res)=>formData.pic=res" @updatePicKey="(res)=>formData.picKey=res" :pic="formData.pic" :picKey="formData.picKey"></uploader-pic>
          </template>
        </el-form-item>
        <el-form-item v-show="type == 'create'" label="套餐天数:"  prop="days" class="required">
          <el-input-number v-model.number="formData.days" style="width:60%" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="达标业绩[USDT]:"  prop="performance" class="required">
          <el-input-number v-model.number="formData.performance" style="width:60%" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最低购入价格:"  prop="priceMin" class="required">
          <el-input-number v-model="formData.priceMin"  style="width:60%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="最高购入价格:"  prop="priceMax" class="required">
          <el-input-number v-model="formData.priceMax"  style="width:60%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="收益率百分比:"  prop="rate" class="required">
          <el-input-number v-model.number="formData.rate" style="width:60%" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="今日总额度:"  prop="limit" class="required">
          <el-input-number v-model.number="formData.limit" style="width:60%" :precision="2"  :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import uploaderPic from '@/view/lukNft/uploaderPic.vue'
export default {
  name: 'LukCombo',
  components: {
    'uploader-pic': uploaderPic
  },
}
</script>

<script setup>
import {
  createLukCombo,
  updateLukCombo,
  findLukCombo,
  getLukComboList
} from '@/api/lukCombo'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  days: 0,
  limit: 0,
  pic: '',
  picKey: '',
  priceMax: 0,
  priceMin: 0,
  rate: 0,
  performance: 0,
})

const elFormRef = ref()

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
  page.value = 1
  pageSize.value = 10
  getTableData()
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
  const table = await getLukComboList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukComboFunc = async(row) => {
    const res = await findLukCombo({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.relukCombo
        dialogFormVisible.value = true
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
    days: 0,
    limit: 0,
    pic: '',
    picKey: '',
    priceMax: 0,
    priceMin: 0,
    rate: 0,
    performance: 0,
  }
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.days <= 0 || formData.value.priceMax <= 0 || formData.value.priceMin <= 0 || formData.value.rate <= 0 || formData.value.limit <= 0 || formData.value.pic == ''){
    ElMessage.error('全部参数为必填项')
    return
  }
  let res
  switch (type.value) {
    case 'create':
      res = await createLukCombo(formData.value)
      break
    case 'update':
      res = await updateLukCombo(formData.value)
      break
    default:
      res = await createLukCombo(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}
// 开关控制
const switchDialog = async(values, num) => {
  let res
  values.status = num
  res = await updateLukCombo(values)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
    })
    getTableData()
  }
}
</script>

<style>
</style>
