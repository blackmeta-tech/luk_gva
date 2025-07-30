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

type LukRevenueHeritageApi struct {
}

var lukRevenueHeritageService = service.ServiceGroupApp.LukServiceGroup.LukRevenueHeritageService

// GetLukRevenueHeritageList 分页获取LukRevenueHeritage列表
// @Tags LukRevenueHeritage
// @Summary 分页获取LukRevenueHeritage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukRevenueHeritageSearch true "分页获取LukRevenueHeritage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukRevenueHeritage/getLukRevenueHeritageList [get]
func (lukRevenueHeritageApi *LukRevenueHeritageApi) GetLukRevenueHeritageList(c *gin.Context) {
	var pageInfo lukReq.LukRevenueHeritageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summaryLuk, summaryUsdt, total, err := lukRevenueHeritageService.GetLukRevenueHeritageInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
