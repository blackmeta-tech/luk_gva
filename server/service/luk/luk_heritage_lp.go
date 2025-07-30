package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukHeritageLpService struct {
}

// CreateLukHeritageLp 创建LukHeritageLp记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageLpService *LukHeritageLpService) CreateLukHeritageLp(lukHeritageLp luk.LukHeritageLp) (err error) {
	lukHeritageLp.Address = strings.ToLower(lukHeritageLp.Address)
	info := luk.LukHeritageLp{}
	global.GVA_DB.Where("address = ?", lukHeritageLp.Address).First(&info)
	if info.ID > 0 {
		err = errors.New("该用户地址已录入")
		return
	}
	lukHeritageLp.Remaining = lukHeritageLp.Balance
	err = global.GVA_DB.Create(&lukHeritageLp).Error
	return err
}

// DeleteLukHeritageLp 删除LukHeritageLp记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageLpService *LukHeritageLpService) DeleteLukHeritageLp(lukHeritageLp luk.LukHeritageLp) (err error) {
	err = global.GVA_DB.Delete(&lukHeritageLp).Error
	return err
}

// UpdateLukHeritageLp 更新LukHeritageLp记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageLpService *LukHeritageLpService) UpdateLukHeritageLp(lukHeritageLp luk.LukHeritageLp) (err error) {
	err = global.GVA_DB.Save(&lukHeritageLp).Error
	return err
}

// GetLukHeritageLp 根据id获取LukHeritageLp记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageLpService *LukHeritageLpService) GetLukHeritageLp(id uint) (lukHeritageLp luk.LukHeritageLp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukHeritageLp).Error
	return
}

// GetLukHeritageLpInfoList 分页获取LukHeritageLp记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukHeritageLpService *LukHeritageLpService) GetLukHeritageLpInfoList(info lukReq.LukHeritageLpSearch) (list []luk.LukHeritageLp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukHeritageLp{})
	var lukHeritageLps []luk.LukHeritageLp
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

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukHeritageLps).Error
	return lukHeritageLps, total, err
}
