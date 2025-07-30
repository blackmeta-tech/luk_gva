package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LukConfigApi struct {
}

func (s *LukConfigApi) GetConfig(c *gin.Context) {
	if config, err := (luk.LukConfig{}).GetData(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(config, "获取成功", c)
	}
}

func (s *LukConfigApi) SetConfig(c *gin.Context) {
	var sys luk.LukConfig
	_ = c.ShouldBindJSON(&sys)
	if err := (luk.LukConfig{}).Update(sys); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		data := global.LUK_CONFIG{}
		global.GVA_DB.Debug().Model(luk.LukConfig{}).Where("id = 1").First(&data)
		global.GVA_LUK_CONFIG = data
		response.OkWithMessage("设置成功", c)
	}
}
