package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukHeritageMetaverseService struct {
}

// CreateLukHeritageMetaverse 创建LukHeritageMetaverse记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseService *LukHeritageMetaverseService) CreateLukHeritageMetaverse(lukHeritageMetaverse luk.LukHeritageMetaverse) (err error) {
	lukHeritageMetaverse.Address = strings.ToLower(lukHeritageMetaverse.Address)
	metaverse := luk.LukHeritageMetaverse{}
	global.GVA_DB.Where("address = ?", lukHeritageMetaverse.Address).First(&metaverse)
	if metaverse.ID > 0 {
		err = errors.New("该用户地址已录入")
		return
	}
	lukHeritageMetaverse.Remaining = lukHeritageMetaverse.Balance
	err = global.GVA_DB.Create(&lukHeritageMetaverse).Error
	return err
}

// DeleteLukHeritageMetaverse 删除LukHeritageMetaverse记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseService *LukHeritageMetaverseService) DeleteLukHeritageMetaverse(lukHeritageMetaverse luk.LukHeritageMetaverse) (err error) {
	err = global.GVA_DB.Delete(&lukHeritageMetaverse).Error
	return err
}

// UpdateLukHeritageMetaverse 更新LukHeritageMetaverse记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseService *LukHeritageMetaverseService) UpdateLukHeritageMetaverse(lukHeritageMetaverse luk.LukHeritageMetaverse) (err error) {
	err = global.GVA_DB.Save(&lukHeritageMetaverse).Error
	return err
}

// GetLukHeritageMetaverse 根据id获取LukHeritageMetaverse记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseService *LukHeritageMetaverseService) GetLukHeritageMetaverse(id uint) (lukHeritageMetaverse luk.LukHeritageMetaverse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukHeritageMetaverse).Error
	return
}

// GetLukHeritageMetaverseInfoList 分页获取LukHeritageMetaverse记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageMetaverseService *LukHeritageMetaverseService) GetLukHeritageMetaverseInfoList(info lukReq.LukHeritageMetaverseSearch) (list []luk.LukHeritageMetaverse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukHeritageMetaverse{})
	var lukHeritageMetaverses []luk.LukHeritageMetaverse
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

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukHeritageMetaverses).Error
	return lukHeritageMetaverses, total, err
}
