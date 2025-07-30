package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukRevenueRouter struct {
}

// InitLukRevenueRouter 初始化 LukRevenue 路由信息
func (s *LukRevenueRouter) InitLukRevenueRouter(Router *gin.RouterGroup) {
	lukRevenueRouter := Router.Group("lukRevenue").Use(middleware.OperationRecord())
	lukRevenueRouterWithoutRecord := Router.Group("lukRevenue")
	var lukRevenueApi = v1.ApiGroupApp.LukApiGroup.LukRevenueApi
	{
		lukRevenueRouter.POST("createLukRevenue", lukRevenueApi.CreateLukRevenue) // 新建LukRevenue
	}
	{
		lukRevenueRouterWithoutRecord.GET("findLukRevenue", lukRevenueApi.FindLukRevenue)       // 根据ID获取LukRevenue
		lukRevenueRouterWithoutRecord.GET("getLukRevenueList", lukRevenueApi.GetLukRevenueList) // 获取LukRevenue列表
	}
	Router.Group("apih5/lukRevenue").GET("list", lukRevenueApi.GetLukRevenueList) // 获取LukRevenue列表
}
