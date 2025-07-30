import service from '@/utils/request'

// @Tags LukRebate
// @Summary 创建LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukRebate true "创建LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heritage/createLukRebate [post]
export const createLukRebate = (data) => {
  return service({
    url: '/lukRebate/createLukRebate',
    method: 'post',
    data
  })
}

// @Tags LukRebate
// @Summary 删除LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukRebate true "删除LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heritage/deleteLukRebate [delete]
export const deleteLukRebate = (data) => {
  return service({
    url: '/lukRebate/deleteLukRebate',
    method: 'delete',
    data
  })
}

// @Tags LukRebate
// @Summary 删除LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heritage/deleteLukRebate [delete]
export const deleteLukRebateByIds = (data) => {
  return service({
    url: '/lukRebate/deleteLukRebateByIds',
    method: 'delete',
    data
  })
}

// @Tags LukRebate
// @Summary 更新LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukRebate true "更新LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /heritage/updateLukRebate [put]
export const updateLukRebate = (data) => {
  return service({
    url: '/lukRebate/updateLukRebate',
    method: 'put',
    data
  })
}

// @Tags LukRebate
// @Summary 用id查询LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukRebate true "用id查询LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /heritage/findLukRebate [get]
export const findLukRebate = (params) => {
  return service({
    url: '/lukRebate/findLukRebate',
    method: 'get',
    params
  })
}

// @Tags LukRebate
// @Summary 分页获取LukRebate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukRebate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heritage/getLukRebateList [get]
export const getLukRebateList = (params) => {
  return service({
    url: '/lukRebate/getLukRebateList',
    method: 'get',
    params
  })
}
