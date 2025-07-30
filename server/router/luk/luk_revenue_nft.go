package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukRevenueNftRouter struct {
}

// InitLukRevenueNftRouter 初始化 LukRevenueNft 路由信息
func (s *LukRevenueNftRouter) InitLukRevenueNftRouter(Router *gin.RouterGroup) {
	lukRevenueRouterWithoutRecord := Router.Group("lukRevenueNft")
	var lukRevenueApi = v1.ApiGroupApp.LukApiGroup.LukRevenueNftApi
	{
		lukRevenueRouterWithoutRecord.GET("findLukRevenueNft", lukRevenueApi.FindLukRevenueNft)       // 根据ID获取LukRevenueNft
		lukRevenueRouterWithoutRecord.GET("getLukRevenueNftList", lukRevenueApi.GetLukRevenueNftList) // 获取LukRevenueNft列表
	}
	Router.Group("apih5/lukRevenueNft").GET("list", lukRevenueApi.GetLukRevenueNftList) // 获取LukRevenueNft列表
}
