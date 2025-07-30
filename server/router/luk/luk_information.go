package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukInformationRouter struct {
}

// InitLukInformationRouter 初始化 LukInformation 路由信息
func (s *LukInformationRouter) InitLukInformationRouter(Router *gin.RouterGroup) {
	lukInformationRouter := Router.Group("lukInformation").Use(middleware.OperationRecord())
	lukInformationRouterWithoutRecord := Router.Group("lukInformation")
	var lukInformationApi = v1.ApiGroupApp.LukApiGroup.LukInformationApi
	{
		lukInformationRouter.POST("createLukInformation", lukInformationApi.CreateLukInformation)   // 新建LukInformation
		lukInformationRouter.DELETE("deleteLukInformation", lukInformationApi.DeleteLukInformation) // 删除LukInformation
		lukInformationRouter.PUT("updateLukInformation", lukInformationApi.UpdateLukInformation)    // 更新LukInformation
	}
	{
		lukInformationRouterWithoutRecord.GET("findLukInformation", lukInformationApi.FindLukInformation)       // 根据ID获取LukInformation
		lukInformationRouterWithoutRecord.GET("getLukInformationList", lukInformationApi.GetLukInformationList) // 获取LukInformation列表
	}

	Router.Group("apih5/lukInformation").GET("list", lukInformationApi.GetLukInformationList)
}
