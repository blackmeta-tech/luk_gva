package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukRevenueDetailedHeritageRouter struct {
}

// InitLukRevenueDetailedHeritageRouter 初始化 LukRevenueDetailedHeritage 路由信息
func (s *LukRevenueDetailedHeritageRouter) InitLukRevenueDetailedHeritageRouter(Router *gin.RouterGroup) {
	lukRevenueDetailedRouterWithoutRecord := Router.Group("lukRevenueDetailedHeritage")
	var lukRevenueDetailedApi = v1.ApiGroupApp.LukApiGroup.LukRevenueDetailedHeritageApi
	{
		lukRevenueDetailedRouterWithoutRecord.GET("getLukRevenueDetailedHeritageList", lukRevenueDetailedApi.GetLukRevenueDetailedHeritageList) // 获取LukRevenueDetailedHeritage列表
	}
	Router.Group("apih5/lukRevenueDetailedHeritage").GET("list", lukRevenueDetailedApi.GetLukRevenueDetailedHeritageList) // 获取LukRevenueDetailedHeritage列表
}
