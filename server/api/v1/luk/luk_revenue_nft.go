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

type LukRevenueNftApi struct {
}

var lukRevenueNftService = service.ServiceGroupApp.LukServiceGroup.LukRevenueNftService

// FindLukRevenueNft 用id查询LukRevenueNft
// @Tags LukRevenueNft
// @Summary 用id查询LukRevenueNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukRevenueNft true "用id查询LukRevenueNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukRevenueNft/findLukRevenueNft [get]
func (lukRevenueNftApi *LukRevenueNftApi) FindLukRevenueNft(c *gin.Context) {
	var lukRevenueNft luk.LukRevenueNft
	_ = c.ShouldBindQuery(&lukRevenueNft)
	if relukRevenueNft, err := lukRevenueNftService.GetLukRevenueNft(lukRevenueNft.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukRevenueNft": relukRevenueNft}, c)
	}
}

// GetLukRevenueNftList 分页获取LukRevenueNft列表
// @Tags LukRevenueNft
// @Summary 分页获取LukRevenueNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueNftSearch true "分页获取LukRevenueNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueNft/getLukRevenueNftList [get]
func (lukRevenueNftApi *LukRevenueNftApi) GetLukRevenueNftList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueNftSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summary, total, err := lukRevenueNftService.GetLukRevenueNftInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
