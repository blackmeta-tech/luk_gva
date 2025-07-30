<template>
  <div>
    <el-tabs v-model="searchInfo.type" type="border-card"  @tab-click="handleClick">
      <el-tab-pane v-for="item in userTypeData" :key=item.value :label="`${item.label}`" :name=item.value ></el-tab-pane>
      <div class="gva-table-box">
          <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="id"
          >
          <el-table-column align="left" label="类型" prop="type" >
            <template #default="scope">{{ userTypeData.find(item => item.value === scope.row.type).label }}</template>
          </el-table-column>
          <el-table-column align="left" label="级别" prop="level" >
            <template #default="scope">{{ partnerTypeData.find(item => item.value === scope.row.level).label }}</template>
          </el-table-column>
          <el-table-column align="left" label="小区交易量" prop="areaMin" />
          <el-table-column align="left" label="大区交易量" prop="areaMax" />
          <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateLukPartnerFunc(scope.row)">变更</el-button>
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
    </el-tabs>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="150px">
        <el-form-item label="小区交易量:">
          <el-input-number v-model.number="formData.areaMin" style="width:60%" :precision="2"  :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="大区交易量:">
          <el-input-number v-model.number="formData.areaMax" style="width:60%" :precision="2"  :clearable="true" placeholder="请输入" />

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
  name: 'LukPartner',
  data() {
    return {
      value: true
    }
  }
}
</script>

<script setup>
import {
  updateLukPartner,
  findLukPartner,
  getLukPartnerList
} from '@/api/lukPartner'

// 全量引入格式化工具 请按需保留
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { getLukConstant } from '@/api/lukRevenue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  areaMax: 0,
  areaMin: 0,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const userTypeData = ref([])
const partnerTypeData = ref([])
const searchInfo = ref({
  type: 1,
})

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

// 搜索
const handleClick = (tab) => {
  searchInfo.value.type = parseInt(tab.index) + 1
  page.value = 1
  pageSize.value = 10
  getTableData()
}


// 查询
const getLukConstantList = async() => {
  const ks = ref(['user_type', 'partner_type'])
  const res = await getLukConstant({ keys: ks.value })
  if (res.code === 0) {
    userTypeData.value = res.data.user_type
    partnerTypeData.value = res.data.partner_type
  }
}

getLukConstantList()

// 查询
const getTableData = async() => {
  const table = await getLukPartnerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukPartnerFunc = async(row) => {
  const res = await findLukPartner({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reLukPartner
    dialogFormVisible.value = true
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    areaMax: 0,
    areaMin: 0,
  }
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.areaMin < 0 || formData.value.areaMax < 0) {
    ElMessage.error('请检查所填的数值')
    return
  }
    let res
    switch (type.value) {
      case 'update':
        res = await updateLukPartner(formData.value)
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
