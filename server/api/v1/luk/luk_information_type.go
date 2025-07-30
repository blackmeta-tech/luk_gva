package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukInformationTypeApi struct {
}

var lukDappInformationTypeService = service.ServiceGroupApp.LukServiceGroup.LukInformationTypeService

// CreateLukInformationType 创建LukInformationType
// @Tags LukInformationType
// @Summary 创建LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformationType true "创建LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukDappInformationType/createLukInformationType [post]
func (lukDappInformationTypeApi *LukInformationTypeApi) CreateLukInformationType(c *gin.Context) {
	var lukDappInformationType luk.LukInformationType
	_ = c.ShouldBindJSON(&lukDappInformationType)
	if err := lukDappInformationTypeService.CreateLukInformationType(lukDappInformationType); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukInformationType 删除LukInformationType
// @Tags LukInformationType
// @Summary 删除LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformationType true "删除LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukDappInformationType/deleteLukInformationType [delete]
func (lukDappInformationTypeApi *LukInformationTypeApi) DeleteLukInformationType(c *gin.Context) {
	var lukDappInformationType luk.LukInformationType
	_ = c.ShouldBindJSON(&lukDappInformationType)
	if err := lukDappInformationTypeService.DeleteLukInformationType(lukDappInformationType); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLukInformationTypeByIds 批量删除LukInformationType
// @Tags LukInformationType
// @Summary 批量删除LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lukDappInformationType/deleteLukInformationTypeByIds [delete]
func (lukDappInformationTypeApi *LukInformationTypeApi) DeleteLukInformationTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lukDappInformationTypeService.DeleteLukInformationTypeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLukInformationType 更新LukInformationType
// @Tags LukInformationType
// @Summary 更新LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformationType true "更新LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukDappInformationType/updateLukInformationType [put]
func (lukDappInformationTypeApi *LukInformationTypeApi) UpdateLukInformationType(c *gin.Context) {
	var lukDappInformationType luk.LukInformationType
	_ = c.ShouldBindJSON(&lukDappInformationType)
	if err := lukDappInformationTypeService.UpdateLukInformationType(lukDappInformationType); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukInformationType 用id查询LukInformationType
// @Tags LukInformationType
// @Summary 用id查询LukInformationType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukInformationType true "用id查询LukInformationType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukDappInformationType/findLukInformationType [get]
func (lukDappInformationTypeApi *LukInformationTypeApi) FindLukInformationType(c *gin.Context) {
	var lukDappInformationType luk.LukInformationType
	_ = c.ShouldBindQuery(&lukDappInformationType)
	if relukDappInformationType, err := lukDappInformationTypeService.GetLukInformationType(lukDappInformationType.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukDappInformationType": relukDappInformationType}, c)
	}
}

// GetLukInformationTypeList 分页获取LukInformationType列表
// @Tags LukInformationType
// @Summary 分页获取LukInformationType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukInformationTypeSearch true "分页获取LukInformationType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukDappInformationType/getLukInformationTypeList [get]
func (lukDappInformationTypeApi *LukInformationTypeApi) GetLukInformationTypeList(c *gin.Context) {
	var pageInfo lukReq.LukInformationTypeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukDappInformationTypeService.GetLukInformationTypeInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
