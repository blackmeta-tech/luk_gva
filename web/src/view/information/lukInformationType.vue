<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item label="是否启用">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择">
              <el-option
                  v-for="item in statusOptions"
                  :key="item.value"
                  :label="`${item.label}`"
                  :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
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
        <el-table-column align="left" label="id" prop="id" />
        <el-table-column align="left" label="分类(中文)" prop="nameCh" />
        <el-table-column align="left" label="分类(英文)" prop="nameEn" />
        <el-table-column align="left" label="排序序号" prop="sortNum" />
        <el-table-column align="left" label="是否启用" >
          <template #default="scope">
            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-color="#13ce66" inactive-color="#ff4949" @change="switchDialog(scope.row, 'status', scope.row.status)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateLukInformationTypeFunc(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="150px">
        <el-form-item label="分类(中文):">
          <el-input v-model="formData.nameCh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类(英文):">
          <el-input v-model="formData.nameEn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序序号:">
          <template #>
            <el-input-number v-model="formData.sortNum"  label="排序序号"></el-input-number>
          </template>

        </el-form-item>
        <el-form-item label="是否启用:">
          <template #>
            <el-switch v-model="formData.status" :active-value="1" :inactive-value="2" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
          </template>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
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
  name: 'LukInformationType',
  data() {
    return {
      value: true
    }
  }
}
</script>

<script setup>
import {
  createLukInformationType,
  deleteLukInformationType,
  deleteLukInformationTypeByIds,
  updateLukInformationType,
  findLukInformationType,
  getLukInformationTypeList
} from '@/api/lukInformationType'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        nameCh: '',
        nameEn: '',
        sortNum: 0,
        status: 1,
        remark: '',
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const statusOptions = ref([])

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
  const table = await getLukInformationTypeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  statusOptions.value = [
    {
      value: '1',
      label: '启用',
      type: 'success'
    },
    {
      value: '2',
      label: '停用',
      type: ''
    }
  ]
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
            deleteLukInformationTypeFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
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
      const res = await deleteLukInformationTypeByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukInformationTypeFunc = async(row) => {
    const res = await findLukInformationType({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.relukDappInformationType
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLukInformationTypeFunc = async (row) => {
    const res = await deleteLukInformationType({ id: row.id })
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
        nameCh: '',
        nameEn: '',
        sortNum: 0,
        status: 1,
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createLukInformationType(formData.value)
          break
        case 'update':
          res = await updateLukInformationType(formData.value)
          break
        default:
          res = await createLukInformationType(formData.value)
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

//开关控制
const switchDialog = async(values, type, num) => {
  let res
  switch (type) {
    case 'status':
      values.status = num
      break
  }
  res = await updateLukInformationType(values)
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
