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

type LukScheduledTaskApi struct {
}

var lukScheduledTaskService = service.ServiceGroupApp.LukServiceGroup.LukScheduledTaskService

// CreateLukScheduledTask 创建LukScheduledTask
// @Tags LukScheduledTask
// @Summary 创建LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTask true "创建LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukScheduledTask/createLukScheduledTask [post]
func (lukScheduledTaskApi *LukScheduledTaskApi) CreateLukScheduledTask(c *gin.Context) {
	var lukScheduledTask luk.LukScheduledTask
	_ = c.ShouldBindJSON(&lukScheduledTask)
	if err := lukScheduledTaskService.CreateLukScheduledTask(lukScheduledTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukScheduledTask 删除LukScheduledTask
// @Tags LukScheduledTask
// @Summary 删除LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTask true "删除LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukScheduledTask/deleteLukScheduledTask [delete]
func (lukScheduledTaskApi *LukScheduledTaskApi) DeleteLukScheduledTask(c *gin.Context) {
	var lukScheduledTask luk.LukScheduledTask
	_ = c.ShouldBindJSON(&lukScheduledTask)
	if err := lukScheduledTaskService.DeleteLukScheduledTask(lukScheduledTask); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLukScheduledTaskByIds 批量删除LukScheduledTask
// @Tags LukScheduledTask
// @Summary 批量删除LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lukScheduledTask/deleteLukScheduledTaskByIds [delete]
func (lukScheduledTaskApi *LukScheduledTaskApi) DeleteLukScheduledTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lukScheduledTaskService.DeleteLukScheduledTaskByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLukScheduledTask 更新LukScheduledTask
// @Tags LukScheduledTask
// @Summary 更新LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukScheduledTask true "更新LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukScheduledTask/updateLukScheduledTask [put]
func (lukScheduledTaskApi *LukScheduledTaskApi) UpdateLukScheduledTask(c *gin.Context) {
	var lukScheduledTask luk.LukScheduledTask
	_ = c.ShouldBindJSON(&lukScheduledTask)
	if err := lukScheduledTaskService.UpdateLukScheduledTask(lukScheduledTask); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukScheduledTask 用id查询LukScheduledTask
// @Tags LukScheduledTask
// @Summary 用id查询LukScheduledTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukScheduledTask true "用id查询LukScheduledTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukScheduledTask/findLukScheduledTask [get]
func (lukScheduledTaskApi *LukScheduledTaskApi) FindLukScheduledTask(c *gin.Context) {
	var lukScheduledTask luk.LukScheduledTask
	_ = c.ShouldBindQuery(&lukScheduledTask)
	if relukScheduledTask, err := lukScheduledTaskService.GetLukScheduledTask(lukScheduledTask.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukScheduledTask": relukScheduledTask}, c)
	}
}

// GetLukScheduledTaskList 分页获取LukScheduledTask列表
// @Tags LukScheduledTask
// @Summary 分页获取LukScheduledTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukScheduledTaskSearch true "分页获取LukScheduledTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukScheduledTask/getLukScheduledTaskList [get]
func (lukScheduledTaskApi *LukScheduledTaskApi) GetLukScheduledTaskList(c *gin.Context) {
	var pageInfo lukReq.LukScheduledTaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukScheduledTaskService.GetLukScheduledTaskInfoList(pageInfo); err != nil {
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

// ExecuteImmediately 立即执行一次
// @Tags LukScheduledTask
// @Summary 立即执行一次
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukScheduledTask true "立即执行一次"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukScheduledTask/executeImmediately [get]
func (lukScheduledTaskApi *LukScheduledTaskApi) ExecuteImmediately(c *gin.Context) {
	var target luk.LukScheduledTask
	_ = c.ShouldBindQuery(&target)
	if err := lukScheduledTaskService.ExecuteImmediately(target); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}
