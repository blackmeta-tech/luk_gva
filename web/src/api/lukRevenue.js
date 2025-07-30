import service from '@/utils/request'

export const getLukConstant = (data) => {
  return service({
    url: '/lukConstant/getLukConstant',
    method: 'post',
    data
  })
}

// @Tags LukRevenue
// @Summary 分页获取LukRevenue列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukRevenue列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenue/getLukRevenueList [get]
export const getLukRevenueList = (params) => {
  return service({
    url: '/lukRevenue/getLukRevenueList',
    method: 'get',
    params
  })
}
