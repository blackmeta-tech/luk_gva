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

type LukHeritageMetaverseOldApi struct {
}

var lukHeritageMetaverseOldService = service.ServiceGroupApp.LukServiceGroup.LukHeritageMetaverseOldService

// CreateLukHeritageMetaverseOld 创建LukHeritageMetaverseOld
// @Tags LukHeritageMetaverseOld
// @Summary 创建LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageMetaverseOld true "创建LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverseOld/createLukHeritageMetaverseOld [post]
func (lukHeritageMetaverseOldApi *LukHeritageMetaverseOldApi) CreateLukHeritageMetaverseOld(c *gin.Context) {
	var lukHeritageMetaverseOld luk.LukHeritageMetaverseOld
	err := c.ShouldBindJSON(&lukHeritageMetaverseOld)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageMetaverseOldService.CreateLukHeritageMetaverseOld(lukHeritageMetaverseOld); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukHeritageMetaverseOld 删除LukHeritageMetaverseOld
// @Tags LukHeritageMetaverseOld
// @Summary 删除LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageMetaverseOld true "删除LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageMetaverseOld/deleteLukHeritageMetaverseOld [delete]
func (lukHeritageMetaverseOldApi *LukHeritageMetaverseOldApi) DeleteLukHeritageMetaverseOld(c *gin.Context) {
	var lukHeritageMetaverseOld luk.LukHeritageMetaverseOld
	err := c.ShouldBindJSON(&lukHeritageMetaverseOld)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageMetaverseOldService.DeleteLukHeritageMetaverseOld(lukHeritageMetaverseOld); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindLukHeritageMetaverseOld 用id查询LukHeritageMetaverseOld
// @Tags LukHeritageMetaverseOld
// @Summary 用id查询LukHeritageMetaverseOld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukHeritageMetaverseOld true "用id查询LukHeritageMetaverseOld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukHeritageMetaverseOld/findLukHeritageMetaverseOld [get]
func (lukHeritageMetaverseOldApi *LukHeritageMetaverseOldApi) FindLukHeritageMetaverseOld(c *gin.Context) {
	var lukHeritageMetaverseOld luk.LukHeritageMetaverseOld
	err := c.ShouldBindQuery(&lukHeritageMetaverseOld)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukHeritageMetaverseOld, err := lukHeritageMetaverseOldService.GetLukHeritageMetaverseOld(lukHeritageMetaverseOld.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukHeritageMetaverseOld": relukHeritageMetaverseOld}, c)
	}
}

// GetLukHeritageMetaverseOldList 分页获取LukHeritageMetaverseOld列表
// @Tags LukHeritageMetaverseOld
// @Summary 分页获取LukHeritageMetaverseOld列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukHeritageMetaverseOldSearch true "分页获取LukHeritageMetaverseOld列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageMetaverseOld/getLukHeritageMetaverseOldList [get]
func (lukHeritageMetaverseOldApi *LukHeritageMetaverseOldApi) GetLukHeritageMetaverseOldList(c *gin.Context) {
	var pageInfo lukReq.LukHeritageMetaverseOldSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukHeritageMetaverseOldService.GetLukHeritageMetaverseOldInfoList(pageInfo); err != nil {
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
