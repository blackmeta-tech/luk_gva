<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px" >
      <el-tabs type="border-card">
        <el-tab-pane label="基础配置">
          <el-form-item label="Luk提现最低限额">
            <el-input-number v-model.number="config.withdrawalQuota" />
          </el-form-item>
          <el-form-item label="Luk每日限额">
            <el-input-number v-model.number="config.withdrawalLuk" />
          </el-form-item>
          <el-form-item label="USDT每日限额">
            <el-input-number v-model.number="config.withdrawalUsdt" />
          </el-form-item>
          <el-form-item label="提现手续费百分比">
            <el-input-number v-model.number="config.withdrawalCharge" />%
          </el-form-item>
          <el-form-item label="NFT发行量">
            <el-input-number v-model.number="config.nftIssued" />
          </el-form-item>
          <el-form-item label="代码发布状态">
            <el-radio v-model="config.online" :label="true" border size="medium">发布中</el-radio>
            <el-radio v-model="config.online" :label="false" border size="medium">未发布</el-radio>
          </el-form-item>
          <el-form-item label="代码发布文本">
            <el-input v-model.number="config.onlineText" />
          </el-form-item>
          <el-form-item label="LP最高返佣额度">
            <el-input-number v-model.number="config.lpRebateHighest" />U
          </el-form-item>
          <el-form-item label="LP返佣百分比">
            <el-input-number v-model.number="config.lpRebate" />%
          </el-form-item>
          <el-form-item label="返佣业绩分红百分比">
            <el-input-number v-model.number="config.rebatePerformance" />%
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="套餐配置">
          <el-form-item label="推荐分红第1代百分比">
            <el-input-number v-model.number="config.comboRecommend1" />%
          </el-form-item>
          <el-form-item label="推荐分红第2-5代百分比">
            <el-input-number v-model.number="config.comboRecommend2" />%
          </el-form-item>
          <el-form-item label="推荐分红第6-10代百分比">
            <el-input-number v-model.number="config.comboRecommend6" />%
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane label="移产配置">
          <el-form-item label="元宇宙模式一（固定）释放百分比">
            <el-input-number v-model.number="config.ratioMetaverse1" />%
          </el-form-item>
          <el-form-item label="元宇宙模式二（套餐）释放百分比">
            <el-input-number v-model.number="config.ratioMetaverse2" />%
          </el-form-item>
          <el-form-item label="元宇宙模式三（LP）释放百分比">
            <el-input-number v-model.number="config.ratioMetaverse3" />%
          </el-form-item>
          <el-form-item label="老元宇宙释放百分比">
            <el-input-number v-model.number="config.ratioMetaverseOld" />%
          </el-form-item>
          <el-form-item label="Lp矿机释放百分比">
            <el-input-number v-model.number="config.ratioLp" />%
          </el-form-item>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="small" @click="update">立即更新</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SystemLuk'
}
</script>
<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getConfig, setConfig } from '@/api/config'

const config = ref({})

const initForm = async() => {
  const res = await getConfig()
  if (res.code === 0) {
    config.value = res.data
  }
}
initForm()
const update = async() => {
  const res = await setConfig(config.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}
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
