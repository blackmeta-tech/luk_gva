import service from '@/utils/request'

// @Tags LukInformation
// @Summary 创建LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformation true "创建LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukInformation/createLukInformation [post]
export const createLukInformation = (data) => {
  return service({
    url: '/lukInformation/createLukInformation',
    method: 'post',
    data
  })
}

// @Tags LukInformation
// @Summary 删除LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformation true "删除LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukInformation/deleteLukInformation [delete]
export const deleteLukInformation = (data) => {
  return service({
    url: '/lukInformation/deleteLukInformation',
    method: 'delete',
    data
  })
}

// @Tags LukInformation
// @Summary 删除LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukInformation/deleteLukInformation [delete]
export const deleteLukInformationByIds = (data) => {
  return service({
    url: '/lukInformation/deleteLukInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags LukInformation
// @Summary 更新LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformation true "更新LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukInformation/updateLukInformation [put]
export const updateLukInformation = (data) => {
  return service({
    url: '/lukInformation/updateLukInformation',
    method: 'put',
    data
  })
}

// @Tags LukInformation
// @Summary 用id查询LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukInformation true "用id查询LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukInformation/findLukInformation [get]
export const findLukInformation = (params) => {
  return service({
    url: '/lukInformation/findLukInformation',
    method: 'get',
    params
  })
}

// @Tags LukInformation
// @Summary 分页获取LukInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukInformation/getLukInformationList [get]
export const getLukInformationList = (params) => {
  return service({
    url: '/lukInformation/getLukInformationList',
    method: 'get',
    params
  })
}
