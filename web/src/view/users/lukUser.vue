<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="注册时间">
        <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
         —
        <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
        <el-form-item label="用户地址">
          <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id">
          <el-input v-model="searchInfo.id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户地址">
          <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社区等级">
          <el-select v-model="searchInfo.communitylevel" clearable placeholder="请选择" @clear="clearSelect('communitylevel')">
            <el-option
                v-for="item in partnerTypeData"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="联盟白名单">
          <el-select v-model="searchInfo.islinkage" clearable placeholder="请选择" @clear="clearSelect('islinkage')">
            <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="联盟等级">
          <el-select v-model="searchInfo.linkagelevel" clearable placeholder="请选择" @clear="clearSelect('linkagelevel')">
            <el-option
                v-for="item in partnerTypeData"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
          <el-button size="small" type="success" icon="plus" @click="openDialog">联盟白名单</el-button>
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
        <el-table-column align="left" label="用户ID" prop="id" />
        <el-table-column align="left" label="用户地址" prop="address" width="380" />
        <el-table-column align="left" label="上级ID" prop="pid" />
        <el-table-column align="left" label="社区等级" prop="communityLevel" >
          <template #default="scope">{{ partnerTypeData.find(item => item.value === scope.row.communityLevel).label }}</template>
        </el-table-column>
        <el-table-column align="left" label="联盟白名单" prop="isLinkage" width="130" >
          <template #default="scope">
            <el-tag v-show="scope.row.isLinkage" type="success">是</el-tag>
            <el-tag v-show="scope.row.isLinkage == false">否</el-tag>
            <el-popover v-model="scope.row.visible" placement="top">
              <p>确定取消白名单吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="updateLukUserFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button v-show="scope.row.isLinkage" size="small" type="primary" link icon="delete" @click="scope.row.visible = true">解除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column align="left" label="联盟等级" prop="linkageLevel" width="100">
          <template #default="scope">
            <span v-show="scope.row.isLinkage">{{ partnerTypeData.find(item => item.value === scope.row.linkageLevel).label }}</span>
            <span v-show="scope.row.isLinkage == false">--</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="USDT数量" prop="usdt" width="120" />
        <el-table-column align="left" label="LUK数量" prop="luk" width="120" />
        <el-table-column align="left" label="小区交易量" prop="areaMin" width="120" />
        <el-table-column align="left" label="大区交易量" prop="areaMax" width="120" />
        <el-table-column align="left" label="层数" prop="level" />
        <el-table-column align="left" label="推荐码" prop="invit" width="100" />
        <el-table-column align="left" label="注册时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="最后登录时间" prop="lastLoginTime" width="180">
          <template #default="scope">{{ formatDate(scope.row.lastLoginTime) }}</template>
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
        <el-form-item label="用户地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
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
  name: 'LukUser'
}
</script>

<script setup>
import {
  updateLukUser,
  updateLinkage,
  getLukUserList
} from '@/api/lukUser'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import { getLukConstant } from '@/api/lukRevenue'
import { ElMessage } from 'element-plus'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const statusOptions = ref([])
const partnerTypeData = ref([])
const formData = ref({
  address: '',
})

// 查询
const getLukConstantList = async() => {
  const ks = ref(['partner_type'])
  const res = await getLukConstant({ keys: ks.value })
  if (res.code === 0) {
    partnerTypeData.value = res.data.partner_type
  }
}

getLukConstantList()
// 清空选项
const clearSelect = (name) => {
  if (name === 'communitylevel') {
    delete searchInfo.value.communitylevel
  } else if (name === 'islinkage') {
    delete searchInfo.value.islinkage
  } else if (name === 'linkagelevel') {
    delete searchInfo.value.linkagelevel
  }
}
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


// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  statusOptions.value = [
    {
      value: '1',
      label: '是',
      type: 'success'
    },
    {
      value: '0',
      label: '否',
      type: ''
    },
  ]
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 查询
const getTableData = async() => {
  const table = await getLukUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)
const openDialog = () => {
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    address: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.address === '') {
    ElMessage.error('地址必填')
    return
  }
  let res = await updateLinkage(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '添加成功'
    })
    closeDialog()
    getTableData()
  }
}

const updateLukUserFunc = async(row) => {
  formData.value = row
  formData.value.isLinkage = 0
  formData.value.linkageLevel = 0
  const res = await updateLukUser(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '解除成功'
    })
    closeDialog()
    getTableData()
    formData.value.address = null
  }
}
getTableData()

</script>

<style>
</style>
