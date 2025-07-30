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
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.updatedAt)) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户地址" prop="address" width="380" />
        <el-table-column align="left" label="销毁个数" prop="luk" />
        <el-table-column align="left" label="销毁哈希" prop="txHash"  width="300" />
        <el-table-column align="left" label="销毁状态">
          <template #default="scope">
            <el-tag v-show="scope.row.status == 1" type="success">成功</el-tag>
            <el-tag v-show="scope.row.status == 0">等待</el-tag>
            <el-tag v-show="scope.row.status == -1" type="danger">失败</el-tag>
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
  name: 'LukDestroy'
}
</script>

<script setup>

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import service from '@/utils/request'

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
  const table = await getDxbaDbaDestroyInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const getDxbaDbaDestroyInfoList = (params) => {
  return service({
    url: '/lukCombo/getLukDestroyInfoList',
    method: 'get',
    params
  })
}

getTableData()


</script>

<style>
</style>
