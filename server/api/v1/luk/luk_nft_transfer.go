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

type LukNftTransferApi struct {
}

var lukNftTransferService = service.ServiceGroupApp.LukServiceGroup.LukNftTransferService

// CreateLukNftTransfer 创建LukNftTransfer
// @Tags LukNftTransfer
// @Summary 创建LukNftTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukNftTransfer true "创建LukNftTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftTransfer/createLukNftTransfer [post]
func (lukNftTransferApi *LukNftTransferApi) CreateLukNftTransfer(c *gin.Context) {
	var lukNftTransfer luk.LukNftTransfer
	err := c.ShouldBindJSON(&lukNftTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukNftTransferService.CreateLukNftTransfer(lukNftTransfer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败，"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindLukNftTransfer 用id查询LukNftTransfer
// @Tags LukNftTransfer
// @Summary 用id查询LukNftTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukNftTransfer true "用id查询LukNftTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukNftTransfer/findLukNftTransfer [get]
func (lukNftTransferApi *LukNftTransferApi) FindLukNftTransfer(c *gin.Context) {
	var lukNftTransfer luk.LukNftTransfer
	err := c.ShouldBindQuery(&lukNftTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukNftTransfer, err := lukNftTransferService.GetLukNftTransfer(lukNftTransfer.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukNftTransfer": relukNftTransfer}, c)
	}
}

// GetLukNftTransferList 分页获取LukNftTransfer列表
// @Tags LukNftTransfer
// @Summary 分页获取LukNftTransfer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukNftTransferSearch true "分页获取LukNftTransfer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNftTransfer/getLukNftTransferList [get]
func (lukNftTransferApi *LukNftTransferApi) GetLukNftTransferList(c *gin.Context) {
	var pageInfo lukReq.LukNftTransferSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukNftTransferService.GetLukNftTransferInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
