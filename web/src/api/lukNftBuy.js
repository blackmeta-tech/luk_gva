import service from '@/utils/request'

// @Tags LukNftBuy
// @Summary 分页获取LukNftBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukNftBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftBuy/getLukNftBuyList [get]
export const getLukNftBuyList = (params) => {
  return service({
    url: '/lukNftBuy/getLukNftBuyList',
    method: 'get',
    params
  })
}
