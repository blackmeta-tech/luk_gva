package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LukHeritageMetaverseApi struct {
}

var lukHeritageMetaverseService = service.ServiceGroupApp.LukServiceGroup.LukHeritageMetaverseService

// CreateLukHeritageMetaverse 创建LukHeritageMetaverse
// @Tags LukHeritageMetaverse
// @Summary 创建LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageMetaverse true "创建LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverse/createLukHeritageMetaverse [post]
func (lukHeritageMetaverseApi *LukHeritageMetaverseApi) CreateLukHeritageMetaverse(c *gin.Context) {
	var lukHeritageMetaverse luk.LukHeritageMetaverse
	err := c.ShouldBindJSON(&lukHeritageMetaverse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageMetaverseService.CreateLukHeritageMetaverse(lukHeritageMetaverse); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukHeritageMetaverse 删除LukHeritageMetaverse
// @Tags LukHeritageMetaverse
// @Summary 删除LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageMetaverse true "删除LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageMetaverse/deleteLukHeritageMetaverse [delete]
func (lukHeritageMetaverseApi *LukHeritageMetaverseApi) DeleteLukHeritageMetaverse(c *gin.Context) {
	var lukHeritageMetaverse luk.LukHeritageMetaverse
	err := c.ShouldBindJSON(&lukHeritageMetaverse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageMetaverseService.DeleteLukHeritageMetaverse(lukHeritageMetaverse); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateLukHeritageMetaverse 更新LukHeritageMetaverse
// @Tags LukHeritageMetaverse
// @Summary 更新LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageMetaverse true "更新LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukHeritageMetaverse/updateLukHeritageMetaverse [put]
func (lukHeritageMetaverseApi *LukHeritageMetaverseApi) UpdateLukHeritageMetaverse(c *gin.Context) {
	var lukHeritageMetaverse luk.LukHeritageMetaverse
	err := c.ShouldBindJSON(&lukHeritageMetaverse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageMetaverseService.UpdateLukHeritageMetaverse(lukHeritageMetaverse); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukHeritageMetaverse 用id查询LukHeritageMetaverse
// @Tags LukHeritageMetaverse
// @Summary 用id查询LukHeritageMetaverse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukHeritageMetaverse true "用id查询LukHeritageMetaverse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukHeritageMetaverse/findLukHeritageMetaverse [get]
func (lukHeritageMetaverseApi *LukHeritageMetaverseApi) FindLukHeritageMetaverse(c *gin.Context) {
	var lukHeritageMetaverse luk.LukHeritageMetaverse
	err := c.ShouldBindQuery(&lukHeritageMetaverse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukHeritageMetaverse, err := lukHeritageMetaverseService.GetLukHeritageMetaverse(lukHeritageMetaverse.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukHeritageMetaverse": relukHeritageMetaverse}, c)
	}
}

// GetLukHeritageMetaverseList 分页获取LukHeritageMetaverse列表
// @Tags LukHeritageMetaverse
// @Summary 分页获取LukHeritageMetaverse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukHeritageMetaverseSearch true "分页获取LukHeritageMetaverse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverse/getLukHeritageMetaverseList [get]
func (lukHeritageMetaverseApi *LukHeritageMetaverseApi) GetLukHeritageMetaverseList(c *gin.Context) {
	var pageInfo lukReq.LukHeritageMetaverseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukHeritageMetaverseService.GetLukHeritageMetaverseInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
