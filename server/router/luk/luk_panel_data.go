package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukPanelDataRouter struct {
}

// InitLukDappRouter 初始化LukPanelData 路由信息
func (s *LukDappRouter) InitLukPanelDataRouter(Router *gin.RouterGroup) {
	var routersAdmin = Router.Group("lukPanelData")
	var lukDappApi = v1.ApiGroupApp.LukApiGroup.LukPanelDataApi
	routersAdmin.GET("queryData", lukDappApi.QueryData) //后端访问
}
