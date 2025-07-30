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

type LukRevenueDetailedApi struct {
}

var lukRevenueDetailedService = service.ServiceGroupApp.LukServiceGroup.LukRevenueDetailedService

// CreateLukRevenueDetailed 创建LukRevenueDetailed
// @Tags LukRevenueDetailed
// @Summary 创建LukRevenueDetailed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukRevenueDetailed true "创建LukRevenueDetailed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailed/createLukRevenueDetailed [post]
func (lukRevenueDetailedApi *LukRevenueDetailedApi) CreateLukRevenueDetailed(c *gin.Context) {
	var lukRevenueDetailed luk.LukRevenueDetailed
	_ = c.ShouldBindJSON(&lukRevenueDetailed)
	if err := lukRevenueDetailedService.CreateLukRevenueDetailed(lukRevenueDetailed); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindLukRevenueDetailed 用id查询LukRevenueDetailed
// @Tags LukRevenueDetailed
// @Summary 用id查询LukRevenueDetailed
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukRevenueDetailed true "用id查询LukRevenueDetailed"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukRevenueDetailed/findLukRevenueDetailed [get]
func (lukRevenueDetailedApi *LukRevenueDetailedApi) FindLukRevenueDetailed(c *gin.Context) {
	var lukRevenueDetailed luk.LukRevenueDetailed
	_ = c.ShouldBindQuery(&lukRevenueDetailed)
	if relukRevenueDetailed, err := lukRevenueDetailedService.GetLukRevenueDetailed(lukRevenueDetailed.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukRevenueDetailed": relukRevenueDetailed}, c)
	}
}

// GetLukRevenueDetailedList 分页获取LukRevenueDetailed列表
// @Tags LukRevenueDetailed
// @Summary 分页获取LukRevenueDetailed列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueDetailedSearch true "分页获取LukRevenueDetailed列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailed/getLukRevenueDetailedList [get]
func (lukRevenueDetailedApi *LukRevenueDetailedApi) GetLukRevenueDetailedList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueDetailedSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summary, total, err := lukRevenueDetailedService.GetLukRevenueDetailedInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]interface{}{
			"list":     list,
			"summary":  summary,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", c)
	}
}
