import service from '@/utils/request'

// @Tags LukComboBuy
// @Summary 创建LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukComboBuy true "创建LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukComboBuy/createLukComboBuy [post]
export const createLukComboBuy = (data) => {
  return service({
    url: '/lukComboBuy/createLukComboBuy',
    method: 'post',
    data
  })
}

// @Tags LukComboBuy
// @Summary 删除LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukComboBuy true "删除LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukComboBuy/deleteLukComboBuy [delete]
export const deleteLukComboBuy = (data) => {
  return service({
    url: '/lukComboBuy/deleteLukComboBuy',
    method: 'delete',
    data
  })
}

// @Tags LukComboBuy
// @Summary 删除LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukComboBuy/deleteLukComboBuy [delete]
export const deleteLukComboBuyByIds = (data) => {
  return service({
    url: '/lukComboBuy/deleteLukComboBuyByIds',
    method: 'delete',
    data
  })
}

// @Tags LukComboBuy
// @Summary 更新LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukComboBuy true "更新LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukComboBuy/updateLukComboBuy [put]
export const updateLukComboBuy = (data) => {
  return service({
    url: '/lukComboBuy/updateLukComboBuy',
    method: 'put',
    data
  })
}

// @Tags LukComboBuy
// @Summary 用id查询LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukComboBuy true "用id查询LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukComboBuy/findLukComboBuy [get]
export const findLukComboBuy = (params) => {
  return service({
    url: '/lukComboBuy/findLukComboBuy',
    method: 'get',
    params
  })
}

// @Tags LukComboBuy
// @Summary 分页获取LukComboBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukComboBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukComboBuy/getLukComboBuyList [get]
export const getLukComboBuyList = (params) => {
  return service({
    url: '/lukComboBuy/getLukComboBuyList',
    method: 'get',
    params
  })
}
