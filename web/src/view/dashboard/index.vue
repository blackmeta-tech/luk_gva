<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>基础数据</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top"  :column="2" style="width: 98%" border>
                <el-descriptions-item label="LUK池U量"  width="25%">{{ panelData.pond.usdt }}</el-descriptions-item>
                <el-descriptions-item label="LUK池量">{{ panelData.pond.luk }}</el-descriptions-item>
                <el-descriptions-item label="LUK销毁量">{{ panelData.pond.destructionLuk }}</el-descriptions-item>
                <el-descriptions-item label="全网用户数量">{{ panelData.pond.address }}</el-descriptions-item>
                <el-descriptions-item label="已售NFT个数">{{ panelData.pond.nft }}</el-descriptions-item>
                <el-descriptions-item label="移除流动性手续费">{{ panelData.pond.removeLiquidityFee }}LUK</el-descriptions-item>
                <el-descriptions-item label="兑换买入手续费">{{ panelData.pond.swapFee }}LUK</el-descriptions-item>
                <el-descriptions-item label="兑换卖出手续费">{{ panelData.pond.sellFee }}LUK</el-descriptions-item>
                <el-descriptions-item label="兑换买入">{{ panelData.pond.swapUsdt }}USDT | {{ panelData.pond.swapLuk }}LUK</el-descriptions-item>
                <el-descriptions-item label="兑换卖出">{{ panelData.pond.sellUsdt }}USDT | {{ panelData.pond.sellLuk }}LUK</el-descriptions-item>
                <el-descriptions-item label="添加流动性总量">{{ panelData.pond.addLiquidity }}</el-descriptions-item>
                <el-descriptions-item label="移除流动性总量">{{ panelData.pond.removeLiquidity }}</el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>套餐相关</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top"  :column="2" style="width: 98%" border>
                <el-descriptions-item v-for="item in combos" :label="`${item.days}` + `天套餐剩余额度`">{{ item.remaining }}</el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>分红数据</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top"  :column="2" style="width: 98%" border>
                <el-descriptions-item label="NFT-推荐分红"  width="25%">{{ (+panelData.dividends.recommend).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="NFT-兑换买入">{{ (+panelData.dividends.nftSwap).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="NFT-兑换卖出">{{ (+panelData.dividends.nftSell).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="NFT-移除流动性">{{ (+panelData.dividends.nftRemoveLiquidity).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="NFT-加权套餐">{{ (+panelData.dividends.nftWeighted).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="NFT-1.5kU价值">{{ (+panelData.dividends.nftValue).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="套餐质押-7天">{{ (+panelData.dividends.combo7).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="套餐质押-15天">{{ (+panelData.dividends.combo15).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="套餐质押-30天">{{ (+panelData.dividends.combo30).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="套餐质押-60天">{{ (+panelData.dividends.combo60).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="套餐推荐"  span="2">{{ (+panelData.dividends.comboRecommend).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="联盟分红-青铜">{{ (+panelData.dividends.linkage1).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="联盟分红-白银">{{ (+panelData.dividends.linkage2).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="联盟分红-黄金">{{ (+panelData.dividends.linkage3).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="联盟分红-钻石">{{ (+panelData.dividends.linkage4).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="社区分红-青铜">{{ (+panelData.dividends.community1).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="社区分红-白银">{{ (+panelData.dividends.community2).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="社区分红-黄金">{{ (+panelData.dividends.community3).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="社区分红-钻石">{{ (+panelData.dividends.community4).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="LP-兑换买入">{{ (+panelData.dividends.lpSwap).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="LP-兑换卖出">{{ (+panelData.dividends.lpSell).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="LP-移除流动性">{{ (+panelData.dividends.lpRemoveLiquidity).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="LP-返佣">{{ (+panelData.dividends.lpRebate).toFixed(8) }}</el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>

    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>回流数据</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top" :column="2" style="width: 98%" border>
                <el-descriptions-item label="8%联盟分红[套餐购买]"  width="25%">{{ (+panelData.precipitation.alliance).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="3%STC底池[套餐购买]">{{ (+panelData.precipitation.stcPot).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="3%STC节点分红[套餐购买]">{{ (+panelData.precipitation.stcNode).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="3%运营钱包[套餐购买]">{{ (+panelData.precipitation.operate).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="1%NFT加权分红[套餐购买]">{{ (+panelData.precipitation.nft).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="1%联盟经费[套餐购买]">{{ (+panelData.precipitation.funding).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="1%联盟运营岗位津贴[套餐购买]"  span="2">{{ (+panelData.precipitation.allowance).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="1%回流底池[Lp]">{{ (+panelData.precipitation.lpPool).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="1%运营钱包[Lp]">{{ (+panelData.precipitation.lpOperate).toFixed(8) }}</el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>

    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>移产数据</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top" :column="2" style="width: 98%" border>
                <el-descriptions-item label="新元宇宙释放总额">{{ (+panelData.heritage.metaverseBalance).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="新元宇宙待释放总额">{{ (+panelData.heritage.metaverseRemaining).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="新元宇宙人数" width="25%"> {{ panelData.heritage.metaverseNum }} </el-descriptions-item>
                <el-descriptions-item label="新元宇宙分红">{{ (+panelData.heritage.metaverse).toFixed(8) }}LUK | {{ (+panelData.heritage.metaverseUsdt).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="老元宇宙释放总额">{{ (+panelData.heritage.metaverseOldBalance).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="老元宇宙待释放总额">{{ (+panelData.heritage.metaverseOldRemaining).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="老元宇宙人数" > {{ panelData.heritage.metaverseOldNum }} </el-descriptions-item>
                <el-descriptions-item label="老元宇宙分红">{{ (+panelData.heritage.metaverseOld).toFixed(8) }}LUK | {{ (+panelData.heritage.metaverseOldUsdt).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="LP矿机释放总额">{{ (+panelData.heritage.lpBalance).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="LP矿机待释放总额">{{ (+panelData.heritage.lpRemaining).toFixed(8) }}USDT</el-descriptions-item>
                <el-descriptions-item label="LP矿机人数" > {{ panelData.heritage.lpNum }} </el-descriptions-item>
                <el-descriptions-item label="LP矿机分红">{{ (+panelData.heritage.lp).toFixed(8) }}LUK | {{ (+panelData.heritage.lpUsdt).toFixed(8) }}USDT</el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>

    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>其他数据</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <template #default>
              <el-descriptions class="margin-top" :column="2" style="width: 98%" border>
                <el-descriptions-item label="USDT提现汇总"  width="25%">{{ (+panelData.other.withdrawUsdt).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item label="LUK提现汇总">{{ (+panelData.other.withdrawLuk).toFixed(8) }}</el-descriptions-item>
                <el-descriptions-item  label="节点" span="2"> {{ panelData.other.wss }} </el-descriptions-item>
                <el-descriptions-item  label="LUK合约" span="2"> {{ panelData.other.lukToken }} </el-descriptions-item>
                <el-descriptions-item  label="NFT合约" span="2"> {{ panelData.other.nftToken }} </el-descriptions-item>
                <el-descriptions-item  label="UDST合约" span="2"> {{ panelData.other.usdtToken }} </el-descriptions-item>
                <el-descriptions-item  label="LP合约" span="2"> {{ panelData.other.liquidityToken }} </el-descriptions-item>
              </el-descriptions>
            </template>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import service from '@/utils/request'
import { queryLukComboList } from '@/api/lukCombo'
const panelData = ref({
  pond: {},
  precipitation: {},
  dividends: {},
  other: {},
  heritage: {},
})
const combos = ref({})
// 获取面板数据
const getPanelData = async() => {
  var res = await getPanelDataDo()
  if (res.code === 0) {
    panelData.value = res.data
    console.log(panelData.value)
  } else {
    ElMessage({
      type: 'warning',
      message: '获取失败'
    })
  }
}

const getPanelDataDo = () => {
  return service({
    url: '/lukPanelData/queryData',
    method: 'get',
  })
}
getPanelData()

// 查询
const getComboList = async() => {
  const table = await queryLukComboList()
  if (table.code === 0) {
    console.log(table)
    combos.value = table.data
  }
}

getComboList()

</script>
<script>
export default {
  name: 'Dashboard'
}
</script>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}
.page {
  background: #f0f2f5;
  padding: 0;
.gva-card-box{
  padding: 12px 16px;
&+.gva-card-box{
   padding-top: 0px;
 }
}
.gva-card {
  box-sizing: border-box;
  background-color: #fff;
  border-radius: 2px;
  height: auto;
  padding: 26px 30px;
  overflow: hidden;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}
.gva-top-card {
  height: 260px;
  @include flex-center;
  justify-content: space-between;
  color: #777;
&-left {
   height: 100%;
   display: flex;
   flex-direction: column;
&-title {
   font-size: 22px;
   color: #343844;
 }
&-dot {
   font-size: 16px;
   color: #6B7687;
   margin-top: 24px;
 }
&-rows {
 // margin-top: 15px;
   margin-top: 18px;
   color: #6B7687;
   width: 600px;
   align-items: center;
 }
&-item{
+.gva-top-card-left-item{
  margin-top: 24px;
}
margin-top: 14px;
}
}
&-right {
   height: 600px;
   width: 600px;
   margin-top: 28px;
 }
}
::v-deep(.el-card__header){
  padding:0;
  border-bottom: none;
}
.card-header{
  padding-bottom: 20px;
  border-bottom: 1px solid #e8e8e8;
}
.quick-entrance-title {
  height: 30px;
  font-size: 22px;
  color: #333;
  width: 100%;
  border-bottom: 1px solid #eee;
}
.quick-entrance-items {
  @include flex-center;
  justify-content: center;
  text-align: center;
  color: #333;
.quick-entrance-item {
  padding: 16px 28px;
  margin-top: -16px;
  margin-bottom: -16px;
  border-radius: 4px;
  transition: all 0.2s;
&:hover{
   box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
 }
cursor: pointer;
height: auto;
text-align: center;
// align-items: center;
&-icon {
   width: 50px;
   height: 50px !important;
   border-radius: 8px;
   @include flex-center;
   justify-content: center;
   margin: 0 auto;
i {
  font-size: 24px;
}
}
p {
  margin-top: 10px;
}
}
}
.echart-box{
  padding: 14px;
}
}
.dashboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}
.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
  .gva-card {
    padding: 20px 10px !important;
  .gva-top-card {
    height: auto;
&-left {
&-title {
   font-size: 20px !important;
 }
&-rows {
   margin-top: 15px;
   align-items: center;
 }
}
&-right {
   display: none;
 }
}
.gva-middle-card {
&-item {
   line-height: 20px;
 }
}
.dashboard-icon {
  font-size: 18px;
}
}
}
</style>
