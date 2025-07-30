package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukConfigRouter struct {
}

// InitLukConfigRouter 初始化 LukConfig 路由信息
func (s *LukConfigRouter) InitLukConfigRouter(Router *gin.RouterGroup) {
	var routersAdmin = Router.Group("lukConfig")
	var lukConfigApi = v1.ApiGroupApp.LukApiGroup.LukConfigApi
	routersAdmin.GET("getConfig", lukConfigApi.GetConfig)                                                 //后端访问
	Router.Group("lukConfig").Use(middleware.OperationRecord()).POST("setConfig", lukConfigApi.SetConfig) //后端访问
}
