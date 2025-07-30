import service from '@/utils/request'

// @Tags LukHeritageLp
// @Summary 创建LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageLp true "创建LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageLp/createLukHeritageLp [post]
export const createLukHeritageLp = (data) => {
  return service({
    url: '/lukHeritageLp/createLukHeritageLp',
    method: 'post',
    data
  })
}

// @Tags LukHeritageLp
// @Summary 删除LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageLp true "删除LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageLp/deleteLukHeritageLp [delete]
export const deleteLukHeritageLp = (data) => {
  return service({
    url: '/lukHeritageLp/deleteLukHeritageLp',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageLp
// @Summary 删除LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageLp/deleteLukHeritageLp [delete]
export const deleteLukHeritageLpByIds = (data) => {
  return service({
    url: '/lukHeritageLp/deleteLukHeritageLpByIds',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageLp
// @Summary 更新LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageLp true "更新LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukHeritageLp/updateLukHeritageLp [put]
export const updateLukHeritageLp = (data) => {
  return service({
    url: '/lukHeritageLp/updateLukHeritageLp',
    method: 'put',
    data
  })
}

// @Tags LukHeritageLp
// @Summary 用id查询LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukHeritageLp true "用id查询LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukHeritageLp/findLukHeritageLp [get]
export const findLukHeritageLp = (params) => {
  return service({
    url: '/lukHeritageLp/findLukHeritageLp',
    method: 'get',
    params
  })
}

// @Tags LukHeritageLp
// @Summary 分页获取LukHeritageLp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukHeritageLp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageLp/getLukHeritageLpList [get]
export const getLukHeritageLpList = (params) => {
  return service({
    url: '/lukHeritageLp/getLukHeritageLpList',
    method: 'get',
    params
  })
}
