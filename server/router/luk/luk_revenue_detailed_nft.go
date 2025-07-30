package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukRevenueDetailedNftRouter struct {
}

// InitLukRevenueDetailedNftRouter 初始化 LukRevenueDetailedNft 路由信息
func (s *LukRevenueDetailedNftRouter) InitLukRevenueDetailedNftRouter(Router *gin.RouterGroup) {
	lukRevenueDetailedRouterWithoutRecord := Router.Group("lukRevenueDetailedNft")
	var lukRevenueDetailedApi = v1.ApiGroupApp.LukApiGroup.LukRevenueDetailedNftApi
	{
		lukRevenueDetailedRouterWithoutRecord.GET("findLukRevenueDetailedNft", lukRevenueDetailedApi.FindLukRevenueDetailedNft)       // 根据ID获取LukRevenueDetailedNft
		lukRevenueDetailedRouterWithoutRecord.GET("getLukRevenueDetailedNftList", lukRevenueDetailedApi.GetLukRevenueDetailedNftList) // 获取LukRevenueDetailedNft列表
	}
	Router.Group("apih5/lukRevenueDetailedNft").GET("list", lukRevenueDetailedApi.GetLukRevenueDetailedNftList) // 获取LukRevenueDetailedNft列表
}
