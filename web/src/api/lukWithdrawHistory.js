import service from '@/utils/request'

// @Tags LukWithdrawHistory
// @Summary 更新LukWithdrawHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukWithdrawHistory true "更新LukWithdrawHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /WithdrawHistory/updateLukWithdrawHistory [put]
export const updateLukWithdrawHistory = (data) => {
  return service({
    url: '/lukWithdrawHistory/updateLukWithdrawHistory',
    method: 'put',
    data
  })
}

// @Tags LukWithdrawHistory
// @Summary 用id查询LukWithdrawHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukWithdrawHistory true "用id查询LukWithdrawHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /WithdrawHistory/findLukWithdrawHistory [get]
export const findLukWithdrawHistory = (params) => {
  return service({
    url: '/lukWithdrawHistory/findLukWithdrawHistory',
    method: 'get',
    params
  })
}

// @Tags LukWithdrawHistory
// @Summary 分页获取LukWithdrawHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukWithdrawHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /WithdrawHistory/getLukWithdrawHistoryList [get]
export const getLukWithdrawHistoryList = (params) => {
  return service({
    url: '/lukWithdrawHistory/getLukWithdrawHistoryList',
    method: 'get',
    params
  })
}
