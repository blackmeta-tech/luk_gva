<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
        <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
         —
        <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
        <el-form-item label="用户地址">
          <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="NFT名称">
          <el-input v-model="searchInfo.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上架状态">
          <el-select v-model="searchInfo.statu" clearable placeholder="请选择" @clear="clearSelect('statu')">
            <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="1.5kU黑名单">
          <el-select v-model="searchInfo.isBlacklist" clearable placeholder="请选择" @clear="clearSelect('isBlacklist')">
            <el-option
                v-for="item in blackOptions"
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
          <el-button size="small" icon="plus" @click="openBlackDialog">1.5kU黑名单</el-button>
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
        <el-table-column align="left" label="tokenID" prop="id"  />
        <el-table-column align="left" label="图片"  width="150">
          <template #default="scope">
            <el-image style="width: 100px; height: 100px" :src="scope.row.pic"></el-image>
          </template>
        </el-table-column>
        <el-table-column align="left" label="NFT名称" prop="name" width="100" />
        <el-table-column align="left" label="归属用户地址" prop="address" width="380" />
        <el-table-column align="left" label="上架价格[U]" prop="price" width="100" />
        <el-table-column align="left" label="1.5kU黑名单"  width="130">
          <template #default="scope">
            <span v-show="scope.row.blacklist == false">--</span>
            <span v-show="scope.row.blacklist">是</span>
            <span v-show="scope.row.blacklist">
              <el-popover  v-model:visible="scope.row.doExecuteImmediatelyVisible" placement="top" width="160">
                <p v-show="scope.row.prohibit == false" >确认禁止该NFT不参与全部分红吗</p>
                <p v-show="scope.row.prohibit" >确认允许该NFT参与其他分红吗</p>
                <p>大区交易量：{{ scope.row.areaMax }}</p>
                <p>小区区交易量：{{ scope.row.areaMin }}</p>
                <p>当前状态:
                  <span v-show="scope.row.prohibit == false">可参与其他分红</span>
                  <span v-show="scope.row.prohibit">禁止全部分红</span>
                </p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button size="small" type="text" @click="scope.row.doExecuteImmediatelyVisible = false">取消</el-button>
                  <el-button size="small" type="primary" @click="doProhibit(scope.row)">确认</el-button>
                </div>
                <template #reference>
                  <el-button type="text" size="small" @click="scope.row.doExecuteImmediatelyVisible = true">分红审核</el-button>
                </template>
              </el-popover>
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="上架状态" >
          <template #default="scope">
            <el-switch v-model="scope.row.status" :active-value="0" :inactive-value="-1" active-color="#13ce66" inactive-color="#ff4949" @change="switchDialog(scope.row,  scope.row.status)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createdAt)) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button v-show="scope.row.address == ''" type="primary" link icon="edit" size="small" class="table-button" @click="updateLukNftFunc(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item>
          <template #>
            <uploader-pic @updatePic="(res)=>formData.pic=res" @updatePicKey="(res)=>formData.picKey=res" :pic="formData.pic" :picKey="formData.picKey"></uploader-pic>
          </template>
        </el-form-item>
        <el-form-item label="NFT名称:"  prop="name" >
          <el-input v-model="formData.name"  :clearable="true"  />
        </el-form-item>
        <el-form-item label="上架价格[U]:"  prop="price" style="width:60%" >
          <template #>
            <el-input-number v-model="formData.price"  style="width:100%" :precision="2" />
          </template>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="dialogBlackVisible" :before-close="closeBlackDialog" title="加入黑名单">
      <el-form :model="formBlackData" label-position="right" label-width="150px">
        <el-form-item label="NFT编号ID:">
          <el-input-number v-model="formBlackData.id" clearable placeholder="请输入" style="width:60%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeBlackDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterBlackDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import uploaderPic from '@/view/lukNft/uploaderPic.vue'
export default {
  name: 'LukNftList',
  components: {
    'uploader-pic': uploaderPic
  },
}
</script>

<script setup>
import {
  createLukNft,
  updateLukNft,
  findLukNft,
  getLukNftList,
  updateBlack
} from '@/api/lukNft'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  pic: '',
  picKey: '',
  price: 0,
  name: '',
})
const formBlackData = ref({
  id: null,
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const statusOptions = ref([])
const blackOptions = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 清空选项
const clearSelect = (name) => {
  if (name === 'statu') {
    delete searchInfo.value.statu
  } else if (name === 'isBlacklist') {
    delete searchInfo.value.isBlacklist
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
  const table = await getLukNftList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
      value: '0',
      label: '上架',
      type: 'success'
    },
    {
      value: '-1',
      label: '下架',
      type: ''
    }
  ]
  blackOptions.value = [
    {
      value: '0',
      label: '否',
      type: 'success'
    },
    {
      value: '1',
      label: '是',
      type: ''
    }
  ]
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukNftFunc = async(row) => {
  const res = await findLukNft({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.relukNft
    dialogFormVisible.value = true
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = async() => {
  type.value = 'create'
  // 查询最后一条记录补齐表单
  if (tableData.value.length > 0) {
    const res = await findLukNft({ id: tableData.value[0].id })
    if (res.code === 0) {
      var nftData = res.data.relukNft
      formData.value = {
        pic: nftData.pic,
        picKey: nftData.picKey,
        price: nftData.price,
        name: nftData.name,
      }
    }
  }
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.price <= 0) {
    ElMessage.error('价格不得小于0!')
    return
  }
 elFormRef.value?.validate( async (valid) => {
   if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createLukNft(formData.value)
        break
      case 'update':
        res = await updateLukNft(formData.value)
        break
      default:
        res = await createLukNft(formData.value)
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
  })
}

// 修改禁止分红
const doProhibit = async(values) => {
  formData.value = values
  if (formData.value.prohibit) {
    formData.value.prohibit = false
  } else {
    formData.value.prohibit = true
  }
  let res
  res = await updateLukNft(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
    })
    getTableData()
  }
}

//开关控制
const switchDialog = async(values, num) => {
  let res
  values.status = num
  res = await updateLukNft(values)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
    })
    getTableData()
  }
}

// 弹窗控制标记
const dialogBlackVisible = ref(false)
const openBlackDialog = () => {
  dialogBlackVisible.value = true
}

// 关闭弹窗
const closeBlackDialog = () => {
  dialogBlackVisible.value = false
  formBlackData.value = {
    id: null,
  }
}
// 弹窗确定
const enterBlackDialog = async () => {
  if (formBlackData.value.id <= 0) {
    ElMessage.error('编号必填')
    return
  }
  let res = await updateBlack(formBlackData.value)
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
