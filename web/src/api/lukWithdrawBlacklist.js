import service from '@/utils/request'

// @Tags LukWithdrawBlacklist
// @Summary 创建LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukWithdrawBlacklist true "创建LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawBlacklist/createLukWithdrawBlacklist [post]
export const createLukWithdrawBlacklist = (data) => {
  return service({
    url: '/lukWithdrawBlacklist/createLukWithdrawBlacklist',
    method: 'post',
    data
  })
}

// @Tags LukWithdrawBlacklist
// @Summary 删除LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukWithdrawBlacklist true "删除LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukWithdrawBlacklist/deleteLukWithdrawBlacklist [delete]
export const deleteLukWithdrawBlacklist = (data) => {
  return service({
    url: '/lukWithdrawBlacklist/deleteLukWithdrawBlacklist',
    method: 'delete',
    data
  })
}

// @Tags LukWithdrawBlacklist
// @Summary 删除LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukWithdrawBlacklist/deleteLukWithdrawBlacklist [delete]
export const deleteLukWithdrawBlacklistByIds = (data) => {
  return service({
    url: '/lukWithdrawBlacklist/deleteLukWithdrawBlacklistByIds',
    method: 'delete',
    data
  })
}

// @Tags LukWithdrawBlacklist
// @Summary 更新LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukWithdrawBlacklist true "更新LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukWithdrawBlacklist/updateLukWithdrawBlacklist [put]
export const updateLukWithdrawBlacklist = (data) => {
  return service({
    url: '/lukWithdrawBlacklist/updateLukWithdrawBlacklist',
    method: 'put',
    data
  })
}

// @Tags LukWithdrawBlacklist
// @Summary 用id查询LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukWithdrawBlacklist true "用id查询LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukWithdrawBlacklist/findLukWithdrawBlacklist [get]
export const findLukWithdrawBlacklist = (params) => {
  return service({
    url: '/lukWithdrawBlacklist/findLukWithdrawBlacklist',
    method: 'get',
    params
  })
}

// @Tags LukWithdrawBlacklist
// @Summary 分页获取LukWithdrawBlacklist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukWithdrawBlacklist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawBlacklist/getLukWithdrawBlacklistList [get]
export const getLukWithdrawBlacklistList = (params) => {
  return service({
    url: '/lukWithdrawBlacklist/getLukWithdrawBlacklistList',
    method: 'get',
    params
  })
}
