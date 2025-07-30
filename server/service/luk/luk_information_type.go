package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
)

type LukInformationTypeService struct {
}

// CreateLukInformationType 创建LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) CreateLukInformationType(lukDappInformationType luk.LukInformationType) (err error) {
	err = global.GVA_DB.Create(&lukDappInformationType).Error
	return err
}

// DeleteLukInformationType 删除LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) DeleteLukInformationType(lukDappInformationType luk.LukInformationType) (err error) {
	err = global.GVA_DB.Delete(&lukDappInformationType).Error
	return err
}

// DeleteLukInformationTypeByIds 批量删除LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) DeleteLukInformationTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukInformationType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukInformationType 更新LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) UpdateLukInformationType(lukDappInformationType luk.LukInformationType) (err error) {
	err = global.GVA_DB.Save(&lukDappInformationType).Error
	return err
}

// GetLukInformationType 根据id获取LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) GetLukInformationType(id uint) (lukDappInformationType luk.LukInformationType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukDappInformationType).Error
	return
}

// GetLukInformationTypeInfoList 分页获取LukInformationType记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukDappInformationTypeService *LukInformationTypeService) GetLukInformationTypeInfoList(info lukReq.LukInformationTypeSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukInformationType{})
	var lukDappInformationTypes []luk.LukInformationType
	if info.Status != nil {
		if *info.Status > 0 {
			db = db.Where("status = ? ", info.Status)
		}
	}
	if info.ID > 0 {
		db = db.Where("id = ? ", info.ID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort_num asc,id desc").Find(&lukDappInformationTypes).Error
	return lukDappInformationTypes, total, err
}
