import service from '@/utils/request'

// @Tags LukWithdrawApply
// @Summary 更新LukWithdrawApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukWithdrawApply true "更新LukWithdrawApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukWithdrawApply/updateLukWithdrawApply [put]
export const updateLukWithdrawApply = (data) => {
  return service({
    url: '/lukWithdrawApply/updateLukWithdrawApply',
    method: 'put',
    data
  })
}

// @Tags LukWithdrawApply
// @Summary 用id查询LukWithdrawApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukWithdrawApply true "用id查询LukWithdrawApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukWithdrawApply/findLukWithdrawApply [get]
export const findLukWithdrawApply = (params) => {
  return service({
    url: '/lukWithdrawApply/findLukWithdrawApply',
    method: 'get',
    params
  })
}

// @Tags LukWithdrawApply
// @Summary 分页获取LukWithdrawApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukWithdrawApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawApply/getLukWithdrawApplyList [get]
export const getLukWithdrawApplyList = (params) => {
  return service({
    url: '/lukWithdrawApply/getLukWithdrawApplyList',
    method: 'get',
    params
  })
}

export const getWithdrawalBalance = () => {
  return service({
    url: '/lukWithdrawApply/getWithdrawalBalance',
    method: 'get',
  })
}
