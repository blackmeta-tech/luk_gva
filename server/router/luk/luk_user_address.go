package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukUserAddressRouter struct {
}

// InitLukUserAddressRouter 初始化 LukUserAddress 路由信息
func (s *LukUserAddressRouter) InitLukUserAddressRouter(Router *gin.RouterGroup) {
	lukUserAddressRouter := Router.Group("lukUserAddress").Use(middleware.OperationRecord())
	lukUserAddressRouterWithoutRecord := Router.Group("lukUserAddress")
	var lukUserAddressApi = v1.ApiGroupApp.LukApiGroup.LukUserAddressApi
	{
		lukUserAddressRouter.POST("updateLukUserAddress", lukUserAddressApi.UpdateLukUserAddress) // 更新LukUserAddress
	}
	{
		lukUserAddressRouterWithoutRecord.GET("getLukUserAddressList", lukUserAddressApi.GetLukUserAddressList) // 获取LukUserAddress列表
	}
	Router.Group("apih5/lukUserAddress").GET("findLukUserAddress", lukUserAddressApi.FindLukUserAddress) // 获取LukUserAddress列表
}
