package luk

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukRebateService struct {
}

// CreateLukRebate 创建LukRebate记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRebateService *LukRebateService) CreateLukRebate(lukRebate luk.LukRebate) (err error) {
	lukRebate.Address = strings.ToLower(lukRebate.Address)
	metaverse := luk.LukRebate{}
	global.GVA_DB.Where("address = ?", lukRebate.Address).First(&metaverse)
	if metaverse.ID > 0 {
		err = errors.New("该用户地址已录入")
		return
	}
	lukRebate.Remaining = lukRebate.Balance
	err = global.GVA_DB.Create(&lukRebate).Error
	return err
}

// DeleteLukRebate 删除LukRebate记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRebateService *LukRebateService) DeleteLukRebate(lukRebate luk.LukRebate) (err error) {
	err = global.GVA_DB.Delete(&lukRebate).Error
	return err
}

// UpdateLukRebate 更新LukRebate记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRebateService *LukRebateService) UpdateLukRebate(lukRebate luk.LukRebate) (err error) {
	err = global.GVA_DB.Save(&lukRebate).Error
	return err
}

// GetLukRebate 根据id获取LukRebate记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRebateService *LukRebateService) GetLukRebate(id uint) (lukRebate luk.LukRebate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRebate).Error
	return
}

func (lukRebateService *LukRebateService) QueryLukRebateByAddress(address string) (rebate interface{}, err error) {
	rebateData := luk.LukRebate{}
	err = global.GVA_DB.Debug().Model(&luk.LukRebate{}).Where("address = ?", strings.ToLower(address)).First(&rebateData).Error
	b, _ := json.Marshal(&rebateData)
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	data["rebatePerformance"] = global.GVA_LUK_CONFIG.RebatePerformance
	rebate = data
	return rebate, err
}

// GetLukRebateInfoList 分页获取LukRebate记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRebateService *LukRebateService) GetLukRebateInfoList(info lukReq.LukRebateSearch) (list []luk.LukRebate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukRebate{})
	var lukRebates []luk.LukRebate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db.Where("address = ?", strings.ToLower(info.Address))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukRebates).Error
	return lukRebates, total, err
}

func (lukRebateService *LukRebateService) QueryAll() (address []string) {
	address = make([]string, 0)
	data := []luk.LukRebate{}
	_ = global.GVA_DB.Debug().Model(&luk.LukRebate{}).Find(&data).Error
	for _, item := range data {
		address = append(address, item.Address)
	}
	return
}
