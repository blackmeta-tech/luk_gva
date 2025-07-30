import service from '@/utils/request'

// @Tags LukRevenueDetailed
// @Summary 分页获取LukRevenueDetailed列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukRevenueDetailed列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailed/getLukRevenueDetailedList [get]
export const getLukRevenueDetailedList = (params) => {
  return service({
    url: '/lukRevenueDetailed/getLukRevenueDetailedList',
    method: 'get',
    params
  })
}
