<template>
  <div class="gva-table-box">
    <el-collapse v-model="activeNames" @change="handleChange">
      <el-collapse-item title="合约节点相关配置信息" name="1">
        <div>节点：{{ tokensData.wss }}</div>
        <div>NFT合约：{{ tokensData.nftToken }}</div>
        <div>LUK合约：{{ tokensData.lukToken }}</div>
        <div>USDT合约：{{ tokensData.usdtToken }}</div>
      </el-collapse-item>
      <el-collapse-item title="钱包信息" name="2">
        <el-table
            ref="multipleTable"
            style="width: 100%"
            tooltip-effect="dark"
            :data="tableData"
        >
          <el-table-column align="left" label="钱包配置名称" prop="abbr"/>
          <el-table-column align="left" label="钱包名称" prop="name"/>
          <el-table-column align="left" label="钱包地址" prop="address" width="400"/>
          <el-table-column align="left" label="钱包私钥" prop="privateKey">
            <template #default="scope">
              <el-tag v-show="scope.row.privateKey != ''" type="success">已配置</el-tag>
              <el-tag v-show="scope.row.privateKey == ''">无配置</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="Usdt余额" prop="usdt" />
          <el-table-column align="left" label="LUK余额" prop="luk" />
          <el-table-column align="left" label="BNB余额" prop="bnb" />
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
export default {
  name: 'Wallet'
}
</script>
<script setup>
import { ref } from 'vue'
import service from '@/utils/request'

const tableData = ref([])
const tokensData = ref([])

const initForm = async() => {
  const res = await getWallet()
  if (res.code === 0) {
    tableData.value = res.data.walletDatas
    tokensData.value = res.data.tokens
  }
}
const getWallet = () => {
  return service({
    url: '/lukWalletLog/getWallet',
    method: 'get'
  })
}
initForm()

</script>

<style lang="scss">
.system {
  background: #fff;
  padding:36px;
  border-radius: 2px;
h2 {
  padding: 10px;
  margin: 10px 0;
  font-size: 16px;
  box-shadow: -4px 0px 0px 0px #e7e8e8;
}
::v-deep(.el-input-number__increase){
  top:5px !important;
}
.gva-btn-list{
  margin-top:16px;
}
}
</style>
