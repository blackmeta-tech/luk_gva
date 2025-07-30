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

type LukBannerApi struct {
}

var wLukBannerService = service.ServiceGroupApp.LukServiceGroup.LukBannerService

// CreateLukBanner 创建LukBanner
// @Tags LukBanner
// @Summary 创建LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukBanner true "创建LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wLukBanner/createLukBanner [post]
func (wLukBannerApi *LukBannerApi) CreateLukBanner(c *gin.Context) {
	var wLukBanner luk.LukBanner
	_ = c.ShouldBindJSON(&wLukBanner)
	if err := wLukBannerService.CreateLukBanner(wLukBanner); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLukBanner 删除LukBanner
// @Tags LukBanner
// @Summary 删除LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukBanner true "删除LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wLukBanner/deleteLukBanner [delete]
func (wLukBannerApi *LukBannerApi) DeleteLukBanner(c *gin.Context) {
	var wLukBanner luk.LukBanner
	_ = c.ShouldBindJSON(&wLukBanner)
	//查出id对应的key
	info, _ := wLukBannerService.GetLukBanner(wLukBanner.ID)
	if err := wLukBannerService.DeleteLukBanner(wLukBanner); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		//需将图片删除，以免浪费内存
		if info.PicKey != "" {
			if err = service.ServiceGroupApp.LukServiceGroup.FileUploadService.DeleteFile(info.PicKey); err != nil {
				global.GVA_LOG.Error("删除图片失败!", zap.Error(err))
			}
		}
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateLukBanner 更新LukBanner
// @Tags LukBanner
// @Summary 更新LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukBanner true "更新LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wLukBanner/updateLukBanner [put]
func (wLukBannerApi *LukBannerApi) UpdateLukBanner(c *gin.Context) {
	var wLukBanner luk.LukBanner
	_ = c.ShouldBindJSON(&wLukBanner)
	//查出id对应的key
	old, _ := wLukBannerService.GetLukBanner(wLukBanner.ID)
	if err := wLukBannerService.UpdateLukBanner(wLukBanner); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		//更新图片需将旧图删除，以免浪费内存
		if old.PicKey != "" && old.PicKey != wLukBanner.PicKey {
			if err = service.ServiceGroupApp.LukServiceGroup.FileUploadService.DeleteFile(old.PicKey); err != nil {
				global.GVA_LOG.Error("删除图片失败!", zap.Error(err))
			}
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukBanner 用id查询LukBanner
// @Tags LukBanner
// @Summary 用id查询LukBanner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukBanner true "用id查询LukBanner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wLukBanner/findLukBanner [get]
func (wLukBannerApi *LukBannerApi) FindLukBanner(c *gin.Context) {
	var wLukBanner luk.LukBanner
	_ = c.ShouldBindQuery(&wLukBanner)
	if rewLukBanner, err := wLukBannerService.GetLukBanner(wLukBanner.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewLukBanner": rewLukBanner}, c)
	}
}

// GetLukBannerList 分页获取LukBanner列表
// @Tags LukBanner
// @Summary 分页获取LukBanner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukBannerSearch true "分页获取LukBanner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wLukBanner/getLukBannerList [get]
func (wLukBannerApi *LukBannerApi) GetLukBannerList(c *gin.Context) {
	var pageInfo lukReq.LukBannerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := wLukBannerService.GetLukBannerInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
