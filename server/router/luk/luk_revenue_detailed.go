package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukRevenueDetailedRouter struct {
}

// InitLukRevenueDetailedRouter 初始化 LukRevenueDetailed 路由信息
func (s *LukRevenueDetailedRouter) InitLukRevenueDetailedRouter(Router *gin.RouterGroup) {
	lukRevenueDetailedRouter := Router.Group("lukRevenueDetailed").Use(middleware.OperationRecord())
	lukRevenueDetailedRouterWithoutRecord := Router.Group("lukRevenueDetailed")
	var lukRevenueDetailedApi = v1.ApiGroupApp.LukApiGroup.LukRevenueDetailedApi
	{
		lukRevenueDetailedRouter.POST("createLukRevenueDetailed", lukRevenueDetailedApi.CreateLukRevenueDetailed) // 新建LukRevenueDetailed
	}
	{
		lukRevenueDetailedRouterWithoutRecord.GET("findLukRevenueDetailed", lukRevenueDetailedApi.FindLukRevenueDetailed)       // 根据ID获取LukRevenueDetailed
		lukRevenueDetailedRouterWithoutRecord.GET("getLukRevenueDetailedList", lukRevenueDetailedApi.GetLukRevenueDetailedList) // 获取LukRevenueDetailed列表
	}
	Router.Group("apih5/lukRevenueDetailed").GET("list", lukRevenueDetailedApi.GetLukRevenueDetailedList) // 获取LukRevenueDetailed列表
}
