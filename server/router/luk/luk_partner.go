package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LukPartnerRouter struct {
}

// InitLukPartnerRouter 初始化 LukPartner 路由信息
func (s *LukPartnerRouter) InitLukPartnerRouter(Router *gin.RouterGroup) {
	lukPartnerRouter := Router.Group("lukPartner").Use(middleware.OperationRecord())
	lukPartnerRouterWithoutRecord := Router.Group("lukPartner")
	var lukPartnerApi = v1.ApiGroupApp.LukApiGroup.LukPartnerApi
	{
		lukPartnerRouter.PUT("updateLukPartner", lukPartnerApi.UpdateLukPartner) // 更新LukPartner
	}
	{
		lukPartnerRouterWithoutRecord.GET("findLukPartner", lukPartnerApi.FindLukPartner)       // 根据ID获取LukPartner
		lukPartnerRouterWithoutRecord.GET("getLukPartnerList", lukPartnerApi.GetLukPartnerList) // 获取LukPartner列表
	}
	Router.Group("apih5/lukPartner").GET("list", lukPartnerApi.QueryLukPartnerList) //前端访问
}
