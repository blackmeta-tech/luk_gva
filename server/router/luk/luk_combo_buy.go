package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LukComboBuyRouter struct {
}

// InitLukComboBuyRouter 初始化 LukComboBuy 路由信息
func (s *LukComboBuyRouter) InitLukComboBuyRouter(Router *gin.RouterGroup) {
	lukComboBuyRouterWithoutRecord := Router.Group("lukComboBuy")
	h5 := Router.Group("apih5/lukComboBuy")
	var lukComboBuyApi = v1.ApiGroupApp.LukApiGroup.LukComboBuyApi
	{
		lukComboBuyRouterWithoutRecord.GET("findLukComboBuy", lukComboBuyApi.FindLukComboBuy)       // 根据ID获取LukComboBuy
		lukComboBuyRouterWithoutRecord.GET("getLukComboBuyList", lukComboBuyApi.GetLukComboBuyList) // 获取LukComboBuy列表
	}
	h5.POST("updateHash", lukComboBuyApi.UpdateLukComboBuyByHash) // 编辑
	h5.POST("update", lukComboBuyApi.UpdateLukComboBuy)           // 编辑
	h5.POST("create", lukComboBuyApi.CreateLukComboBuy)           // 新增
	h5.POST("delete", lukComboBuyApi.DeleteLukComboBuy)           // 删除
	h5.GET("list", lukComboBuyApi.GetList)                        //
	h5.GET("getGeneration", lukComboBuyApi.GetGeneration)         //

}
