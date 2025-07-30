package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukHeritageMetaverseOldService struct {
}

// CreateLukHeritageMetaverseOld 创建LukHeritageMetaverseOld记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseOldService *LukHeritageMetaverseOldService) CreateLukHeritageMetaverseOld(lukHeritageMetaverseOld luk.LukHeritageMetaverseOld) (err error) {
	lukHeritageMetaverseOld.Address = strings.ToLower(lukHeritageMetaverseOld.Address)
	info := luk.LukHeritageMetaverseOld{}
	global.GVA_DB.Where("address = ?", lukHeritageMetaverseOld.Address).First(&info)
	if info.ID > 0 {
		err = errors.New("该用户地址已录入")
		return
	}
	lukHeritageMetaverseOld.Remaining = lukHeritageMetaverseOld.Balance
	err = global.GVA_DB.Create(&lukHeritageMetaverseOld).Error
	return err
}

// DeleteLukHeritageMetaverseOld 删除LukHeritageMetaverseOld记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseOldService *LukHeritageMetaverseOldService) DeleteLukHeritageMetaverseOld(lukHeritageMetaverseOld luk.LukHeritageMetaverseOld) (err error) {
	err = global.GVA_DB.Delete(&lukHeritageMetaverseOld).Error
	return err
}

// UpdateLukHeritageMetaverseOld 更新LukHeritageMetaverseOld记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseOldService *LukHeritageMetaverseOldService) UpdateLukHeritageMetaverseOld(lukHeritageMetaverseOld luk.LukHeritageMetaverseOld) (err error) {
	err = global.GVA_DB.Save(&lukHeritageMetaverseOld).Error
	return err
}

// GetLukHeritageMetaverseOld 根据id获取LukHeritageMetaverseOld记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseOldService *LukHeritageMetaverseOldService) GetLukHeritageMetaverseOld(id uint) (lukHeritageMetaverseOld luk.LukHeritageMetaverseOld, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukHeritageMetaverseOld).Error
	return
}

// GetLukHeritageMetaverseOldInfoList 分页获取LukHeritageMetaverseOld记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseOldService *LukHeritageMetaverseOldService) GetLukHeritageMetaverseOldInfoList(info lukReq.LukHeritageMetaverseOldSearch) (list []luk.LukHeritageMetaverseOld, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukHeritageMetaverseOld{})
	var lukHeritageMetaverseOlds []luk.LukHeritageMetaverseOld
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

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukHeritageMetaverseOlds).Error
	return lukHeritageMetaverseOlds, total, err
}
