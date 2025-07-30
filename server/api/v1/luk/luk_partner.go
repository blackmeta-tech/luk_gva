package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukPartnerApi struct {
}

var LukPartnerService = service.ServiceGroupApp.LukServiceGroup.LukPartnerService

// UpdateLukPartner 更新LukPartner
// @Tags LukPartner
// @Summary 更新LukPartner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukPartner true "更新LukPartner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /LukPartner/updateLukPartner [put]
func (LukPartnerApi *LukPartnerApi) UpdateLukPartner(c *gin.Context) {
	var LukPartner luk.LukPartner
	_ = c.ShouldBindJSON(&LukPartner)
	//查出id对应的key
	if err := LukPartnerService.UpdateLukPartner(LukPartner); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukPartner 用id查询LukPartner
// @Tags LukPartner
// @Summary 用id查询LukPartner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukPartner true "用id查询LukPartner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /LukPartner/findLukPartner [get]
func (LukPartnerApi *LukPartnerApi) FindLukPartner(c *gin.Context) {
	var LukPartner luk.LukPartner
	_ = c.ShouldBindQuery(&LukPartner)
	if reLukPartner, err := LukPartnerService.GetLukPartner(LukPartner.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reLukPartner": reLukPartner}, c)
	}
}

// GetLukPartnerList 分页获取LukPartner列表
// @Tags LukPartner
// @Summary 分页获取LukPartner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukPartnerSearch true "分页获取LukPartner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /LukPartner/getLukPartnerList [get]
func (LukPartnerApi *LukPartnerApi) GetLukPartnerList(c *gin.Context) {
	var pageInfo lukReq.LukPartnerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := LukPartnerService.GetLukPartnerInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//前端获取联盟社区配置
func (LukPartnerApi *LukPartnerApi) QueryLukPartnerList(c *gin.Context) {
	var LukPartner luk.LukPartner
	_ = c.ShouldBindQuery(&LukPartner)
	if list, err := LukPartnerService.QueryList(LukPartner.Type, "asc"); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(list, c)
	}
}
