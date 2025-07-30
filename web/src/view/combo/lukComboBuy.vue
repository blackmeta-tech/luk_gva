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
        <el-form-item label="套餐ID">
          <el-input v-model="searchInfo.comboId" clearable placeholder="请输入" />
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
        >
        <el-table-column align="left" label="记录ID" prop="id" width="120" />
        <el-table-column align="left" label="套餐ID" prop="comboId" width="120" />
        <el-table-column align="left" label="套餐天数" prop="days" width="120" >
          <template #default="scope">
            {{ scope.row.days }}天套餐
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户地址" prop="address" width="380" />
        <el-table-column align="left" label="购入价格[LUK]" prop="price" width="150" />
        <el-table-column align="left" label="购入价格[USDT]" prop="priceUsdt" width="150" />
        <el-table-column align="left" label="收益率百分比" prop="rate" width="120" />
        <el-table-column align="left" label="复购状态">
          <template #default="scope">
            <el-tag v-show="scope.row.repeat == 1" type="success">复购</el-tag>
            <el-tag v-show="scope.row.repeat == 0">不复购</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="150" >
          <template #default="scope">
            <el-tag v-show="scope.row.status == 0" type="success">正常</el-tag>
            <el-tag v-show="scope.row.status == 1" type="success">到期已复购</el-tag>
            <el-tag v-show="scope.row.status == -1">已到期</el-tag>
            <el-tag v-show="scope.row.status == -2">复购失败,{{ scope.row.statusMsg }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="购买时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.buyAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="到期时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.maturityAt) }}</template>
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
  name: 'LukComboBuy'
}
</script>

<script setup>
import {
  getLukComboBuyList
} from '@/api/lukComboBuy'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'

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
  const table = await getLukComboBuyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

</script>

<style>
</style>
