package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukNftTransferRouter struct {
}

// InitLukNftTransferRouter 初始化 LukNftTransfer 路由信息
func (s *LukNftTransferRouter) InitLukNftTransferRouter(Router *gin.RouterGroup) {
	lukNftTransferRouterWithoutRecord := Router.Group("lukNftTransfer")
	var lukNftTransferApi = v1.ApiGroupApp.LukApiGroup.LukNftTransferApi
	{
		Router.Group("apih5/lukNftTransfer").POST("create", lukNftTransferApi.CreateLukNftTransfer) // 新建LukNftTransfer
	}
	{
		lukNftTransferRouterWithoutRecord.GET("findLukNftTransfer", lukNftTransferApi.FindLukNftTransfer)       // 根据ID获取LukNftTransfer
		lukNftTransferRouterWithoutRecord.GET("getLukNftTransferList", lukNftTransferApi.GetLukNftTransferList) // 获取LukNftTransfer列表
	}
	Router.Group("apih5/lukNftTransfer").GET("list", lukNftTransferApi.GetLukNftTransferList) // 获取LukNftTransfer列表
}
