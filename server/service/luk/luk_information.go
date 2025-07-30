package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
)

type LukInformationList struct {
	luk.LukInformation
	TypeName string `json:"typeName"`
}
type LukInformationService struct {
}

// CreateLukInformation 创建LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) CreateLukInformation(lukDappInformation luk.LukInformation) (err error) {
	err = global.GVA_DB.Create(&lukDappInformation).Error
	return err
}

// DeleteLukInformation 删除LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) DeleteLukInformation(lukDappInformation luk.LukInformation) (err error) {
	err = global.GVA_DB.Delete(&lukDappInformation).Error
	return err
}

// DeleteLukInformationByIds 批量删除LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) DeleteLukInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukInformation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukInformation 更新LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) UpdateLukInformation(lukDappInformation luk.LukInformation) (err error) {
	err = global.GVA_DB.Save(&lukDappInformation).Error
	return err
}

// GetLukInformation 根据id获取LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) GetLukInformation(id uint) (lukDappInformation luk.LukInformation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukDappInformation).Error
	return
}

// GetLukInformationInfoList 分页获取LukInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationService *LukInformationService) GetLukInformationInfoList(info lukReq.LukInformationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukInformation{})
	var lukDappInformations []LukInformationList
	if info.Status != nil {
		if *info.Status > 0 {
			db = db.Where("luk_information.status = ? ", info.Status)
		}
	}
	if info.TypeId != nil {
		if *info.TypeId > 0 {
			db = db.Where("type_id = ? ", info.TypeId)
		}
	}
	if info.ID > 0 {
		db = db.Where("luk_information.id = ? ", info.ID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Select("luk_information.*, type.name_ch type_name").Joins("LEFT JOIN luk_information_type type ON type.id = luk_information.type_id", "type").Limit(limit).Offset(offset).Order("id desc").Find(&lukDappInformations).Error
	return lukDappInformations, total, err
}
