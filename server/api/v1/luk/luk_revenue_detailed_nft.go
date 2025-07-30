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

type LukRevenueDetailedNftApi struct {
}

var lukRevenueDetailedNftService = service.ServiceGroupApp.LukServiceGroup.LukRevenueDetailedNftService

// FindLukRevenueDetailedNft 用id查询LukRevenueDetailedNft
// @Tags LukRevenueDetailedNft
// @Summary 用id查询LukRevenueDetailedNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukRevenueDetailedNft true "用id查询LukRevenueDetailedNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukRevenueDetailedNft/findLukRevenueDetailedNft [get]
func (lukRevenueDetailedNftApi *LukRevenueDetailedNftApi) FindLukRevenueDetailedNft(c *gin.Context) {
	var lukRevenueDetailedNft luk.LukRevenueDetailedNft
	_ = c.ShouldBindQuery(&lukRevenueDetailedNft)
	if relukRevenueDetailedNft, err := lukRevenueDetailedNftService.GetLukRevenueDetailedNft(lukRevenueDetailedNft.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukRevenueDetailedNft": relukRevenueDetailedNft}, c)
	}
}

// GetLukRevenueDetailedNftList 分页获取LukRevenueDetailedNft列表
// @Tags LukRevenueDetailedNft
// @Summary 分页获取LukRevenueDetailedNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueDetailedNftSearch true "分页获取LukRevenueDetailedNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailedNft/getLukRevenueDetailedNftList [get]
func (lukRevenueDetailedNftApi *LukRevenueDetailedNftApi) GetLukRevenueDetailedNftList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueDetailedNftSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summary, total, err := lukRevenueDetailedNftService.GetLukRevenueDetailedNftInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
