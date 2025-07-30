import service from '@/utils/request'

// @Tags LukRevenueDetailedNft
// @Summary 分页获取LukRevenueDetailedNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukRevenueDetailedNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailed/getLukRevenueDetailedNftList [get]
export const getLukRevenueDetailedNftList = (params) => {
  return service({
    url: '/lukRevenueDetailedNft/getLukRevenueDetailedNftList',
    method: 'get',
    params
  })
}

export const getLukRevenueDetailedHeritageList = (params) => {
  return service({
    url: '/lukRevenueDetailedHeritage/getLukRevenueDetailedHeritageList',
    method: 'get',
    params
  })
}
