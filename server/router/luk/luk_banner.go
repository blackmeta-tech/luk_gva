package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukBannerRouter struct {
}

// InitLukBannerRouter 初始化 LukBanner 路由信息
func (s *LukBannerRouter) InitLukBannerRouter(Router *gin.RouterGroup) {
	lukBannerRouter := Router.Group("lukBanner").Use(middleware.OperationRecord())
	lukBannerRouterWithoutRecord := Router.Group("lukBanner")
	var lukBannerApi = v1.ApiGroupApp.LukApiGroup.LukBannerApi
	{
		lukBannerRouter.POST("createLukBanner", lukBannerApi.CreateLukBanner)   // 新建LukBanner
		lukBannerRouter.DELETE("deleteLukBanner", lukBannerApi.DeleteLukBanner) // 删除LukBanner
		lukBannerRouter.PUT("updateLukBanner", lukBannerApi.UpdateLukBanner)    // 更新LukBanner
	}
	{
		lukBannerRouterWithoutRecord.GET("findLukBanner", lukBannerApi.FindLukBanner)       // 根据ID获取LukBanner
		lukBannerRouterWithoutRecord.GET("getLukBannerList", lukBannerApi.GetLukBannerList) // 获取LukBanner列表
	}
	Router.Group("apih5/lukBanner").GET("list", lukBannerApi.GetLukBannerList) //前端访问
}
