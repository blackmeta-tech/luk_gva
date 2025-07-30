package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukNftBuyRouter struct {
}

// InitLukNftBuyRouter 初始化 LukNftBuy 路由信息
func (s *LukNftBuyRouter) InitLukNftBuyRouter(Router *gin.RouterGroup) {
	lukNftBuyRouterWithoutRecord := Router.Group("lukNftBuy")
	var lukNftBuyApi = v1.ApiGroupApp.LukApiGroup.LukNftBuyApi
	{
		Router.Group("apih5/lukNftBuy").POST("createLukNftBuy", lukNftBuyApi.CreateLukNftBuy) // 新建LukNftBuy
	}
	{
		lukNftBuyRouterWithoutRecord.GET("findLukNftBuy", lukNftBuyApi.FindLukNftBuy)       // 根据ID获取LukNftBuy
		lukNftBuyRouterWithoutRecord.GET("getLukNftBuyList", lukNftBuyApi.GetLukNftBuyList) // 获取LukNftBuy列表
	}
	Router.Group("apih5/lukNftBuy").GET("list", lukNftBuyApi.GetLukNftBuyList) // 获取LukNftBuy列表
}
