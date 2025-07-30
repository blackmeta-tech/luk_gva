package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukRevenueHeritageRouter struct {
}

// InitLukRevenueHeritageRouter 初始化 LukRevenueHeritage 路由信息
func (s *LukRevenueHeritageRouter) InitLukRevenueHeritageRouter(Router *gin.RouterGroup) {
	lukRevenueRouterWithoutRecord := Router.Group("lukRevenueHeritage")
	var lukRevenueApi = v1.ApiGroupApp.LukApiGroup.LukRevenueHeritageApi
	{
		lukRevenueRouterWithoutRecord.GET("getLukRevenueHeritageList", lukRevenueApi.GetLukRevenueHeritageList) // 获取LukRevenueHeritage列表
	}
	Router.Group("apih5/lukRevenueHeritage").GET("list", lukRevenueApi.GetLukRevenueHeritageList) // 获取LukRevenueHeritage列表
}
