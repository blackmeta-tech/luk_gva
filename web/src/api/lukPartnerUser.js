import service from '@/utils/request'

export const createLukPartnerUser = (data) => {
  return service({
    url: '/lukPartnerUser/createLukPartnerUser',
    method: 'post',
    data
  })
}

// @Tags LukPartnerUser
// @Summary 更新LukPartnerUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukPartnerUser true "更新LukPartnerUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukPartnerUser/updateLukPartnerUser [put]
export const updateLukPartnerUser = (data) => {
  return service({
    url: '/lukPartnerUser/updateLukPartnerUser',
    method: 'put',
    data
  })
}

// @Tags LukPartnerUser
// @Summary 用id查询LukPartnerUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukPartnerUser true "用id查询LukPartnerUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukPartnerUser/findLukPartnerUser [get]
export const findLukPartnerUser = (params) => {
  return service({
    url: '/lukPartnerUser/findLukPartnerUser',
    method: 'get',
    params
  })
}

// @Tags LukPartnerUser
// @Summary 分页获取LukPartnerUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukPartnerUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukPartnerUser/getLukPartnerUserList [get]
export const getLukPartnerUserList = (params) => {
  return service({
    url: '/lukPartnerUser/getLukPartnerUserList',
    method: 'get',
    params
  })
}
