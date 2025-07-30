package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	lukSer "github.com/flipped-aurora/gin-vue-admin/server/service/luk"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukWithdrawHistoryApi struct {
}

var WithdrawHistoryService = service.ServiceGroupApp.LukServiceGroup.LukWithdrawHistoryService

// UpdateLukWithdrawHistory 更新LukWithdrawHistory
// @Tags LukWithdrawHistory
// @Summary 更新LukWithdrawHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawHistory true "更新LukWithdrawHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /WithdrawHistory/updateLukWithdrawHistory [put]
func (WithdrawHistoryApi *LukWithdrawHistoryApi) UpdateLukWithdrawHistory(c *gin.Context) {
	var WithdrawHistory luk.LukWithdrawHistory
	_ = c.ShouldBindJSON(&WithdrawHistory)
	if err := WithdrawHistoryService.Update(WithdrawHistory); err != nil {
		response.FailWithMessage("更新失败，"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukWithdrawHistory 用id查询LukWithdrawHistory
// @Tags LukWithdrawHistory
// @Summary 用id查询LukWithdrawHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukWithdrawHistory true "用id查询LukWithdrawHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /WithdrawHistory/findLukWithdrawHistory [get]
func (WithdrawHistoryApi *LukWithdrawHistoryApi) FindLukWithdrawHistory(c *gin.Context) {
	var WithdrawHistory luk.LukWithdrawHistory
	_ = c.ShouldBindQuery(&WithdrawHistory)
	if reWithdrawHistory, err := WithdrawHistoryService.GetLukWithdrawHistory(WithdrawHistory.ID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reWithdrawHistory": reWithdrawHistory}, c)
	}
}

// GetLukWithdrawHistoryList 分页获取LukWithdrawHistory列表
// @Tags LukWithdrawHistory
// @Summary 分页获取LukWithdrawHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukWithdrawHistorySearch true "分页获取LukWithdrawHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /WithdrawHistory/getLukWithdrawHistoryList [get]
func (WithdrawHistoryApi *LukWithdrawHistoryApi) GetLukWithdrawHistoryList(c *gin.Context) {
	var pageInfo lukReq.LukWithdrawHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	query := lukSer.QueryLukWithdrawHistory{
		Info: pageInfo,
	}
	if list, total, err := query.GetLukWithdrawHistoryInfoList(); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		//获取汇总数据
		countAddress, _ := query.GetCountAddress()
		lukAmount, usdtAmount := query.GetCountAmount()
		res := map[string]interface{}{
			"list":         list,
			"total":        total,
			"page":         pageInfo.Page,
			"pageSize":     pageInfo.PageSize,
			"countAddress": countAddress,
			"lukAmount":    lukAmount,
			"usdtAmount":   usdtAmount,
		}
		response.OkWithDetailed(res, "获取成功", c)
	}
}
