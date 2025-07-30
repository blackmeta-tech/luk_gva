package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukRevenueDetailedHeritageApi struct {
}

var lukRevenueDetailedHeritageService = service.ServiceGroupApp.LukServiceGroup.LukRevenueDetailedHeritageService

// GetLukRevenueDetailedHeritageList 分页获取LukRevenueDetailedHeritage列表
// @Tags LukRevenueDetailedHeritage
// @Summary 分页获取LukRevenueDetailedHeritage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueDetailedHeritageSearch true "分页获取LukRevenueDetailedHeritage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueDetailedHeritage/getLukRevenueDetailedHeritageList [get]
func (lukRevenueDetailedHeritageApi *LukRevenueDetailedHeritageApi) GetLukRevenueDetailedHeritageList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueDetailedHeritageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summaryLuk, summaryUsdt, total, err := lukRevenueDetailedHeritageService.GetLukRevenueDetailedHeritageInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]interface{}{
			"list":        list,
			"summaryLuk":  summaryLuk,
			"summaryUsdt": summaryUsdt,
			"total":       total,
			"page":        pageInfo.Page,
			"pageSize":    pageInfo.PageSize,
		}, "获取成功", c)
	}
}
