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

type LukRevenueApi struct {
}

var lukRevenueService = service.ServiceGroupApp.LukServiceGroup.LukRevenueService

// CreateLukRevenue 创建LukRevenue
// @Tags LukRevenue
// @Summary 创建LukRevenue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukRevenue true "创建LukRevenue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenue/createLukRevenue [post]
func (lukRevenueApi *LukRevenueApi) CreateLukRevenue(c *gin.Context) {
	var lukRevenue luk.LukRevenue
	_ = c.ShouldBindJSON(&lukRevenue)
	if err := lukRevenueService.CreateLukRevenue(lukRevenue); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindLukRevenue 用id查询LukRevenue
// @Tags LukRevenue
// @Summary 用id查询LukRevenue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukRevenue true "用id查询LukRevenue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukRevenue/findLukRevenue [get]
func (lukRevenueApi *LukRevenueApi) FindLukRevenue(c *gin.Context) {
	var lukRevenue luk.LukRevenue
	_ = c.ShouldBindQuery(&lukRevenue)
	if relukRevenue, err := lukRevenueService.GetLukRevenue(lukRevenue.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukRevenue": relukRevenue}, c)
	}
}

// GetLukRevenueList 分页获取LukRevenue列表
// @Tags LukRevenue
// @Summary 分页获取LukRevenue列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueSearch true "分页获取LukRevenue列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenue/getLukRevenueList [get]
func (lukRevenueApi *LukRevenueApi) GetLukRevenueList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summary, total, err := lukRevenueService.GetLukRevenueInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
