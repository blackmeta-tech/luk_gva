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
	"strconv"
)

type LukNftApi struct {
}

var lukNftService = service.ServiceGroupApp.LukServiceGroup.LukNftService

// CreateLukNft 创建LukNft
// @Tags LukNft
// @Summary 创建LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukNft true "创建LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNft/createLukNft [post]
func (lukNftApi *LukNftApi) CreateLukNft(c *gin.Context) {
	var lukNft luk.LukNft
	err := c.ShouldBindJSON(&lukNft)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukNftService.CreateLukNft(lukNft); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateLukNft 更新LukNft
// @Tags LukNft
// @Summary 更新LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukNft true "更新LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lukNft/updateLukNft [put]
func (lukNftApi *LukNftApi) UpdateLukNft(c *gin.Context) {
	var lukNft luk.LukNft
	err := c.ShouldBindJSON(&lukNft)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukNftService.UpdateLukNft(lukNft); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (lukNftApi *LukNftApi) UpdateLukNftBlack(c *gin.Context) {
	var lukNft luk.LukNft
	err := c.ShouldBindJSON(&lukNft)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lukNftService.UpdateLukNftBlack(lukNft.ID); err != nil {
		response.FailWithMessage("添加失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// FindLukNft 用id查询LukNft
// @Tags LukNft
// @Summary 用id查询LukNft
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukNft true "用id查询LukNft"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukNft/findLukNft [get]
func (lukNftApi *LukNftApi) FindLukNft(c *gin.Context) {
	var lukNft luk.LukNft
	err := c.ShouldBindQuery(&lukNft)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukNft, err := lukNftService.GetLukNft(lukNft.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukNft": relukNft}, c)
	}
}

// GetLukNftList 分页获取LukNft列表
// @Tags LukNft
// @Summary 分页获取LukNft列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukNftSearch true "分页获取LukNft列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukNft/getLukNftList [get]
func (lukNftApi *LukNftApi) GetLukNftList(c *gin.Context) {
	var pageInfo lukReq.LukNftSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukNftService.GetLukNftInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

//NFT购买
func (lukNftApi *LukNftApi) BuyLukNft(c *gin.Context) {
	var lukNft luk.LukNft
	err := c.ShouldBindQuery(&lukNft)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pass := lukNft.Address + strconv.Itoa(int(lukNft.ID))
	pass64, err := lukNftService.BuyLukNftAes(pass)
	if err != nil {
		response.FailWithMessage("获取失败,"+err.Error(), c)
	} else {
		response.OkWithDetailed(pass64, "获取成功", c)
	}
}

//NFT汇总
func (lukNftApi *LukNftApi) Summary(c *gin.Context) {
	circulation, hold, history := lukNftService.Summary()
	data := map[string]interface{}{
		"nftIssued":      global.GVA_LUK_CONFIG.NftIssued,
		"circulationNum": circulation,
		"holdNum":        hold,
		"historyNum":     history,
	}
	response.OkWithDetailed(data, "获取成功", c)
}
