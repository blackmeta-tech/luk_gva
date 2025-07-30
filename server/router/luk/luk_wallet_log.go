package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukWalletLogRouter struct {
}

// InitLukWalletLogRouter 初始化 LukWalletLog 路由信息
func (s *LukWalletLogRouter) InitLukWalletLogRouter(Router *gin.RouterGroup) {
	lukWalletLogRouterWithoutRecord := Router.Group("lukWalletLog")
	var lukWalletLogApi = v1.ApiGroupApp.LukApiGroup.LukWalletLogApi
	{
		lukWalletLogRouterWithoutRecord.GET("getWallet", lukWalletLogApi.GetWallet)
		lukWalletLogRouterWithoutRecord.GET("getWalletConfigure", lukWalletLogApi.GetWalletConfigure)
	}
}
