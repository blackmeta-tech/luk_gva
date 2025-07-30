package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukScheduledTaskLogRouter struct {
}

// InitLukScheduledTaskLogRouter 初始化 LukScheduledTaskLog 路由信息
func (s *LukScheduledTaskLogRouter) InitLukScheduledTaskLogRouter(Router *gin.RouterGroup) {
	lukScheduledTaskLogRouter := Router.Group("lukScheduledTaskLog").Use(middleware.OperationRecord())
	lukScheduledTaskLogRouterWithoutRecord := Router.Group("lukScheduledTaskLog")
	var lukScheduledTaskLogApi = v1.ApiGroupApp.LukApiGroup.LukScheduledTaskLogApi
	{
		lukScheduledTaskLogRouter.POST("createLukScheduledTaskLog", lukScheduledTaskLogApi.CreateLukScheduledTaskLog)             // 新建LukScheduledTaskLog
		lukScheduledTaskLogRouter.DELETE("deleteLukScheduledTaskLog", lukScheduledTaskLogApi.DeleteLukScheduledTaskLog)           // 删除LukScheduledTaskLog
		lukScheduledTaskLogRouter.DELETE("deleteLukScheduledTaskLogByIds", lukScheduledTaskLogApi.DeleteLukScheduledTaskLogByIds) // 批量删除LukScheduledTaskLog
		lukScheduledTaskLogRouter.PUT("updateLukScheduledTaskLog", lukScheduledTaskLogApi.UpdateLukScheduledTaskLog)              // 更新LukScheduledTaskLog
	}
	{
		lukScheduledTaskLogRouterWithoutRecord.GET("findLukScheduledTaskLog", lukScheduledTaskLogApi.FindLukScheduledTaskLog)       // 根据ID获取LukScheduledTaskLog
		lukScheduledTaskLogRouterWithoutRecord.GET("getLukScheduledTaskLogList", lukScheduledTaskLogApi.GetLukScheduledTaskLogList) // 获取LukScheduledTaskLog列表
	}
}
