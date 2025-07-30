package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukConstantRouter struct {
}

// InitLukConstantRouter 初始化 LukConstant 路由信息
func (s *LukConstantRouter) InitLukConstantRouter(Router *gin.RouterGroup) {
	lukConstantRouterWithoutRecord := Router.Group("lukConstant").Use(middleware.OperationRecord())
	var lukConstantApi = v1.ApiGroupApp.LukApiGroup.LukConstantApi
	lukConstantRouterWithoutRecord.POST("getLukConstant", lukConstantApi.GetLukConstant)
}
