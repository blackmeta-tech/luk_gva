package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukRebateRouter struct {
}

// InitLukRebateRouter 初始化 LukRebate 路由信息
func (s *LukRebateRouter) InitLukRebateRouter(Router *gin.RouterGroup) {
	lukRebateRouter := Router.Group("lukRebate").Use(middleware.OperationRecord())
	lukRebateRouterWithoutRecord := Router.Group("lukRebate")
	var lukRebateApi = v1.ApiGroupApp.LukApiGroup.LukRebateApi
	{
		lukRebateRouter.POST("createLukRebate", lukRebateApi.CreateLukRebate)   // 新建LukRebate
		lukRebateRouter.DELETE("deleteLukRebate", lukRebateApi.DeleteLukRebate) // 删除LukRebate
		lukRebateRouter.PUT("updateLukRebate", lukRebateApi.UpdateLukRebate)    // 更新LukRebate
	}
	{
		lukRebateRouterWithoutRecord.GET("findLukRebate", lukRebateApi.FindLukRebate)       // 根据ID获取LukRebate
		lukRebateRouterWithoutRecord.GET("getLukRebateList", lukRebateApi.GetLukRebateList) // 获取LukRebate列表
	}
	Router.Group("apih5/lukRebate").GET("query", lukRebateApi.QueryLukRebate) // 获取LukRebate列表
}
