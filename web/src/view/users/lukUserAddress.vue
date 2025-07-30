<template>
<div>
  <div class="gva-search-box">
    <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="用户地址">
        <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item>
        <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
        <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
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
        >
        <el-table-column align="left" label="更新日期" width="180">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.updatedAt)) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户地址" prop="address" width="350" />
        <el-table-column align="left" label="零手续费白名单" prop="isCharge">
          <template #default="scope">
            <el-tag v-show="scope.row.isCharge == 1" type="success">是</el-tag>
            <el-tag v-show="scope.row.isCharge == 0" type="danger">否</el-tag>
            <el-popover v-model="scope.row.visible" placement="top">
              <p>确定取消白名单吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="updateLukUserWalletFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button v-show="scope.row.isCharge == 1" size="small" type="primary" link icon="delete" @click="scope.row.visible = true">取消白名单</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" />
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
  <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作" width="400px">
    <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
      <el-form-item label="用户地址:"  prop="address" >
        <el-input v-model="formData.address" :clearable="true"  placeholder="请输入" />
      </el-form-item>
      <el-form-item label="备注:"  prop="remark" >
        <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入" />
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
  name: 'LukUserAddress'
}
</script>

<script setup>
import { getLukUserAddressList, updateLukUserAddress } from '@/api/lukUseAddress'
// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  address: null,
  type: 1,
  remark: null,
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
  searchInfo.value = {
    stop: 1,
  }
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
  const table = await getLukUserAddressList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async() =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 清退
const updateLukUserWalletFunc = async(row) => {
  formData.value.type = 0
  formData.value.address = row.address
  formData.value.remark = row.remark
  const res = await updateLukUserAddress(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
    formData.value.address = null
    formData.value.remark = null
  }
}

// 打开弹窗
const openDialog = () => {
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value.address = null
  formData.value.remark = null
}
// 弹窗确定
const enterDialog = async() => {
  formData.value.type = 1
  var res = await updateLukUserAddress(formData.value)
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
