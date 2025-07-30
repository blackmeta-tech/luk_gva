<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="调用目标字符串">
          <el-input v-model="searchInfo.invokeTarget" placeholder="请输入调用目标字符串" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="定时任务执行字符串" prop="invokeTarget"/>
        <el-table-column align="left" label="执行状态" prop="status">
          <template #default="scope">
            <el-tag v-if="scope.row.status == 1" type="success">成功</el-tag>
            <el-tag v-if="scope.row.status == 2" type="danger">失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" />
        <el-table-column align="left" label="执行时间" prop="createdAt" width="300" >
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createdAt)) }}
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
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="createdBy字段:">
          <el-input v-model.number="formData.createdBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="deletedBy字段:">
          <el-input v-model.number="formData.deletedBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="invokeTarget字段:">
          <el-input v-model="formData.invokeTarget" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="remark字段:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="status字段:">
          <el-input v-model.number="formData.status" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="updatedBy字段:">
          <el-input v-model.number="formData.updatedBy" clearable placeholder="请输入" />
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
export default {
  name: 'LukScheduledTaskLog'
}
</script>

<script setup>
import {
  createLukScheduledTaskLog,
  updateLukScheduledTaskLog,
  getLukScheduledTaskLogList
} from '@/api/lukScheduledTaskLog'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  createdBy: 0,
  deletedBy: 0,
  invokeTarget: '',
  remark: '',
  status: 0,
  updatedBy: 0,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
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
  const table = await getLukScheduledTaskLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    createdBy: 0,
    deletedBy: 0,
    invokeTarget: '',
    remark: '',
    status: 0,
    updatedBy: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createLukScheduledTaskLog(formData.value)
      break
    case 'update':
      res = await updateLukScheduledTaskLog(formData.value)
      break
    default:
      res = await createLukScheduledTaskLog(formData.value)
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
</script>

<style>
</style>
