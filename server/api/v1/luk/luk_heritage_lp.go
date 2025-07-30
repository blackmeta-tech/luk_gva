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

type LukHeritageLpApi struct {
}

var lukHeritageLpService = service.ServiceGroupApp.LukServiceGroup.LukHeritageLpService

// CreateLukHeritageLp 创建LukHeritageLp
// @Tags LukHeritageLp
// @Summary 创建LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageLp true "创建LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageLp/createLukHeritageLp [post]
func (lukHeritageLpApi *LukHeritageLpApi) CreateLukHeritageLp(c *gin.Context) {
	var lukHeritageLp luk.LukHeritageLp
	err := c.ShouldBindJSON(&lukHeritageLp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageLpService.CreateLukHeritageLp(lukHeritageLp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukHeritageLp 删除LukHeritageLp
// @Tags LukHeritageLp
// @Summary 删除LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukHeritageLp true "删除LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukHeritageLp/deleteLukHeritageLp [delete]
func (lukHeritageLpApi *LukHeritageLpApi) DeleteLukHeritageLp(c *gin.Context) {
	var lukHeritageLp luk.LukHeritageLp
	err := c.ShouldBindJSON(&lukHeritageLp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukHeritageLpService.DeleteLukHeritageLp(lukHeritageLp); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindLukHeritageLp 用id查询LukHeritageLp
// @Tags LukHeritageLp
// @Summary 用id查询LukHeritageLp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukHeritageLp true "用id查询LukHeritageLp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukHeritageLp/findLukHeritageLp [get]
func (lukHeritageLpApi *LukHeritageLpApi) FindLukHeritageLp(c *gin.Context) {
	var lukHeritageLp luk.LukHeritageLp
	err := c.ShouldBindQuery(&lukHeritageLp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukHeritageLp, err := lukHeritageLpService.GetLukHeritageLp(lukHeritageLp.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukHeritageLp": relukHeritageLp}, c)
	}
}

// GetLukHeritageLpList 分页获取LukHeritageLp列表
// @Tags LukHeritageLp
// @Summary 分页获取LukHeritageLp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukHeritageLpSearch true "分页获取LukHeritageLp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukHeritageLp/getLukHeritageLpList [get]
func (lukHeritageLpApi *LukHeritageLpApi) GetLukHeritageLpList(c *gin.Context) {
	var pageInfo lukReq.LukHeritageLpSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukHeritageLpService.GetLukHeritageLpInfoList(pageInfo); err != nil {
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
