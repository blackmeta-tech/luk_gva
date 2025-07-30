package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukInformationTypeRouter struct {
}

// InitLukInformationTypeRouter 初始化 LukInformationType 路由信息
func (s *LukInformationTypeRouter) InitLukInformationTypeRouter(Router *gin.RouterGroup) {
	lukInformationTypeRouter := Router.Group("lukInformationType").Use(middleware.OperationRecord())
	lukInformationTypeRouterWithoutRecord := Router.Group("lukInformationType")
	var lukInformationTypeApi = v1.ApiGroupApp.LukApiGroup.LukInformationTypeApi
	{
		lukInformationTypeRouter.POST("createLukInformationType", lukInformationTypeApi.CreateLukInformationType)             // 新建LukInformationType
		lukInformationTypeRouter.DELETE("deleteLukInformationType", lukInformationTypeApi.DeleteLukInformationType)           // 删除LukInformationType
		lukInformationTypeRouter.DELETE("deleteLukInformationTypeByIds", lukInformationTypeApi.DeleteLukInformationTypeByIds) // 批量删除LukInformationType
		lukInformationTypeRouter.PUT("updateLukInformationType", lukInformationTypeApi.UpdateLukInformationType)              // 更新LukInformationType
	}
	{
		lukInformationTypeRouterWithoutRecord.GET("findLukInformationType", lukInformationTypeApi.FindLukInformationType)       // 根据ID获取LukInformationType
		lukInformationTypeRouterWithoutRecord.GET("getLukInformationTypeList", lukInformationTypeApi.GetLukInformationTypeList) // 获取LukInformationType列表
	}

	Router.Group("apih5/lukInformationType").GET("list", lukInformationTypeApi.GetLukInformationTypeList)
}
