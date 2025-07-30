<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
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
          <el-button size="small" type="success" icon="plus" @click="openDialog">新增</el-button>
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
        <el-table-column align="left" label="id" prop="id" />
        <el-table-column align="left" label="图片" >
          <template #default="scope">
            <el-image style="width: 100px; height: 100px" :src="scope.row.pic"></el-image>
          </template>
        </el-table-column>
        <el-table-column align="left" label="链接" prop="link"/>
        <el-table-column align="left" label="是否启用" >
          <template #default="scope">
            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-color="#13ce66" inactive-color="#ff4949" @change="switchDialog(scope.row,  scope.row.status)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateLukBannerFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-form-item>
          <template #>
            <div>
              <uploader-pic @updatePic="(res)=>formData.pic=res" @updatePicKey="(res)=>formData.picKey=res" :pic="formData.pic" :picKey="formData.picKey"></uploader-pic>
              <p style="child-align: bottom">图片建议尺寸：694x340px</p>
            </div>
          </template>
        </el-form-item>
        <el-form-item label="链接:">
          <el-input v-model="formData.link" clearable placeholder="请输入" />
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
import uploaderPic from '@/view/lukNft/uploaderPic.vue'
export default {
  name: 'LukBanner',
  components: {
    'uploader-pic': uploaderPic
  },
}
</script>

<script setup>
import {
  createLukBanner,
  deleteLukBanner,
  updateLukBanner,
  findLukBanner,
  getLukBannerList
} from '@/api/lukBanner'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  link: '',
  pic: '',
  status: 1,
  picKey: '',
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
  const table = await getLukBannerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteLukBannerFunc(row)
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukBannerFunc = async(row) => {
  const res = await findLukBanner({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rewLukBanner
    dialogFormVisible.value = true
    uploaderPic.data().lastFile.url = formData.value.pic
    uploaderPic.data().lastFile.key = formData.value.picKey
  }
}


// 删除行
const deleteLukBannerFunc = async (row) => {
  const res = await deleteLukBanner({ id: row.id })
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
    link: '',
    pic: '',
    status: 1,
    picKey: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.pic === '') {
    ElMessage.error('图片不得为空!')
    return
  }
  let res
  switch (type.value) {
    case 'create':
      res = await createLukBanner(formData.value)
      break
    case 'update':
      res = await updateLukBanner(formData.value)
      break
    default:
      res = await createLukBanner(formData.value)
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
const switchDialog = async(values, num) => {
  let res
  values.status = num
  res = await updateLukBanner(values)
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
