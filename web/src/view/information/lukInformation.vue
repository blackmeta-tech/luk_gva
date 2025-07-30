<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="分类">
          <el-select v-model="searchInfo.typeId" clearable placeholder="请选择">
            <el-option
                v-for="item in typesOptions"
                :key="item.id"
                :label="`${item.nameCh}`"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
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
        <el-table-column align="left" label="图片" width="150">
          <template #default="scope">
            <el-image v-show="scope.row.pic != ''" style="width: 140px; " :src="scope.row.pic"></el-image>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标题(中文)" prop="titleCh" width="200" :show-overflow-tooltip='true' />
        <el-table-column align="left" label="标题(英文)" prop="titleEn" width="200" :show-overflow-tooltip='true' />
        <el-table-column align="left" label="内容(中文)" prop="articleCh" width="200">
          <template #default="scope">
            <el-tooltip class="item" effect="dark" placement="top" :content="scope.row.articleCh" raw-content>
              <div class="online">{{scope.row.articleCh}}</div>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="left" label="内容(英文)" prop="articleEn" width="200">
          <template #default="scope">
            <el-tooltip class="item" effect="dark" placement="top" :content="scope.row.articleEn" raw-content>
              <div class="online">{{scope.row.articleEn}}</div>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="left" label="分类" prop="typeName" width="80" />
        <el-table-column align="left" label="是否启用">
          <template #default="scope">
            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-color="#13ce66" inactive-color="#ff4949" @change="switchDialog(scope.row, scope.row.status)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="更新日期" width="180">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.timeline)) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
          <el-button type="text" icon="edit" size="small" class="table-button" @click="updateLukInformationFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" z-index="100" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="150px">
        <el-form-item label="标题(中文):">
          <el-input v-model="formData.titleCh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题(英文):">
          <el-input v-model="formData.titleEn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <template #>
            <div>
              <uploader-pic @updatePic="(res)=>formData.pic=res" @updatePicKey="(res)=>formData.picKey=res" :pic="formData.pic" :picKey="formData.picKey"></uploader-pic>
              <p style="child-align: bottom">图片建议16:9，尺寸：768x432px</p>
            </div>
          </template>
        </el-form-item>
        <el-form-item label="是否启用:" >
          <template #>
            <el-switch v-model="formData.status" :active-value="1" :inactive-value="2" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
          </template>
        </el-form-item>
        <el-form-item label="时间:">
          <el-date-picker v-model="formData.timeline" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="内容(中文):">
             <vue3-tinymce v-model="formData.articleCh" :setting="setting" />
        </el-form-item>
        <el-form-item label="内容(英文):">
           <vue3-tinymce v-model="formData.articleEn" :setting="setting" />
        </el-form-item>
        <el-form-item label="分类id:">
          <el-select v-model="formData.typeId" clearable placeholder="请选择">
            <el-option
                v-for="item in typesOptions"
                :key="item.id"
                :label="`${item.nameCh}`"
                :value="item.id"
            />
          </el-select>
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
  name: 'LukInformation',
  components: {
    'uploader-pic': uploaderPic
  },
}
</script>

<script setup>
import {
  createLukInformation,
  deleteLukInformation,
  updateLukInformation,
  findLukInformation,
  getLukInformationList
} from '@/api/lukInformation'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { getLukInformationTypeList } from '@/api/lukInformationType'
import { formatDate } from '@/utils/format'
import Vue3Tinymce from '@/components/Vue3Tinymce/Vue3Tinymce'
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        articleCh: '',
        articleEn: '',
        status: 1,
        timeline: new Date(),
        titleCh: '',
        titleEn: '',
        typeId: 0,
        pic: '',
        picKey: '',
        setting: {
       height: 190,
  toolbar:
    'undo redo | fullscreen | formatselect alignleft aligncenter alignright alignjustify | link unlink | numlist bullist | image media table | fontsizeselect forecolor backcolor | bold italic underline strikethrough | indent outdent | superscript subscript | removeformat |',
  toolbar_drawer: 'sliding',
  quickbars_selection_toolbar:
    'removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor',
  plugins: 'link image media table lists fullscreen quickbars',
  fontsize_formats: '12px 14px 16px 18px',
  default_link_target: '_blank',
  link_title: false,
  nonbreaking_force_tab: true,
  // 以中文简体为例
  language: 'zh_CN',
  language_url:
    'https://unpkg.com/@jsdawn/vue3-tinymce@1.1.6/dist/tinymce/langs/zh_CN.js'
        }
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const statusOptions = ref([])
const typesOptions = ref([])

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
  const table = await getLukInformationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  const table = await getLukInformationTypeList()
  if (table.code === 0) {
    typesOptions.value = table.data.list
    if (table.data.list.length > 0) {
      formData.value.typeId = table.data.list[0].id
    }
  }
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
            deleteLukInformationFunc(row)
        })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukInformationFunc = async(row) => {
    const res = await findLukInformation({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.relukDappInformation
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLukInformationFunc = async (row) => {
    const res = await deleteLukInformation({ id: row.id })
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
        articleCh: '',
        articleEn: '',
        status: 1,
        timeline: new Date(Date.parse(Date.now())),
        titleCh: '',
        titleEn: '',
        typeId: 0,
        pic: '',
        picKey: '',
        }
  if (typesOptions.value.length > 0) {
    formData.value.typeId = typesOptions.value[0].id
  }
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.titleCh == '' || formData.value.titleEn == '' || formData.value.articleCh == '' || formData.value.articleEn == '' || formData.value.timeline == ''){
    ElMessage.error('请补齐参数！')
    return
  }
  let res
  switch (type.value) {
    case 'create':
      res = await createLukInformation(formData.value)
      break
    case 'update':
      res = await updateLukInformation(formData.value)
      break
    default:
      res = await createLukInformation(formData.value)
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
  res = await updateLukInformation(values)
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
.online{
  height: 50px;
  width: 100px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
</style>
