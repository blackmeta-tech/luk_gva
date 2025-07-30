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

type LukUserApi struct {
}

var lukUserService = service.ServiceGroupApp.LukServiceGroup.LukUserService

// CreateLukUser 创建LukUser
// @Tags LukUser
// @Summary 创建LukUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luk.LukUser true "创建LukUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukUser/createLukUser [post]
func (lukUserApi *LukUserApi) CreateLukUser(c *gin.Context) {
	var lukUser luk.LukUser
	err := c.ShouldBindJSON(&lukUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, errNUm, err := lukUserService.CreateLukUser(lukUser)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(errNUm, "注册失败，"+err.Error(), c)
	} else {
		response.OkWithDetailed(user, "注册成功", c)
	}
}
func (lukUserApi *LukUserApi) UpdateLukUser(c *gin.Context) {
	var lukUser luk.LukUser
	_ = c.ShouldBindJSON(&lukUser)
	err := lukUserService.UpdateLukUser(lukUser)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

//编辑白名单
func (lukUserApi *LukUserApi) UpdateLinkage(c *gin.Context) {
	var lukUser lukReq.QuerryUser
	err := c.ShouldBindJSON(&lukUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if lukUserService.UpdateLinkage(lukUser.Address) != nil {
		response.FailWithMessage("修改失败，"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// FindLukUser 用id查询LukUser
// @Tags LukUser
// @Summary 用id查询LukUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query luk.LukUser true "用id查询LukUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lukUser/findLukUser [get]
func (lukUserApi *LukUserApi) FindLukUser(c *gin.Context) {
	var lukUser luk.LukUser
	err := c.ShouldBindQuery(&lukUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relukUser, err := lukUserService.GetLukUser(lukUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relukUser": relukUser}, c)
	}
}

// GetLukUserList 分页获取LukUser列表
// @Tags LukUser
// @Summary 分页获取LukUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lukReq.LukUserSearch true "分页获取LukUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lukUser/getLukUserList [get]
func (lukUserApi *LukUserApi) GetLukUserList(c *gin.Context) {
	var pageInfo lukReq.LukUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lukUserService.GetLukUserInfoList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

func (lukUserApi *LukUserApi) QueryUser(c *gin.Context) {
	var u lukReq.QuerryUser
	err := c.ShouldBindQuery(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := lukUserService.QueryUserByAddress(u.Address)
	if err != nil || user.ID == 0 {
		response.FailWithMessage("获取失败,请绑定填写邀请码", c)
	} else {
		response.OkWithDetailed(user, "获取成功", c)
	}
}

//获取直推关系
func (lukUserApi *LukUserApi) QuerySubclass(c *gin.Context) {
	var u lukReq.QuerryUser
	err := c.ShouldBindQuery(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := lukUserService.QuerySubclass(u.Address); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
