import service from '@/utils/request'

// @Tags LukPartner
// @Summary 更新LukPartner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukPartner true "更新LukPartner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukPartner/updateLukPartner [put]
export const updateLukPartner = (data) => {
  return service({
    url: '/lukPartner/updateLukPartner',
    method: 'put',
    data
  })
}

// @Tags LukPartner
// @Summary 用id查询LukPartner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukPartner true "用id查询LukPartner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukPartner/findLukPartner [get]
export const findLukPartner = (params) => {
  return service({
    url: '/lukPartner/findLukPartner',
    method: 'get',
    params
  })
}

// @Tags LukPartner
// @Summary 分页获取LukPartner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukPartner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukPartner/getLukPartnerList [get]
export const getLukPartnerList = (params) => {
  return service({
    url: '/lukPartner/getLukPartnerList',
    method: 'get',
    params
  })
}
