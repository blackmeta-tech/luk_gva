import service from '@/utils/request'

// @Tags LukScheduledTask
// @Summary 创建LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTask true "创建LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukamiScheduledTask/createLukScheduledTask [post]
export const createLukScheduledTask = (data) => {
  return service({
    url: '/lukScheduledTask/createLukScheduledTask',
    method: 'post',
    data
  })
}

// @Tags LukScheduledTask
// @Summary 删除LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTask true "删除LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukamiScheduledTask/deleteLukScheduledTask [delete]
export const deleteLukScheduledTask = (data) => {
  return service({
    url: '/lukScheduledTask/deleteLukScheduledTask',
    method: 'delete',
    data
  })
}

// @Tags LukScheduledTask
// @Summary 删除LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukamiScheduledTask/deleteLukScheduledTask [delete]
export const deleteLukScheduledTaskByIds = (data) => {
  return service({
    url: '/lukScheduledTask/deleteLukScheduledTaskByIds',
    method: 'delete',
    data
  })
}

// @Tags LukScheduledTask
// @Summary 更新LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTask true "更新LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukamiScheduledTask/updateLukScheduledTask [put]
export const updateLukScheduledTask = (data) => {
  return service({
    url: '/lukScheduledTask/updateLukScheduledTask',
    method: 'put',
    data
  })
}

// @Tags LukScheduledTask
// @Summary 用id查询LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukScheduledTask true "用id查询LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukamiScheduledTask/findLukScheduledTask [get]
export const findLukScheduledTask = (params) => {
  return service({
    url: '/lukScheduledTask/findLukScheduledTask',
    method: 'get',
    params
  })
}

// @Tags LukScheduledTask
// @Summary 分页获取LukScheduledTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukScheduledTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukamiScheduledTask/getLukScheduledTaskList [get]
export const getLukScheduledTaskList = (params) => {
  return service({
    url: '/lukScheduledTask/getLukScheduledTaskList',
    method: 'get',
    params
  })
}

export const executeImmediately = (params) => {
  return service({
    url: '/lukScheduledTask/executeImmediately',
    method: 'get',
    params
  })
}
