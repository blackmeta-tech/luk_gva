package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukUserRouter struct {
}

// InitLukUserRouter 初始化 LukUser 路由信息
func (s *LukUserRouter) InitLukUserRouter(Router *gin.RouterGroup) {
	lukUserRouterWithoutRecord := Router.Group("lukUser")
	var lukUserApi = v1.ApiGroupApp.LukApiGroup.LukUserApi
	{
		Router.Group("apih5/lukUser").POST("create", lukUserApi.CreateLukUser)                                       // 新建LukUser
		lukUserRouterWithoutRecord.Use(middleware.OperationRecord()).POST("updateLinkage", lukUserApi.UpdateLinkage) // 编辑
		lukUserRouterWithoutRecord.Use(middleware.OperationRecord()).POST("updateLukUser", lukUserApi.UpdateLukUser) // 编辑
	}
	{
		lukUserRouterWithoutRecord.GET("findLukUser", lukUserApi.FindLukUser)       // 根据ID获取LukUser
		lukUserRouterWithoutRecord.GET("getLukUserList", lukUserApi.GetLukUserList) // 获取LukUser列表
	}
	Router.Group("apih5/lukUser").GET("query", lukUserApi.QueryUser)        // 获取LukUser列表
	Router.Group("apih5/lukUser").GET("subclass", lukUserApi.QuerySubclass) // 获取LukUser列表
}
