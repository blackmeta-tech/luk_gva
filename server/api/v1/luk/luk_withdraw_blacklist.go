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

type LukWithdrawBlacklistApi struct {
}

var lukWithdrawBlacklistService = service.ServiceGroupApp.LukServiceGroup.LukWithdrawBlacklistService

// CreateLukWithdrawBlacklist 创建LukWithdrawBlacklist
// @Tags LukWithdrawBlacklist
// @Summary 创建LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawBlacklist true "创建LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawBlacklist/createLukWithdrawBlacklist [post]
func (lukWithdrawBlacklistApi *LukWithdrawBlacklistApi) CreateLukWithdrawBlacklist(c *gin.Context) {
	var lukWithdrawBlacklist luk.LukWithdrawBlacklist
	err := c.ShouldBindJSON(&lukWithdrawBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukWithdrawBlacklistService.CreateLukWithdrawBlacklist(lukWithdrawBlacklist); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败，"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukWithdrawBlacklist 删除LukWithdrawBlacklist
// @Tags LukWithdrawBlacklist
// @Summary 删除LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawBlacklist true "删除LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukWithdrawBlacklist/deleteLukWithdrawBlacklist [delete]
func (lukWithdrawBlacklistApi *LukWithdrawBlacklistApi) DeleteLukWithdrawBlacklist(c *gin.Context) {
	var lukWithdrawBlacklist luk.LukWithdrawBlacklist
	err := c.ShouldBindJSON(&lukWithdrawBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukWithdrawBlacklistService.DeleteLukWithdrawBlacklist(lukWithdrawBlacklist); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateLukWithdrawBlacklist 更新LukWithdrawBlacklist
// @Tags LukWithdrawBlacklist
// @Summary 更新LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawBlacklist true "更新LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukWithdrawBlacklist/updateLukWithdrawBlacklist [put]
func (lukWithdrawBlacklistApi *LukWithdrawBlacklistApi) UpdateLukWithdrawBlacklist(c *gin.Context) {
	var lukWithdrawBlacklist luk.LukWithdrawBlacklist
	err := c.ShouldBindJSON(&lukWithdrawBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukWithdrawBlacklistService.UpdateLukWithdrawBlacklist(lukWithdrawBlacklist); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukWithdrawBlacklist 用id查询LukWithdrawBlacklist
// @Tags LukWithdrawBlacklist
// @Summary 用id查询LukWithdrawBlacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukWithdrawBlacklist true "用id查询LukWithdrawBlacklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukWithdrawBlacklist/findLukWithdrawBlacklist [get]
func (lukWithdrawBlacklistApi *LukWithdrawBlacklistApi) FindLukWithdrawBlacklist(c *gin.Context) {
	var lukWithdrawBlacklist luk.LukWithdrawBlacklist
	err := c.ShouldBindQuery(&lukWithdrawBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukWithdrawBlacklist, err := lukWithdrawBlacklistService.GetLukWithdrawBlacklist(lukWithdrawBlacklist.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukWithdrawBlacklist": relukWithdrawBlacklist}, c)
	}
}

// GetLukWithdrawBlacklistList 分页获取LukWithdrawBlacklist列表
// @Tags LukWithdrawBlacklist
// @Summary 分页获取LukWithdrawBlacklist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukWithdrawBlacklistSearch true "分页获取LukWithdrawBlacklist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawBlacklist/getLukWithdrawBlacklistList [get]
func (lukWithdrawBlacklistApi *LukWithdrawBlacklistApi) GetLukWithdrawBlacklistList(c *gin.Context) {
	var pageInfo lukReq.LukWithdrawBlacklistSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukWithdrawBlacklistService.GetLukWithdrawBlacklistInfoList(pageInfo); err != nil {
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
