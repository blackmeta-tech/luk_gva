package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukComboRouter struct {
}

// InitLukComboRouter 初始化 LukCombo 路由信息
func (s *LukComboRouter) InitLukComboRouter(Router *gin.RouterGroup) {
	lukComboRouter := Router.Group("lukCombo").Use(middleware.OperationRecord())
	lukComboRouterWithoutRecord := Router.Group("lukCombo")
	var lukComboApi = v1.ApiGroupApp.LukApiGroup.LukComboApi
	{
		lukComboRouter.POST("createLukCombo", lukComboApi.CreateLukCombo) // 新建LukCombo
		lukComboRouter.PUT("updateLukCombo", lukComboApi.UpdateLukCombo)  // 更新LukCombo
	}
	{
		lukComboRouterWithoutRecord.GET("findLukCombo", lukComboApi.FindLukCombo)                   // 根据ID获取LukCombo
		lukComboRouterWithoutRecord.GET("getLukComboList", lukComboApi.GetLukComboList)             // 获取LukCombo列表
		lukComboRouterWithoutRecord.GET("getLukDestroyInfoList", lukComboApi.GetLukDestroyInfoList) // 销毁套餐
		lukComboRouterWithoutRecord.GET("getLukReflowInfoList", lukComboApi.GetLukReflowInfoList)   // 销毁套餐
		lukComboRouterWithoutRecord.GET("list", lukComboApi.QueryLukComboList)                      // 获取LukCombo列表
	}
	Router.Group("apih5/lukCombo").GET("list", lukComboApi.QueryLukComboList)               // 获取LukCombo列表
	Router.Group("apih5/lukCombo").GET("performance", lukComboApi.QueryLukComboPerformance) // 获取LukCombo列表
	Router.Group("apih5/lukCombo").GET("destroyList", lukComboApi.GetLukDestroyInfoList)    // 前端获取销毁记录
}
