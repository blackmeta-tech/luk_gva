import service from '@/utils/request'

// @Tags LukBanner
// @Summary 创建LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukBanner true "创建LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukBanner/createLukBanner [post]
export const createLukBanner = (data) => {
  return service({
    url: '/lukBanner/createLukBanner',
    method: 'post',
    data
  })
}

// @Tags LukBanner
// @Summary 删除LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukBanner true "删除LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukBanner/deleteLukBanner [delete]
export const deleteLukBanner = (data) => {
  return service({
    url: '/lukBanner/deleteLukBanner',
    method: 'delete',
    data
  })
}

// @Tags LukBanner
// @Summary 删除LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukBanner/deleteLukBanner [delete]
export const deleteLukBannerByIds = (data) => {
  return service({
    url: '/lukBanner/deleteLukBannerByIds',
    method: 'delete',
    data
  })
}

// @Tags LukBanner
// @Summary 更新LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukBanner true "更新LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukBanner/updateLukBanner [put]
export const updateLukBanner = (data) => {
  return service({
    url: '/lukBanner/updateLukBanner',
    method: 'put',
    data
  })
}

// @Tags LukBanner
// @Summary 用id查询LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukBanner true "用id查询LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukBanner/findLukBanner [get]
export const findLukBanner = (params) => {
  return service({
    url: '/lukBanner/findLukBanner',
    method: 'get',
    params
  })
}

// @Tags LukBanner
// @Summary 分页获取LukBanner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukBanner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukBanner/getLukBannerList [get]
export const getLukBannerList = (params) => {
  return service({
    url: '/lukBanner/getLukBannerList',
    method: 'get',
    params
  })
}
