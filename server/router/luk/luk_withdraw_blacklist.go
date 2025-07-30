package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukWithdrawBlacklistRouter struct {
}

// InitLukWithdrawBlacklistRouter 初始化 LukWithdrawBlacklist 路由信息
func (s *LukWithdrawBlacklistRouter) InitLukWithdrawBlacklistRouter(Router *gin.RouterGroup) {
	lukWithdrawBlacklistRouter := Router.Group("lukWithdrawBlacklist").Use(middleware.OperationRecord())
	lukWithdrawBlacklistRouterWithoutRecord := Router.Group("lukWithdrawBlacklist")
	var lukWithdrawBlacklistApi = v1.ApiGroupApp.LukApiGroup.LukWithdrawBlacklistApi
	{
		lukWithdrawBlacklistRouter.POST("createLukWithdrawBlacklist", lukWithdrawBlacklistApi.CreateLukWithdrawBlacklist)   // 新建LukWithdrawBlacklist
		lukWithdrawBlacklistRouter.DELETE("deleteLukWithdrawBlacklist", lukWithdrawBlacklistApi.DeleteLukWithdrawBlacklist) // 删除LukWithdrawBlacklist
		lukWithdrawBlacklistRouter.PUT("updateLukWithdrawBlacklist", lukWithdrawBlacklistApi.UpdateLukWithdrawBlacklist)    // 更新LukWithdrawBlacklist
	}
	{
		lukWithdrawBlacklistRouterWithoutRecord.GET("findLukWithdrawBlacklist", lukWithdrawBlacklistApi.FindLukWithdrawBlacklist)       // 根据ID获取LukWithdrawBlacklist
		lukWithdrawBlacklistRouterWithoutRecord.GET("getLukWithdrawBlacklistList", lukWithdrawBlacklistApi.GetLukWithdrawBlacklistList) // 获取LukWithdrawBlacklist列表
	}
}
