package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukInformationApi struct {
}

var lukDappInformationService = service.ServiceGroupApp.LukServiceGroup.LukInformationService

// CreateLukInformation 创建LukInformation
// @Tags LukInformation
// @Summary 创建LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformation true "创建LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukDappInformation/createLukInformation [post]
func (lukDappInformationApi *LukInformationApi) CreateLukInformation(c *gin.Context) {
	var lukDappInformation luk.LukInformation
	_ = c.ShouldBindJSON(&lukDappInformation)
	if err := lukDappInformationService.CreateLukInformation(lukDappInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukInformation 删除LukInformation
// @Tags LukInformation
// @Summary 删除LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformation true "删除LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lukDappInformation/deleteLukInformation [delete]
func (lukDappInformationApi *LukInformationApi) DeleteLukInformation(c *gin.Context) {
	var lukDappInformation luk.LukInformation
	_ = c.ShouldBindJSON(&lukDappInformation)
	//查出id对应的key
	old, _ := lukDappInformationService.GetLukInformation(lukDappInformation.ID)
	if err := lukDappInformationService.DeleteLukInformation(lukDappInformation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		//需将图片删除，以免浪费内存
		if old.PicKey != "" {
			if err = service.ServiceGroupApp.LukServiceGroup.FileUploadService.DeleteFile(old.PicKey); err != nil {
				global.GVA_LOG.Error("删除图片失败!", zap.Error(err))
			}
		}
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateLukInformation 更新LukInformation
// @Tags LukInformation
// @Summary 更新LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukInformation true "更新LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukDappInformation/updateLukInformation [put]
func (lukDappInformationApi *LukInformationApi) UpdateLukInformation(c *gin.Context) {
	var lukDappInformation luk.LukInformation
	_ = c.ShouldBindJSON(&lukDappInformation)
	//查出id对应的key
	old, _ := lukDappInformationService.GetLukInformation(lukDappInformation.ID)
	if err := lukDappInformationService.UpdateLukInformation(lukDappInformation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		//更新图片需将旧图删除，以免浪费内存
		if old.PicKey != "" && old.PicKey != lukDappInformation.PicKey {
			if err = service.ServiceGroupApp.LukServiceGroup.FileUploadService.DeleteFile(old.PicKey); err != nil {
				global.GVA_LOG.Error("删除图片失败!", zap.Error(err))
			}
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukInformation 用id查询LukInformation
// @Tags LukInformation
// @Summary 用id查询LukInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukInformation true "用id查询LukInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukDappInformation/findLukInformation [get]
func (lukDappInformationApi *LukInformationApi) FindLukInformation(c *gin.Context) {
	var lukDappInformation luk.LukInformation
	_ = c.ShouldBindQuery(&lukDappInformation)
	if relukDappInformation, err := lukDappInformationService.GetLukInformation(lukDappInformation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukDappInformation": relukDappInformation}, c)
	}
}

// GetLukInformationList 分页获取LukInformation列表
// @Tags LukInformation
// @Summary 分页获取LukInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukInformationSearch true "分页获取LukInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukDappInformation/getLukInformationList [get]
func (lukDappInformationApi *LukInformationApi) GetLukInformationList(c *gin.Context) {
	var pageInfo lukReq.LukInformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukDappInformationService.GetLukInformationInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
