import service from '@/utils/request'

// @Tags LukInformationType
// @Summary 创建LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformationType true "创建LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukInformationType/createLukInformationType [post]
export const createLukInformationType = (data) => {
  return service({
    url: '/lukInformationType/createLukInformationType',
    method: 'post',
    data
  })
}

// @Tags LukInformationType
// @Summary 删除LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformationType true "删除LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukInformationType/deleteLukInformationType [delete]
export const deleteLukInformationType = (data) => {
  return service({
    url: '/lukInformationType/deleteLukInformationType',
    method: 'delete',
    data
  })
}

// @Tags LukInformationType
// @Summary 删除LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukInformationType/deleteLukInformationType [delete]
export const deleteLukInformationTypeByIds = (data) => {
  return service({
    url: '/lukInformationType/deleteLukInformationTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags LukInformationType
// @Summary 更新LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukInformationType true "更新LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukInformationType/updateLukInformationType [put]
export const updateLukInformationType = (data) => {
  return service({
    url: '/lukInformationType/updateLukInformationType',
    method: 'put',
    data
  })
}

// @Tags LukInformationType
// @Summary 用id查询LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukInformationType true "用id查询LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukInformationType/findLukInformationType [get]
export const findLukInformationType = (params) => {
  return service({
    url: '/lukInformationType/findLukInformationType',
    method: 'get',
    params
  })
}

// @Tags LukInformationType
// @Summary 分页获取LukInformationType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukInformationType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukInformationType/getLukInformationTypeList [get]
export const getLukInformationTypeList = (params) => {
  return service({
    url: '/lukInformationType/getLukInformationTypeList',
    method: 'get',
    params
  })
}
