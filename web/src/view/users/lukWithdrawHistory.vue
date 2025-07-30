<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="地址">
          <el-input v-model="searchInfo.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="币种类型">
          <el-select v-model="searchInfo.tokenType" clearable placeholder="请选择" @clear="clearSelect('tokenType')">
            <el-option
                v-for="item in tokenTypeData"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="clearSelect('status')">
            <el-option
                v-for="item in statusTypeData"
                :key="item.value"
                :label="`${item.label}`"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="提现时间">
          <div class="block">
            <el-date-picker
                v-model="searchInfoTimes"
                type="daterange"
                align="right"
                unlink-panels
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
                value-format='YYYY-MM-DD'
            >
            </el-date-picker>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
        <br />
        <el-form-item label="汇总">
          LUK提现总数：<el-tag>{{ lukAmount }}</el-tag>
          &nbsp; USDT提现总数：<el-tag>{{ usdtAmount }}</el-tag>
          &nbsp; 地址总数：<el-tag>{{ countAddress }}</el-tag>
          &nbsp; 提现条数：<el-tag>{{ total }}</el-tag>
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
            @sort-change="tableSortChange"
        >
          <el-table-column align="left" label="提现时间" prop="time" width="230">
            <template #default="scope">
              {{ formatDate(new Date(scope.row.time)) }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="用户地址" prop="address" width="400" />
          <el-table-column align="left" label="代币类型" prop="tokenType">
            <template #default="scope">{{ tokenTypeData.find(item => item.value === scope.row.tokenType).label }}</template>
          </el-table-column>
          <el-table-column align="left" label="申请提现金额" prop="amountPrimary" sortable="custom" width="140"  />
          <el-table-column align="left" label="手续费" prop="procedures" />
          <el-table-column align="left" label="实际提现金额" prop="amount" width="120"  />
          <el-table-column align="left" label="提现哈希" prop="signedTx" width="300" />
          <el-table-column align="left" label="交易状态">
            <template #default="scope">
              <el-tag v-show="scope.row.status == 1" type="success">成功</el-tag>
              <el-tag v-show="scope.row.status == 0">等待</el-tag>
              <el-tag v-show="scope.row.status == -1" type="danger">失败</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="按钮组">
            <template #default="scope">
              <el-button v-show="scope.row.status == 0" type="text" icon="edit" size="small" class="table-button" @click="updateLukWithdrawHistoryFunc(scope.row)">操作</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作" width="300px" center>
      <el-form :model="formData" label-position="right" label-width="30px">
        <el-form-item>
          <div style="margin-top: 20px">
            <el-radio v-model="formData.status" :label="1" border size="medium">成功</el-radio>
            <el-radio v-model="formData.status" :label="-1" border size="medium">失败</el-radio>
          </div>
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
  name: 'LukWithdrawHistory',
}
</script>

<script setup>
import {
  updateLukWithdrawHistory,
  findLukWithdrawHistory,
  getLukWithdrawHistoryList
} from '@/api/lukWithdrawHistory'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { toSQLLine } from '@/utils/stringFun'
import { getLukConstant } from '@/api/lukRevenue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  status: 0,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfoTimes = ref([])
const searchInfo = ref({})
const countAddress = ref(0)
const lukAmount = ref(0)
const usdtAmount = ref(0)

const tokenTypeData = ref([])
const statusTypeData = ref([])

// 查询
const getLukConstantList = async() => {
  const ks = ref(['token_type', 'status_type'])
  const res = await getLukConstant({ keys: ks.value })
  if (res.code === 0) {
    tokenTypeData.value = res.data.token_type
    statusTypeData.value = res.data.status_type
  }
}
getLukConstantList()

// 清空选项
const clearSelect = (name) => {
  if (name === 'type') {
    delete searchInfo.value.type
  } else if (name === 'tokenType') {
    delete searchInfo.value.tokenType
  } else if (name === 'status') {
    delete searchInfo.value.status
  }
}
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

// 排序
const tableSortChange = ({ prop, order }) => {
  searchInfo.value.orderKey = toSQLLine(prop)
  searchInfo.value.desc = order === 'descending'
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
  if (searchInfoTimes.value.length > 0) {
    searchInfo.value.startDate = searchInfoTimes.value[0]
    searchInfo.value.endDate = searchInfoTimes.value[1]
  }
  const table = await getLukWithdrawHistoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    countAddress.value = table.data.countAddress
    usdtAmount.value = table.data.usdtAmount
    lukAmount.value = table.data.lukAmount
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
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLukWithdrawHistoryFunc = async(row) => {
    const res = await findLukWithdrawHistory({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reWithdrawHistory
        dialogFormVisible.value = true
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        status: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'update':
          res = await updateLukWithdrawHistory(formData.value)
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
