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

type LukNftBuyApi struct {
}

var lukNftBuyService = service.ServiceGroupApp.LukServiceGroup.LukNftBuyService

// CreateLukNftBuy 创建LukNftBuy
// @Tags LukNftBuy
// @Summary 创建LukNftBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukNftBuy true "创建LukNftBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftBuy/createLukNftBuy [post]
func (lukNftBuyApi *LukNftBuyApi) CreateLukNftBuy(c *gin.Context) {
	var lukNftBuy luk.LukNftBuy
	err := c.ShouldBindJSON(&lukNftBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukNftBuyService.CreateLukNftBuy(lukNftBuy); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindLukNftBuy 用id查询LukNftBuy
// @Tags LukNftBuy
// @Summary 用id查询LukNftBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukNftBuy true "用id查询LukNftBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukNftBuy/findLukNftBuy [get]
func (lukNftBuyApi *LukNftBuyApi) FindLukNftBuy(c *gin.Context) {
	var lukNftBuy luk.LukNftBuy
	err := c.ShouldBindQuery(&lukNftBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukNftBuy, err := lukNftBuyService.GetLukNftBuy(lukNftBuy.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukNftBuy": relukNftBuy}, c)
	}
}

// GetLukNftBuyList 分页获取LukNftBuy列表
// @Tags LukNftBuy
// @Summary 分页获取LukNftBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukNftBuySearch true "分页获取LukNftBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftBuy/getLukNftBuyList [get]
func (lukNftBuyApi *LukNftBuyApi) GetLukNftBuyList(c *gin.Context) {
	var pageInfo lukReq.LukNftBuySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukNftBuyService.GetLukNftBuyInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
