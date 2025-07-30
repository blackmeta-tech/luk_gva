package luk

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	tokenService "github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strings"
	"time"
)

type LukUserAddressService struct {
}

// 更新or创建数据
func (LukUserAddressService) Update(data lukReq.LukUserAddressUpdate) (err error) {
	if data.Address == "" {
		err = errors.New("地址不得为空")
		return
	}
	address := strings.ToLower(data.Address)
	updateInfo := map[string]interface{}{
		"address":    address,
		"updated_at": time.Now(),
		"remark":     data.Remark,
	}
	tokenService := &tokenService.LukTokenService{}
	//请求调用链上接口
	err = tokenService.SetWhiteList(common.HexToAddress(address), data.Type)

	if data.Type == 1 {
		status, _ := tokenService.IsWhiteList(address)
		if !status {
			fmt.Println("=====", status)
			err = errors.New("添加失败，合约返回错误")
			return
		}
	}
	if err != nil {
		return
	}
	updateInfo["is_charge"] = data.Type
	err = global.GVA_DB.Debug().Where(luk.LukUserAddress{Address: address}).Assign(updateInfo).FirstOrCreate(&luk.LukUserAddress{}).Error
	return
}

func (LukUserAddressService) GetList(info lukReq.LukUserAddressSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Debug().Model(&luk.LukUserAddress{})
	var lukUserAddress []luk.LukUserAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("is_charge = 1")
	if info.Address != "" {
		db = db.Where("address = ? ", strings.ToLower(info.Address))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukUserAddress).Error
	return lukUserAddress, total, err
}

// Author [piexlmax](https://github.com/piexlmax)
func (LukUserAddressService) GetOne(address string) (userAddres luk.LukUserAddress, err error) {
	err = global.GVA_DB.Model(&luk.LukUserAddress{}).Where("address = ?", address).First(&userAddres).Error
	return
}

//获取白名单地址
func (LukUserAddressService) QueryWhitelist() (list []string) {
	data := []luk.LukUserAddress{}
	err := global.GVA_DB.Model(&luk.LukUserAddress{}).Where("is_charge = 1").Find(&data).Error
	if err != nil {
		return
	}
	for _, item := range data {
		list = append(list, item.Address)
	}
	return
}

func (LukUserAddressService) QueryLpPro(tx *gorm.DB) (data map[string]decimal.Decimal) {
	data = make(map[string]decimal.Decimal, 0)
	userAddres := []luk.LukUserAddress{}
	err := tx.Model(&luk.LukUserAddress{}).Where("lp > 0 and is_charge = 0").Find(&userAddres).Error
	if err != nil {
		return
	}
	var totol decimal.Decimal
	for _, item := range userAddres {
		totol = totol.Add(item.Lp)
	}
	if totol.IsZero() {
		return
	}
	for _, item := range userAddres {
		data[item.Address] = item.Lp.Div(totol)
	}
	return
}
