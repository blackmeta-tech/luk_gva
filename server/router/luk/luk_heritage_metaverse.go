package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukHeritageMetaverseRouter struct {
}

// InitLukHeritageMetaverseRouter 初始化 LukHeritageMetaverse 路由信息
func (s *LukHeritageMetaverseRouter) InitLukHeritageMetaverseRouter(Router *gin.RouterGroup) {
	lukHeritageMetaverseRouter := Router.Group("lukHeritageMetaverse").Use(middleware.OperationRecord())
	lukHeritageMetaverseRouterWithoutRecord := Router.Group("lukHeritageMetaverse")
	var lukHeritageMetaverseApi = v1.ApiGroupApp.LukApiGroup.LukHeritageMetaverseApi
	{
		lukHeritageMetaverseRouter.POST("createLukHeritageMetaverse", lukHeritageMetaverseApi.CreateLukHeritageMetaverse)   // 新建LukHeritageMetaverse
		lukHeritageMetaverseRouter.DELETE("deleteLukHeritageMetaverse", lukHeritageMetaverseApi.DeleteLukHeritageMetaverse) // 删除LukHeritageMetaverse
		lukHeritageMetaverseRouter.PUT("updateLukHeritageMetaverse", lukHeritageMetaverseApi.UpdateLukHeritageMetaverse)    // 更新LukHeritageMetaverse
	}
	{
		lukHeritageMetaverseRouterWithoutRecord.GET("findLukHeritageMetaverse", lukHeritageMetaverseApi.FindLukHeritageMetaverse)       // 根据ID获取LukHeritageMetaverse
		lukHeritageMetaverseRouterWithoutRecord.GET("getLukHeritageMetaverseList", lukHeritageMetaverseApi.GetLukHeritageMetaverseList) // 获取LukHeritageMetaverse列表
	}
}
