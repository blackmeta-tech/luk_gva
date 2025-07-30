package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukWithdrawHistoryRouter struct {
}

// InitLukWithdrawHistoryRouter 初始化 LukWithdrawHistory 路由信息
func (s *LukWithdrawHistoryRouter) InitLukWithdrawHistoryRouter(Router *gin.RouterGroup) {
	WithdrawHistoryRouter := Router.Group("lukWithdrawHistory").Use(middleware.OperationRecord())
	WithdrawHistoryRouterWithoutRecord := Router.Group("lukWithdrawHistory")
	var WithdrawHistoryApi = v1.ApiGroupApp.LukApiGroup.LukWithdrawHistoryApi
	{
		WithdrawHistoryRouter.PUT("updateLukWithdrawHistory", WithdrawHistoryApi.UpdateLukWithdrawHistory) // 更新LukWithdrawHistory
	}
	{
		WithdrawHistoryRouterWithoutRecord.GET("findLukWithdrawHistory", WithdrawHistoryApi.FindLukWithdrawHistory)       // 根据ID获取LukWithdrawHistory
		WithdrawHistoryRouterWithoutRecord.GET("getLukWithdrawHistoryList", WithdrawHistoryApi.GetLukWithdrawHistoryList) // 获取LukWithdrawHistory列表
	}

	Router.Group("apih5/lukWithdrawHistory").GET("list", WithdrawHistoryApi.GetLukWithdrawHistoryList)
}
