package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukNftRouter struct {
}

// InitLukNftRouter 初始化 LukNft 路由信息
func (s *LukNftRouter) InitLukNftRouter(Router *gin.RouterGroup) {
	lukNftRouter := Router.Group("lukNft").Use(middleware.OperationRecord())
	lukNftRouterWithoutRecord := Router.Group("lukNft")
	var lukNftApi = v1.ApiGroupApp.LukApiGroup.LukNftApi
	{
		lukNftRouter.POST("createLukNft", lukNftApi.CreateLukNft)          // 新建LukNft
		lukNftRouter.PUT("updateLukNft", lukNftApi.UpdateLukNft)           // 更新LukNft
		lukNftRouter.PUT("updateLukNftBlack", lukNftApi.UpdateLukNftBlack) // 更新LukNft
	}
	{
		lukNftRouterWithoutRecord.GET("findLukNft", lukNftApi.FindLukNft)       // 根据ID获取LukNft
		lukNftRouterWithoutRecord.GET("getLukNftList", lukNftApi.GetLukNftList) // 获取LukNft列表
	}
	Router.Group("apih5/lukNft").GET("getLukNftList", lukNftApi.GetLukNftList) // 获取LukNft列表
	Router.Group("apih5/lukNft").GET("encryption", lukNftApi.BuyLukNft)        // 加密
	Router.Group("apih5/lukNft").GET("summary", lukNftApi.Summary)             // 加密
}
