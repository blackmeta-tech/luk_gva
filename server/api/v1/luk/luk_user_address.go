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

type LukUserAddressApi struct {
}

var lukUserAddressService = service.ServiceGroupApp.LukServiceGroup.LukUserAddressService

func (lukUserAddressApi *LukUserAddressApi) UpdateLukUserAddress(c *gin.Context) {
	var user lukReq.LukUserAddressUpdate
	_ = c.ShouldBindJSON(&user)
	if err := lukUserAddressService.Update(user); err != nil {
		global.GVA_LOG.Error("创建Or更新失败!", zap.Error(err))
		response.FailWithMessage("创建Or更新失败,"+err.Error(), c)
	} else {
		response.OkWithMessage("创建Or更新成功", c)
	}
}

func (wLukUserAddressApi *LukUserAddressApi) GetLukUserAddressList(c *gin.Context) {
	var pageInfo lukReq.LukUserAddressSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lukUserAddressService.GetList(pageInfo); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

func (wLukUserAddressApi *LukUserAddressApi) FindLukUserAddress(c *gin.Context) {
	var userAddress luk.LukUserAddress
	_ = c.ShouldBindQuery(&userAddress)
	if userAddress.Address == "" {
		response.FailWithMessage("地址不能为空", c)
		return
	}
	if relukUserWallet, err := lukUserAddressService.GetOne(userAddress.Address); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(relukUserWallet, c)
	}
}
