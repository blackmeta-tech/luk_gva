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

type LukRebateApi struct {
}

var lukRebateService = service.ServiceGroupApp.LukServiceGroup.LukRebateService

// CreateLukRebate 创建LukRebate
// @Tags LukRebate
// @Summary 创建LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukRebate true "创建LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRebate/createLukRebate [post]
func (lukRebateApi *LukRebateApi) CreateLukRebate(c *gin.Context) {
	var lukRebate luk.LukRebate
	err := c.ShouldBindJSON(&lukRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukRebateService.CreateLukRebate(lukRebate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukRebate 删除LukRebate
// @Tags LukRebate
// @Summary 删除LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukRebate true "删除LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukRebate/deleteLukRebate [delete]
func (lukRebateApi *LukRebateApi) DeleteLukRebate(c *gin.Context) {
	var lukRebate luk.LukRebate
	err := c.ShouldBindJSON(&lukRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukRebateService.DeleteLukRebate(lukRebate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateLukRebate 更新LukRebate
// @Tags LukRebate
// @Summary 更新LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukRebate true "更新LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukRebate/updateLukRebate [put]
func (lukRebateApi *LukRebateApi) UpdateLukRebate(c *gin.Context) {
	var lukRebate luk.LukRebate
	err := c.ShouldBindJSON(&lukRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukRebateService.UpdateLukRebate(lukRebate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukRebate 用id查询LukRebate
// @Tags LukRebate
// @Summary 用id查询LukRebate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukRebate true "用id查询LukRebate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukRebate/findLukRebate [get]
func (lukRebateApi *LukRebateApi) FindLukRebate(c *gin.Context) {
	var lukRebate luk.LukRebate
	err := c.ShouldBindQuery(&lukRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukRebate, err := lukRebateService.GetLukRebate(lukRebate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukRebate": relukRebate}, c)
	}
}

// GetLukRebateList 分页获取LukRebate列表
// @Tags LukRebate
// @Summary 分页获取LukRebate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRebateSearch true "分页获取LukRebate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRebate/getLukRebateList [get]
func (lukRebateApi *LukRebateApi) GetLukRebateList(c *gin.Context) {
	var pageInfo lukReq.LukRebateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukRebateService.GetLukRebateInfoList(pageInfo); err != nil {
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

func (lukRebateApi *LukRebateApi) QueryLukRebate(c *gin.Context) {
	var lukRebate luk.LukRebate
	err := c.ShouldBindQuery(&lukRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	relukRebate, _ := lukRebateService.QueryLukRebateByAddress(lukRebate.Address)
	response.OkWithData(relukRebate, c)
}
