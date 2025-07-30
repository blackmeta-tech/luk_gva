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

type LukComboBuyApi struct {
}

var lukComboBuyService = service.ServiceGroupApp.LukServiceGroup.LukComboBuyService

// CreateLukComboBuy 创建LukComboBuy
// @Tags LukComboBuy
// @Summary 创建LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukComboBuy true "创建LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukComboBuy/createLukComboBuy [post]
func (lukComboBuyApi *LukComboBuyApi) CreateLukComboBuy(c *gin.Context) {
	var lukComboBuy luk.LukComboBuy
	err := c.ShouldBindJSON(&lukComboBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if errnum, err := lukComboBuyService.CreateLukComboBuy(lukComboBuy); err != nil {
		global.GVA_LOG.Error("购买失败!", zap.Error(err))
		response.FailWithDetailed(errnum, err.Error(), c)
	} else {
		response.OkWithDetailed(errnum, "购买成功", c)
	}
}

// UpdateLukComboBuy 更新LukComboBuy
// @Tags LukComboBuy
// @Summary 更新LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukComboBuy true "更新LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukComboBuy/updateLukComboBuy [post]
func (lukComboBuyApi *LukComboBuyApi) UpdateLukComboBuy(c *gin.Context) {
	var lukComboBuy luk.LukComboBuy
	err := c.ShouldBindJSON(&lukComboBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukComboBuyService.UpdateLukComboBuy(lukComboBuy); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (lukComboBuyApi *LukComboBuyApi) UpdateLukComboBuyByHash(c *gin.Context) {
	var lukComboBuy luk.LukComboBuy
	err := c.ShouldBindJSON(&lukComboBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukComboBuyService.UpdateLukComboBuyByHash(lukComboBuy); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (wLukComboBuyApi *LukComboBuyApi) DeleteLukComboBuy(c *gin.Context) {
	var b luk.LukComboBuy
	_ = c.ShouldBindJSON(&b)
	//查出id对应的key
	if err := lukComboBuyService.DeleteLukComboBuy(b.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// FindLukComboBuy 用id查询LukComboBuy
// @Tags LukComboBuy
// @Summary 用id查询LukComboBuy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukComboBuy true "用id查询LukComboBuy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukComboBuy/findLukComboBuy [get]
func (lukComboBuyApi *LukComboBuyApi) FindLukComboBuy(c *gin.Context) {
	var lukComboBuy luk.LukComboBuy
	err := c.ShouldBindQuery(&lukComboBuy)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukComboBuy, err := lukComboBuyService.GetLukComboBuy(lukComboBuy.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukComboBuy": relukComboBuy}, c)
	}
}

// GetLukComboBuyList 分页获取LukComboBuy列表
// @Tags LukComboBuy
// @Summary 分页获取LukComboBuy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukComboBuySearch true "分页获取LukComboBuy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukComboBuy/getLukComboBuyList [get]
func (lukComboBuyApi *LukComboBuyApi) GetLukComboBuyList(c *gin.Context) {
	var pageInfo lukReq.LukComboBuySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukComboBuyService.GetLukComboBuyInfoList(pageInfo); err != nil {
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

//前端获取购买记录
func (lukComboBuyApi *LukComboBuyApi) GetList(c *gin.Context) {
	var pageInfo lukReq.LukComboBuySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukComboBuyService.QueryList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		summary := lukComboBuyService.QuerySummary(pageInfo)
		response.OkWithDetailed(map[string]interface{}{
			"summary":  summary,
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//前端获取20代下级购买套餐记录
func (lukComboBuyApi *LukComboBuyApi) GetGeneration(c *gin.Context) {
	var pageInfo lukReq.LukComboBuySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukComboBuyService.GetGenerationByAddresss(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]interface{}{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", c)
	}
}
