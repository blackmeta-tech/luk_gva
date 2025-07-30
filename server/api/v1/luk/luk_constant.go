package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	lukSer "github.com/flipped-aurora/gin-vue-admin/server/service/luk"
	"github.com/gin-gonic/gin"
)

type LukConstantApi struct {
}

var lukConstantService = service.ServiceGroupApp.LukServiceGroup.LukConstantService

//GetLukConstant 获取常量定义
func (lukDappInformationApi *LukConstantApi) GetLukConstant(c *gin.Context) {
	var data lukSer.LukConstantService
	_ = c.ShouldBindJSON(&data)
	lukConstantService.Keys = data.Keys
	list := lukConstantService.GetLukConstant()
	response.OkWithData(list, c)
}
