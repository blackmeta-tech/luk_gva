package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//主要是提供给DAPP前端接口
type LukDappApi struct {
}

var lukDappService = service.ServiceGroupApp.LukServiceGroup.LukDappService

// 获取底池数据
func (lukDappApi *LukDappApi) GetPoolBasic(c *gin.Context) {
	data, err := service.ServiceGroupApp.LukServiceGroup.LukPoolBasicService.GetDataDapp()
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(data, c)
	}
}

func (lukDappApi *LukDappApi) TransactionRecord(c *gin.Context) {
	data := lukDappService.TransactionRecord()
	response.OkWithData(data, c)
}

//移除跟添加流动操作接口
func (lukDappApi *LukDappApi) BlockScanRecordList(c *gin.Context) {
	var pageInfo lukReq.BlockScanRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, summary, total, err := lukDappService.BlockScanRecordList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]interface{}{
			"summary":  summary,
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//获取上线状态
func (lukDappApi *LukDappApi) QueryOnline(c *gin.Context) {
	response.OkWithDetailed(global.GVA_LUK_CONFIG.Online, global.GVA_LUK_CONFIG.OnlineText, c)
}

//获取遗产数据
func (lukDappApi *LukDappApi) QueryHeritage(c *gin.Context) {
	var user lukReq.QuerryUser
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	datas := lukDappService.QueryHeritage(user.Address)
	response.OkWithDetailed(datas, "获取成功", c)
}
