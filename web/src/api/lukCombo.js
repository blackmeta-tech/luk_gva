import service from '@/utils/request'

// @Tags LukCombo
// @Summary 创建LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukCombo true "创建LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukCombo/createLukCombo [post]
export const createLukCombo = (data) => {
  return service({
    url: '/lukCombo/createLukCombo',
    method: 'post',
    data
  })
}

// @Tags LukCombo
// @Summary 删除LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukCombo true "删除LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukCombo/deleteLukCombo [delete]
export const deleteLukCombo = (data) => {
  return service({
    url: '/lukCombo/deleteLukCombo',
    method: 'delete',
    data
  })
}

// @Tags LukCombo
// @Summary 删除LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukCombo/deleteLukCombo [delete]
export const deleteLukComboByIds = (data) => {
  return service({
    url: '/lukCombo/deleteLukComboByIds',
    method: 'delete',
    data
  })
}

// @Tags LukCombo
// @Summary 更新LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukCombo true "更新LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukCombo/updateLukCombo [put]
export const updateLukCombo = (data) => {
  return service({
    url: '/lukCombo/updateLukCombo',
    method: 'put',
    data
  })
}

// @Tags LukCombo
// @Summary 用id查询LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukCombo true "用id查询LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukCombo/findLukCombo [get]
export const findLukCombo = (params) => {
  return service({
    url: '/lukCombo/findLukCombo',
    method: 'get',
    params
  })
}

// @Tags LukCombo
// @Summary 分页获取LukCombo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukCombo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukCombo/getLukComboList [get]
export const getLukComboList = (params) => {
  return service({
    url: '/lukCombo/getLukComboList',
    method: 'get',
    params
  })
}

export const queryLukComboList = (params) => {
  return service({
    url: '/lukCombo/list',
    method: 'get',
    params
  })
}
