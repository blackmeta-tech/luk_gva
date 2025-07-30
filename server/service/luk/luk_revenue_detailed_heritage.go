package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type LukRevenueDetailedHeritageService struct {
}

// CreateLukRevenueDetailedHeritage 创建LukRevenueDetailedHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedHeritageService) CreateLukRevenueDetailedHeritage(lukRevenueDetailed luk.LukRevenueDetailedHeritage) (err error) {
	err = global.GVA_DB.Create(&lukRevenueDetailed).Error
	return err
}

// GetLukRevenueDetailedHeritage 根据id获取LukRevenueDetailedHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedHeritageService) GetLukRevenueDetailedHeritage(id uint) (lukRevenueDetailed luk.LukRevenueDetailedHeritage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenueDetailed).Error
	return
}

func (lukRevenueDetailedService *LukRevenueDetailedHeritageService) _Query(info lukReq.LukRevenueDetailedHeritageSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenueDetailedHeritage{})
	if info.ID > 0 {
		db.Where("id = ? ", info.ID)
	}
	if info.Address != "" {
		db.Where("address = ? ", info.Address)
	}
	if info.Type > 0 {
		db.Where("type = ? ", info.Type)
	}
	return
}

// GetLukRevenueDetailedHeritageInfoList 分页获取LukRevenueDetailedHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedHeritageService) GetLukRevenueDetailedHeritageInfoList(info lukReq.LukRevenueDetailedHeritageSearch) (list []luk.LukRevenueDetailedHeritage, summaryLuk, summaryUsdt decimal.Decimal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := lukRevenueDetailedService._Query(info)
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error

	var lukRevenueDetailed luk.LukRevenueDetailedHeritage
	err = lukRevenueDetailedService._Query(info).Select("sum(luk) as luk, sum(usdt) as usdt").First(&lukRevenueDetailed).Error
	if err == nil {
		summaryLuk = lukRevenueDetailed.Luk
		summaryUsdt = lukRevenueDetailed.Usdt
	}
	return
}
