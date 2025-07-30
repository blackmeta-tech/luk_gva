import service from '@/utils/request'

// @Tags LukNftTransfer
// @Summary 分页获取LukNftTransfer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukNftTransfer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftTransfer/getLukNftTransferList [get]
export const getLukNftTransferList = (params) => {
  return service({
    url: '/lukNftTransfer/getLukNftTransferList',
    method: 'get',
    params
  })
}
