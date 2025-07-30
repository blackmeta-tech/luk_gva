import service from '@/utils/request'

// @Tags LukHeritageMetaverse
// @Summary 创建LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverse true "创建LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heritage/createLukHeritageMetaverse [post]
export const createLukHeritageMetaverse = (data) => {
  return service({
    url: '/lukHeritageMetaverse/createLukHeritageMetaverse',
    method: 'post',
    data
  })
}

// @Tags LukHeritageMetaverse
// @Summary 删除LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverse true "删除LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heritage/deleteLukHeritageMetaverse [delete]
export const deleteLukHeritageMetaverse = (data) => {
  return service({
    url: '/lukHeritageMetaverse/deleteLukHeritageMetaverse',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageMetaverse
// @Summary 删除LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heritage/deleteLukHeritageMetaverse [delete]
export const deleteLukHeritageMetaverseByIds = (data) => {
  return service({
    url: '/lukHeritageMetaverse/deleteLukHeritageMetaverseByIds',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageMetaverse
// @Summary 更新LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverse true "更新LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /heritage/updateLukHeritageMetaverse [put]
export const updateLukHeritageMetaverse = (data) => {
  return service({
    url: '/lukHeritageMetaverse/updateLukHeritageMetaverse',
    method: 'put',
    data
  })
}

// @Tags LukHeritageMetaverse
// @Summary 用id查询LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukHeritageMetaverse true "用id查询LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /heritage/findLukHeritageMetaverse [get]
export const findLukHeritageMetaverse = (params) => {
  return service({
    url: '/lukHeritageMetaverse/findLukHeritageMetaverse',
    method: 'get',
    params
  })
}

// @Tags LukHeritageMetaverse
// @Summary 分页获取LukHeritageMetaverse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukHeritageMetaverse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heritage/getLukHeritageMetaverseList [get]
export const getLukHeritageMetaverseList = (params) => {
  return service({
    url: '/lukHeritageMetaverse/getLukHeritageMetaverseList',
    method: 'get',
    params
  })
}
