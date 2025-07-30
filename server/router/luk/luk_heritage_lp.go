package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukHeritageLpRouter struct {
}

// InitLukHeritageLpRouter 初始化 LukHeritageLp 路由信息
func (s *LukHeritageLpRouter) InitLukHeritageLpRouter(Router *gin.RouterGroup) {
	lukHeritageLpRouter := Router.Group("lukHeritageLp").Use(middleware.OperationRecord())
	lukHeritageLpRouterWithoutRecord := Router.Group("lukHeritageLp")
	var lukHeritageLpApi = v1.ApiGroupApp.LukApiGroup.LukHeritageLpApi
	{
		lukHeritageLpRouter.POST("createLukHeritageLp", lukHeritageLpApi.CreateLukHeritageLp)   // 新建LukHeritageLp
		lukHeritageLpRouter.DELETE("deleteLukHeritageLp", lukHeritageLpApi.DeleteLukHeritageLp) // 删除LukHeritageLp
	}
	{
		lukHeritageLpRouterWithoutRecord.GET("findLukHeritageLp", lukHeritageLpApi.FindLukHeritageLp)       // 根据ID获取LukHeritageLp
		lukHeritageLpRouterWithoutRecord.GET("getLukHeritageLpList", lukHeritageLpApi.GetLukHeritageLpList) // 获取LukHeritageLp列表
	}
}
