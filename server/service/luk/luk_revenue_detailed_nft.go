package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type LukRevenueDetailedNftService struct {
}

// CreateLukRevenueDetailedNft 创建LukRevenueDetailedNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedNftService) CreateLukRevenueDetailedNft(lukRevenueDetailed luk.LukRevenueDetailedNft) (err error) {
	err = global.GVA_DB.Create(&lukRevenueDetailed).Error
	return err
}

// GetLukRevenueDetailedNft 根据id获取LukRevenueDetailedNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedNftService) GetLukRevenueDetailedNft(id uint) (lukRevenueDetailed luk.LukRevenueDetailedNft, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenueDetailed).Error
	return
}

func (lukRevenueDetailedService *LukRevenueDetailedNftService) _Query(info lukReq.LukRevenueDetailedNftSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenueDetailedNft{})
	if info.ID > 0 {
		db.Where("id = ? ", info.ID)
	}
	if info.TokenType > 0 {
		db.Where("token_type = ? ", info.TokenType)
	}
	if info.Address != "" {
		db.Where("address = ? ", info.Address)
	}
	if info.Type > 0 {
		db.Where("type = ? ", info.Type)
	}
	return
}

// GetLukRevenueDetailedNftInfoList 分页获取LukRevenueDetailedNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedNftService) GetLukRevenueDetailedNftInfoList(info lukReq.LukRevenueDetailedNftSearch) (list []luk.LukRevenueDetailedNft, summary decimal.Decimal, total int64, err error) {
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

	var lukRevenueDetailed luk.LukRevenueDetailedNft
	err = lukRevenueDetailedService._Query(info).Select("sum(amount) as amount").First(&lukRevenueDetailed).Error
	if err == nil {
		summary = lukRevenueDetailed.Amount
	}
	return
}
