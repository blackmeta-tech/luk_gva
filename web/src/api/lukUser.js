import service from '@/utils/request'

export const updateLinkage = (data) => {
  return service({
    url: '/lukUser/updateLinkage',
    method: 'post',
    data
  })
}

export const updateLukUser = (data) => {
  return service({
    url: '/lukUser/updateLukUser',
    method: 'post',
    data
  })
}

// @Tags LukUser
// @Summary 用id查询LukUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukUser true "用id查询LukUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukUser/findLukUser [get]
export const findLukUser = (params) => {
  return service({
    url: '/lukUser/findLukUser',
    method: 'get',
    params
  })
}

// @Tags LukUser
// @Summary 分页获取LukUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukUser/getLukUserList [get]
export const getLukUserList = (params) => {
  return service({
    url: '/lukUser/getLukUserList',
    method: 'get',
    params
  })
}
