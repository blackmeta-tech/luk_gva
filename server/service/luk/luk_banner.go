package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
)

type LukBannerService struct {
}

// CreateLukBanner 创建LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) CreateLukBanner(lukBanner luk.LukBanner) (err error) {
	err = global.GVA_DB.Create(&lukBanner).Error
	return err
}

// DeleteLukBanner 删除LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) DeleteLukBanner(lukBanner luk.LukBanner) (err error) {
	err = global.GVA_DB.Delete(&lukBanner).Error
	return err
}

// DeleteLukBannerByIds 批量删除LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) DeleteLukBannerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukBanner{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukBanner 更新LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) UpdateLukBanner(lukBanner luk.LukBanner) (err error) {
	err = global.GVA_DB.Save(&lukBanner).Error
	return err
}

// GetLukBanner 根据id获取LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) GetLukBanner(id uint) (lukBanner luk.LukBanner, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukBanner).Error
	return
}

// GetLukBannerInfoList 分页获取LukBanner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBannerService *LukBannerService) GetLukBannerInfoList(info lukReq.LukBannerSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukBanner{})
	var lukBanners []luk.LukBanner
	// 如果有条件搜索 下方会自动创建搜索语句
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
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukBanners).Error
	return lukBanners, total, err
}
