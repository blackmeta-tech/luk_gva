package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LukScheduledTaskLogApi struct {
}

var lukScheduledTaskLogService = service.ServiceGroupApp.LukServiceGroup.LukScheduledTaskLogService

// CreateLukScheduledTaskLog 创建LukScheduledTaskLog
// @Tags LukScheduledTaskLog
// @Summary 创建LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTaskLog true "创建LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukScheduledTaskLog/createLukScheduledTaskLog [post]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) CreateLukScheduledTaskLog(c *gin.Context) {
	var lukScheduledTaskLog luk.LukScheduledTaskLog
	_ = c.ShouldBindJSON(&lukScheduledTaskLog)
	if err := lukScheduledTaskLogService.CreateLukScheduledTaskLog(lukScheduledTaskLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukScheduledTaskLog 删除LukScheduledTaskLog
// @Tags LukScheduledTaskLog
// @Summary 删除LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTaskLog true "删除LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukScheduledTaskLog/deleteLukScheduledTaskLog [delete]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) DeleteLukScheduledTaskLog(c *gin.Context) {
	var lukScheduledTaskLog luk.LukScheduledTaskLog
	_ = c.ShouldBindJSON(&lukScheduledTaskLog)
	if err := lukScheduledTaskLogService.DeleteLukScheduledTaskLog(lukScheduledTaskLog); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLukScheduledTaskLogByIds 批量删除LukScheduledTaskLog
// @Tags LukScheduledTaskLog
// @Summary 批量删除LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lukScheduledTaskLog/deleteLukScheduledTaskLogByIds [delete]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) DeleteLukScheduledTaskLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lukScheduledTaskLogService.DeleteLukScheduledTaskLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLukScheduledTaskLog 更新LukScheduledTaskLog
// @Tags LukScheduledTaskLog
// @Summary 更新LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTaskLog true "更新LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukScheduledTaskLog/updateLukScheduledTaskLog [put]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) UpdateLukScheduledTaskLog(c *gin.Context) {
	var lukScheduledTaskLog luk.LukScheduledTaskLog
	_ = c.ShouldBindJSON(&lukScheduledTaskLog)
	if err := lukScheduledTaskLogService.UpdateLukScheduledTaskLog(lukScheduledTaskLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukScheduledTaskLog 用id查询LukScheduledTaskLog
// @Tags LukScheduledTaskLog
// @Summary 用id查询LukScheduledTaskLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukScheduledTaskLog true "用id查询LukScheduledTaskLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukScheduledTaskLog/findLukScheduledTaskLog [get]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) FindLukScheduledTaskLog(c *gin.Context) {
	var lukScheduledTaskLog luk.LukScheduledTaskLog
	_ = c.ShouldBindQuery(&lukScheduledTaskLog)
	if relukScheduledTaskLog, err := lukScheduledTaskLogService.GetLukScheduledTaskLog(lukScheduledTaskLog.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukScheduledTaskLog": relukScheduledTaskLog}, c)
	}
}

// GetLukScheduledTaskLogList 分页获取LukScheduledTaskLog列表
// @Tags LukScheduledTaskLog
// @Summary 分页获取LukScheduledTaskLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukScheduledTaskLogSearch true "分页获取LukScheduledTaskLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukScheduledTaskLog/getLukScheduledTaskLogList [get]
func (lukScheduledTaskLogApi *LukScheduledTaskLogApi) GetLukScheduledTaskLogList(c *gin.Context) {
	var pageInfo lukReq.LukScheduledTaskLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukScheduledTaskLogService.GetLukScheduledTaskLogInfoList(pageInfo); err != nil {
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
