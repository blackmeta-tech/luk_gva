package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukScheduledTaskRouter struct {
}

// InitLukScheduledTaskRouter 初始化 LukScheduledTask 路由信息
func (s *LukScheduledTaskRouter) InitLukScheduledTaskRouter(Router *gin.RouterGroup) {
	lukScheduledTaskRouter := Router.Group("lukScheduledTask").Use(middleware.OperationRecord())
	lukScheduledTaskRouterWithoutRecord := Router.Group("lukScheduledTask")
	var lukScheduledTaskApi = v1.ApiGroupApp.LukApiGroup.LukScheduledTaskApi
	{
		lukScheduledTaskRouter.POST("createLukScheduledTask", lukScheduledTaskApi.CreateLukScheduledTask)             // 新建LukScheduledTask
		lukScheduledTaskRouter.DELETE("deleteLukScheduledTask", lukScheduledTaskApi.DeleteLukScheduledTask)           // 删除LukScheduledTask
		lukScheduledTaskRouter.DELETE("deleteLukScheduledTaskByIds", lukScheduledTaskApi.DeleteLukScheduledTaskByIds) // 批量删除LukScheduledTask
		lukScheduledTaskRouter.PUT("updateLukScheduledTask", lukScheduledTaskApi.UpdateLukScheduledTask)              // 更新LukScheduledTask
	}
	{
		lukScheduledTaskRouterWithoutRecord.GET("findLukScheduledTask", lukScheduledTaskApi.FindLukScheduledTask)       // 根据ID获取LukScheduledTask
		lukScheduledTaskRouterWithoutRecord.GET("getLukScheduledTaskList", lukScheduledTaskApi.GetLukScheduledTaskList) // 获取LukScheduledTask列表
		lukScheduledTaskRouterWithoutRecord.GET("executeImmediately", lukScheduledTaskApi.ExecuteImmediately)           // 立即执行一次
	}
}
