package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LukComboApi struct {
}

var lukComboService = service.ServiceGroupApp.LukServiceGroup.LukComboService

// CreateLukCombo 创建LukCombo
// @Tags LukCombo
// @Summary 创建LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukCombo true "创建LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukCombo/createLukCombo [post]
func (lukComboApi *LukComboApi) CreateLukCombo(c *gin.Context) {
	var lukCombo luk.LukCombo
	err := c.ShouldBindJSON(&lukCombo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Days":     {utils.NotEmpty()},
		"Limit":    {utils.NotEmpty()},
		"Pic":      {utils.NotEmpty()},
		"PriceMax": {utils.NotEmpty()},
		"PriceMin": {utils.NotEmpty()},
		"Rate":     {utils.NotEmpty()},
	}
	if err := utils.Verify(lukCombo, verify); err != nil {
		response.FailWithMessage("必填参数不足", c)
		return
	}
	if err := lukComboService.CreateLukCombo(lukCombo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateLukCombo 更新LukCombo
// @Tags LukCombo
// @Summary 更新LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukCombo true "更新LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukCombo/updateLukCombo [put]
func (lukComboApi *LukComboApi) UpdateLukCombo(c *gin.Context) {
	var lukCombo luk.LukCombo
	err := c.ShouldBindJSON(&lukCombo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Days": {utils.NotEmpty()},
	}
	if err := utils.Verify(lukCombo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukComboService.UpdateLukCombo(lukCombo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLukCombo 用id查询LukCombo
// @Tags LukCombo
// @Summary 用id查询LukCombo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukCombo true "用id查询LukCombo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukCombo/findLukCombo [get]
func (lukComboApi *LukComboApi) FindLukCombo(c *gin.Context) {
	var lukCombo luk.LukCombo
	err := c.ShouldBindQuery(&lukCombo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukCombo, err := lukComboService.GetLukCombo(lukCombo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukCombo": relukCombo}, c)
	}
}

// GetLukComboList 分页获取LukCombo列表
// @Tags LukCombo
// @Summary 分页获取LukCombo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukComboSearch true "分页获取LukCombo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukCombo/getLukComboList [get]
func (lukComboApi *LukComboApi) GetLukComboList(c *gin.Context) {
	var pageInfo lukReq.LukComboSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukComboService.GetLukComboInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

//前端获取套餐列表接口
func (lukComboApi *LukComboApi) QueryLukComboList(c *gin.Context) {
	var q lukReq.LukComboSearchByOne
	err := c.ShouldBindQuery(&q)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := lukComboService.GetLukComboInfoListByAddress(q.Address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

//查看用户下的全部套餐业绩
func (lukComboApi *LukComboApi) QueryLukComboPerformance(c *gin.Context) {
	var q lukReq.LukComboSearchByOne
	err := c.ShouldBindQuery(&q)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, performances, performancesU, err := lukComboService.QueryPerformance(q.Address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		response.OkWithDetailed(map[string]interface{}{
			"list":             list,
			"performances":     performances,
			"performancesUsdt": performancesU,
		}, "获取成功", c)
	}
}

//销毁列表
func (D *LukComboApi) GetLukDestroyInfoList(c *gin.Context) {
	var pageInfo lukReq.LukDestroySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := service.ServiceGroupApp.LukServiceGroup.LukDestroyService.GetLukDestroyInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		summary := service.ServiceGroupApp.LukServiceGroup.LukDestroyService.QuerySummary(pageInfo.Address)
		response.OkWithDetailed(map[string]interface{}{
			"summary":  summary,
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//回流记录
func (D *LukComboApi) GetLukReflowInfoList(c *gin.Context) {
	var pageInfo lukReq.LukReflowSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := service.ServiceGroupApp.LukServiceGroup.LukReflowService.GetLukReflowInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
