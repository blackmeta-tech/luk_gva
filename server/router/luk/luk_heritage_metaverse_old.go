package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukHeritageMetaverseOldRouter struct {
}

// InitLukHeritageMetaverseOldRouter 初始化 LukHeritageMetaverseOld 路由信息
func (s *LukHeritageMetaverseOldRouter) InitLukHeritageMetaverseOldRouter(Router *gin.RouterGroup) {
	lukHeritageMetaverseOldRouter := Router.Group("lukHeritageMetaverseOld").Use(middleware.OperationRecord())
	lukHeritageMetaverseOldRouterWithoutRecord := Router.Group("lukHeritageMetaverseOld")
	var lukHeritageMetaverseOldApi = v1.ApiGroupApp.LukApiGroup.LukHeritageMetaverseOldApi
	{
		lukHeritageMetaverseOldRouter.POST("createLukHeritageMetaverseOld", lukHeritageMetaverseOldApi.CreateLukHeritageMetaverseOld)   // 新建LukHeritageMetaverseOld
		lukHeritageMetaverseOldRouter.DELETE("deleteLukHeritageMetaverseOld", lukHeritageMetaverseOldApi.DeleteLukHeritageMetaverseOld) // 删除LukHeritageMetaverseOld
	}
	{
		lukHeritageMetaverseOldRouterWithoutRecord.GET("findLukHeritageMetaverseOld", lukHeritageMetaverseOldApi.FindLukHeritageMetaverseOld)       // 根据ID获取LukHeritageMetaverseOld
		lukHeritageMetaverseOldRouterWithoutRecord.GET("getLukHeritageMetaverseOldList", lukHeritageMetaverseOldApi.GetLukHeritageMetaverseOldList) // 获取LukHeritageMetaverseOld列表
	}
}
