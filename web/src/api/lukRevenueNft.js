import service from '@/utils/request'

// @Tags LukRevenueNft
// @Summary 分页获取LukRevenueNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukRevenueNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenue/getLukRevenueNftList [get]
export const getLukRevenueNftList = (params) => {
  return service({
    url: '/lukRevenueNft/getLukRevenueNftList',
    method: 'get',
    params
  })
}

export const getLukRevenueHeritageList = (params) => {
  return service({
    url: '/lukRevenueHeritage/getLukRevenueHeritageList',
    method: 'get',
    params
  })
}
