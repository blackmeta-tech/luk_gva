<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户地址">
          <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分红类型">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择">
            <el-option
                v-for="item in revenueTypeNftData"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="币种类型">
          <el-select v-model="searchInfo.tokenType" clearable placeholder="请选择">
            <el-option
                v-for="item in tokenTypeData"
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
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="id"
      >
        <el-table-column align="left" label="用户地址" prop="address" width="380" />
        <el-table-column align="left" label="代币类型" prop="tokenType" >
          <template #default="scope">{{ tokenTypeData.find(item => item.value === scope.row.tokenType).label }}</template>
        </el-table-column>
        <el-table-column align="left" label="币种数量" prop="amount">
          <template #default="scope">
            {{ scope.row.amount }}<span v-show="scope.row.category == 2">[{{ scope.row.remark }}]</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="USDT计价" prop="relatively">
          <template #default="scope">
            <span v-show="scope.row.relatively > 0">{{ scope.row.relatively }}</span>
            <span v-show="scope.row.relatively == 0">--</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="分红类型" prop="type" >
          <template #default="scope">{{ revenueTypeNftData.find(item => item.value === scope.row.type).label }}</template>
        </el-table-column>
        <el-table-column align="left" label="来源" prop="source">
          <template #default="scope">
            <span v-show="scope.row.type == 22">NFT编号[{{ scope.row.source }}]</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="记录日期" width="180">
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
  </div>
</template>

<script>
export default {
  name: 'LukRevenueDetailedNft'
}
</script>

<script setup>
import { getLukConstant } from '@/api/lukRevenue'
import { getLukRevenueDetailedNftList } from '@/api/lukRevenueDetailedNft'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const revenueTypeNftData = ref([])
const tokenTypeData = ref([])
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
  const table = await getLukRevenueDetailedNftList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 查询
const getLukConstantList = async() => {
  const ks = ref(['token_type', 'revenue_type_nft'])
  const res = await getLukConstant({ keys: ks.value })
  if (res.code === 0) {
    revenueTypeNftData.value = res.data.revenue_type_nft
    tokenTypeData.value = res.data.token_type
  }
}

getLukConstantList()
getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()

</script>

<style>
</style>
