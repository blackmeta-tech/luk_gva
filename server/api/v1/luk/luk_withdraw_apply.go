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

type LukWithdrawApplyApi struct {
}

var lukWithdrawApplyService = service.ServiceGroupApp.LukServiceGroup.LukWithdrawApplyService

// CreateLukWithdrawApply 创建LukWithdrawApply
// @Tags LukWithdrawApply
// @Summary 创建LukWithdrawApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawApply true "创建LukWithdrawApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawApply/createLukWithdrawApply [post]
func (lukWithdrawApplyApi *LukWithdrawApplyApi) CreateLukWithdrawApply(c *gin.Context) {
	var lukWithdrawApply luk.LukWithdrawApply
	_ = c.ShouldBindJSON(&lukWithdrawApply)
	msg, err := lukWithdrawApplyService.CreateLukWithdrawApply(lukWithdrawApply)
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("已申请，"+msg, c)
	}
}

// UpdateLukWithdrawApply 更新LukWithdrawApply
// @Tags LukWithdrawApply
// @Summary 更新LukWithdrawApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukWithdrawApply true "更新LukWithdrawApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukWithdrawApply/updateLukWithdrawApply [put]
func (lukWithdrawApplyApi *LukWithdrawApplyApi) UpdateLukWithdrawApply(c *gin.Context) {
	var lukWithdrawApply luk.LukWithdrawApply
	_ = c.ShouldBindJSON(&lukWithdrawApply)
	err := lukWithdrawApplyService.UpdateLukWithdrawApply(lukWithdrawApply)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukWithdrawApply 用id查询LukWithdrawApply
// @Tags LukWithdrawApply
// @Summary 用id查询LukWithdrawApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukWithdrawApply true "用id查询LukWithdrawApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukWithdrawApply/findLukWithdrawApply [get]
func (lukWithdrawApplyApi *LukWithdrawApplyApi) FindLukWithdrawApply(c *gin.Context) {
	var lukWithdrawApply luk.LukWithdrawApply
	_ = c.ShouldBindQuery(&lukWithdrawApply)
	if relukWithdrawApply, err := lukWithdrawApplyService.GetLukWithdrawApply(lukWithdrawApply.ID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukWithdrawApply": relukWithdrawApply}, c)
	}
}

// GetLukWithdrawApplyList 分页获取LukWithdrawApply列表
// @Tags LukWithdrawApply
// @Summary 分页获取LukWithdrawApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukWithdrawApplySearch true "分页获取LukWithdrawApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawApply/getLukWithdrawApplyList [get]
func (lukWithdrawApplyApi *LukWithdrawApplyApi) GetLukWithdrawApplyList(c *gin.Context) {
	var pageInfo lukReq.LukWithdrawApplySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukWithdrawApplyService.GetLukWithdrawApplyInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

// GetWithdrawalBalance 获取提现钱包额度
// @Tags LukWithdrawApply
// @Summary 获取提现钱包额度
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukWithdrawApply/getWithdrawalBalance [get]
func (lukWithdrawApplyApi *LukWithdrawApplyApi) GetWithdrawalBalance(c *gin.Context) {
	services := service.ServiceGroupApp.LukServiceGroup.LukWithdrawHistoryService
	balance := services.GetWithdrawalBalance()
	response.OkWithDetailed(balance, "获取成功", c)
}
