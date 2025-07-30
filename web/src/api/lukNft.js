import service from '@/utils/request'

// @Tags LukNft
// @Summary 创建LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukNft true "创建LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNft/createLukNft [post]
export const createLukNft = (data) => {
  return service({
    url: '/lukNft/createLukNft',
    method: 'post',
    data
  })
}

// @Tags LukNft
// @Summary 更新LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukNft true "更新LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukNft/updateLukNft [put]
export const updateLukNft = (data) => {
  return service({
    url: '/lukNft/updateLukNft',
    method: 'put',
    data
  })
}

export const updateBlack = (data) => {
  return service({
    url: '/lukNft/updateLukNftBlack',
    method: 'put',
    data
  })
}

// @Tags LukNft
// @Summary 用id查询LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukNft true "用id查询LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukNft/findLukNft [get]
export const findLukNft = (params) => {
  return service({
    url: '/lukNft/findLukNft',
    method: 'get',
    params
  })
}

// @Tags LukNft
// @Summary 分页获取LukNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNft/getLukNftList [get]
export const getLukNftList = (params) => {
  return service({
    url: '/lukNft/getLukNftList',
    method: 'get',
    params
  })
}
