package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukWithdrawApplyRouter struct {
}

// InitLukWithdrawApplyRouter 初始化 LukWithdrawApply 路由信息
func (s *LukWithdrawApplyRouter) InitLukWithdrawApplyRouter(Router *gin.RouterGroup) {
	lukWithdrawApplyRouter := Router.Group("lukWithdrawApply").Use(middleware.OperationRecord())
	lukWithdrawApplyRouterWithoutRecord := Router.Group("lukWithdrawApply")
	var lukWithdrawApplyApi = v1.ApiGroupApp.LukApiGroup.LukWithdrawApplyApi
	{
		lukWithdrawApplyRouter.PUT("updateLukWithdrawApply", lukWithdrawApplyApi.UpdateLukWithdrawApply) // 更新LukWithdrawApply
	}
	{
		lukWithdrawApplyRouterWithoutRecord.GET("findLukWithdrawApply", lukWithdrawApplyApi.FindLukWithdrawApply)       // 根据ID获取LukWithdrawApply
		lukWithdrawApplyRouterWithoutRecord.GET("getLukWithdrawApplyList", lukWithdrawApplyApi.GetLukWithdrawApplyList) // 获取LukWithdrawApply列表
		lukWithdrawApplyRouterWithoutRecord.GET("getWithdrawalBalance", lukWithdrawApplyApi.GetWithdrawalBalance)       // 获取提现钱包额度
	}

	Router.Group("apih5/lukWithdrawApply").POST("create", lukWithdrawApplyApi.CreateLukWithdrawApply) // 新建LukWithdrawApply
	Router.Group("apih5/lukWithdrawApply").GET("list", lukWithdrawApplyApi.GetLukWithdrawApplyList)
}
