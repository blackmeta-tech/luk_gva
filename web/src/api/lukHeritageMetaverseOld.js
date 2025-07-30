import service from '@/utils/request'

// @Tags LukHeritageMetaverseOld
// @Summary 创建LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverseOld true "创建LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverseOld/createLukHeritageMetaverseOld [post]
export const createLukHeritageMetaverseOld = (data) => {
  return service({
    url: '/lukHeritageMetaverseOld/createLukHeritageMetaverseOld',
    method: 'post',
    data
  })
}

// @Tags LukHeritageMetaverseOld
// @Summary 删除LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverseOld true "删除LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageMetaverseOld/deleteLukHeritageMetaverseOld [delete]
export const deleteLukHeritageMetaverseOld = (data) => {
  return service({
    url: '/lukHeritageMetaverseOld/deleteLukHeritageMetaverseOld',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageMetaverseOld
// @Summary 删除LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageMetaverseOld/deleteLukHeritageMetaverseOld [delete]
export const deleteLukHeritageMetaverseOldByIds = (data) => {
  return service({
    url: '/lukHeritageMetaverseOld/deleteLukHeritageMetaverseOldByIds',
    method: 'delete',
    data
  })
}

// @Tags LukHeritageMetaverseOld
// @Summary 更新LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LukHeritageMetaverseOld true "更新LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukHeritageMetaverseOld/updateLukHeritageMetaverseOld [put]
export const updateLukHeritageMetaverseOld = (data) => {
  return service({
    url: '/lukHeritageMetaverseOld/updateLukHeritageMetaverseOld',
    method: 'put',
    data
  })
}

// @Tags LukHeritageMetaverseOld
// @Summary 用id查询LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LukHeritageMetaverseOld true "用id查询LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukHeritageMetaverseOld/findLukHeritageMetaverseOld [get]
export const findLukHeritageMetaverseOld = (params) => {
  return service({
    url: '/lukHeritageMetaverseOld/findLukHeritageMetaverseOld',
    method: 'get',
    params
  })
}

// @Tags LukHeritageMetaverseOld
// @Summary 分页获取LukHeritageMetaverseOld列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LukHeritageMetaverseOld列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverseOld/getLukHeritageMetaverseOldList [get]
export const getLukHeritageMetaverseOldList = (params) => {
  return service({
    url: '/lukHeritageMetaverseOld/getLukHeritageMetaverseOldList',
    method: 'get',
    params
  })
}
