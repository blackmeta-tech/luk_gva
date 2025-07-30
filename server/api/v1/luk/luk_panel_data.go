package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

var lukPanelDataService = service.ServiceGroupApp.LukServiceGroup.LukPanelDataService

type LukPanelDataApi struct {
}

func (lukPanelDataApi *LukPanelDataApi) QueryData(c *gin.Context) {
	data := lukPanelDataService.QueryData()
	response.OkWithData(data, c)
}
