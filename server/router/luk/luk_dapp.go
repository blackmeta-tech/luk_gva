package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukDappRouter struct {
}

// InitLukDappRouter 初始化 LukDapp 路由信息
func (s *LukDappRouter) InitLukDappRouter(Router *gin.RouterGroup) {
	var routersAdmin = Router.Group("lukDapp")
	var routers = Router.Group("apih5/lukDapp")
	var lukDappApi = v1.ApiGroupApp.LukApiGroup.LukDappApi
	routers.GET("getPoolBasic", lukDappApi.GetPoolBasic)               //前端访问
	routers.GET("transactionRecord", lukDappApi.TransactionRecord)     //前端访问
	routers.GET("blockScanRecordList", lukDappApi.BlockScanRecordList) //前端访问
	routers.GET("queryOnline", lukDappApi.QueryOnline)                 //前端访问
	routers.GET("queryHeritage", lukDappApi.QueryHeritage)             //前端访问

	routersAdmin.GET("getPoolBasic", lukDappApi.GetPoolBasic) //后端访问
}
