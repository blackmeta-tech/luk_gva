package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukWithdrawBlacklistService struct {
}

// CreateLukWithdrawBlacklist 创建LukWithdrawBlacklist记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) CreateLukWithdrawBlacklist(lukWithdrawBlacklist luk.LukWithdrawBlacklist) (err error) {
	lukWithdrawBlacklist.Address = strings.ToLower(lukWithdrawBlacklist.Address)
	blacklist, _ := lukWithdrawBlacklistService.QueryLukWithdrawBlacklist(lukWithdrawBlacklist.Address)
	if blacklist.ID > 0 {
		err = errors.New("该用户地址已经存在黑名单")
		return
	}
	err = global.GVA_DB.Create(&lukWithdrawBlacklist).Error
	return err
}

// DeleteLukWithdrawBlacklist 删除LukWithdrawBlacklist记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) DeleteLukWithdrawBlacklist(lukWithdrawBlacklist luk.LukWithdrawBlacklist) (err error) {
	err = global.GVA_DB.Delete(&lukWithdrawBlacklist).Error
	return err
}

// UpdateLukWithdrawBlacklist 更新LukWithdrawBlacklist记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) UpdateLukWithdrawBlacklist(lukWithdrawBlacklist luk.LukWithdrawBlacklist) (err error) {
	lukWithdrawBlacklist.Address = strings.ToLower(lukWithdrawBlacklist.Address)
	err = global.GVA_DB.Save(&lukWithdrawBlacklist).Error
	return err
}

// GetLukWithdrawBlacklist 根据id获取LukWithdrawBlacklist记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) GetLukWithdrawBlacklist(id uint) (lukWithdrawBlacklist luk.LukWithdrawBlacklist, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukWithdrawBlacklist).Error
	return
}

func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) QueryLukWithdrawBlacklist(address string) (lukWithdrawBlacklist luk.LukWithdrawBlacklist, err error) {
	err = global.GVA_DB.Where("address = ?", strings.ToLower(address)).First(&lukWithdrawBlacklist).Error
	return
}

// GetLukWithdrawBlacklistInfoList 分页获取LukWithdrawBlacklist记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawBlacklistService *LukWithdrawBlacklistService) GetLukWithdrawBlacklistInfoList(info lukReq.LukWithdrawBlacklistSearch) (list []luk.LukWithdrawBlacklist, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukWithdrawBlacklist{})
	var lukWithdrawBlacklists []luk.LukWithdrawBlacklist
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

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukWithdrawBlacklists).Error
	return lukWithdrawBlacklists, total, err
}
