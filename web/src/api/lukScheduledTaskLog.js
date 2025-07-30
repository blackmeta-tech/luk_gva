import service from '@/utils/request'

// @Tags LukScheduledTaskLog
// @Summary 创建LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTaskLog true "创建LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukamiScheduledTaskLog/createLukScheduledTaskLog [post]
export const createLukScheduledTaskLog = (data) => {
  return service({
    url: '/lukScheduledTaskLog/createLukScheduledTaskLog',
    method: 'post',
    data
  })
}

// @Tags LukScheduledTaskLog
// @Summary 删除LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTaskLog true "删除LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukamiScheduledTaskLog/deleteLukScheduledTaskLog [delete]
export const deleteLukScheduledTaskLog = (data) => {
  return service({
    url: '/lukScheduledTaskLog/deleteLukScheduledTaskLog',
    method: 'delete',
    data
  })
}

// @Tags LukScheduledTaskLog
// @Summary 删除LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukamiScheduledTaskLog/deleteLukScheduledTaskLog [delete]
export const deleteLukScheduledTaskLogByIds = (data) => {
  return service({
    url: '/lukScheduledTaskLog/deleteLukScheduledTaskLogByIds',
    method: 'delete',
    data
  })
}

// @Tags LukScheduledTaskLog
// @Summary 更新LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukScheduledTaskLog true "更新LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukamiScheduledTaskLog/updateLukScheduledTaskLog [put]
export const updateLukScheduledTaskLog = (data) => {
  return service({
    url: '/lukScheduledTaskLog/updateLukScheduledTaskLog',
    method: 'put',
    data
  })
}

// @Tags LukScheduledTaskLog
// @Summary 用id查询LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukScheduledTaskLog true "用id查询LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukamiScheduledTaskLog/findLukScheduledTaskLog [get]
export const findLukScheduledTaskLog = (params) => {
  return service({
    url: '/lukScheduledTaskLog/findLukScheduledTaskLog',
    method: 'get',
    params
  })
}

// @Tags LukScheduledTaskLog
// @Summary 分页获取LukScheduledTaskLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukScheduledTaskLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukamiScheduledTaskLog/getLukScheduledTaskLogList [get]
export const getLukScheduledTaskLogList = (params) => {
  return service({
    url: '/lukScheduledTaskLog/getLukScheduledTaskLogList',
    method: 'get',
    params
  })
}
