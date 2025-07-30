package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type LukRevenueNftService struct {
}

// CreateLukRevenueNft 创建LukRevenueNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueNftService) CreateLukRevenueNft(lukRevenue luk.LukRevenueNft) (err error) {
	err = global.GVA_DB.Create(&lukRevenue).Error
	return err
}

// GetLukRevenueNft 根据id获取LukRevenueNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueNftService) GetLukRevenueNft(id uint) (lukRevenue luk.LukRevenueNft, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenue).Error
	return
}

func (lukRevenueService *LukRevenueNftService) _Query(info lukReq.LukRevenueNftSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenueNft{})
	if info.ID > 0 {
		db.Where("id = ? ", info.ID)
	}
	if info.Address != "" {
		db.Where("address = ? ", info.Address)
	}
	if info.TokenType > 0 {
		db.Where("token_type = ? ", info.TokenType)
	}
	if info.Type > 0 {
		db.Where("type = ? ", info.Type)
	}
	return
}

// GetLukRevenueNftInfoList 分页获取LukRevenueNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueNftService) GetLukRevenueNftInfoList(info lukReq.LukRevenueNftSearch) (list []luk.LukRevenueNft, summary decimal.Decimal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := lukRevenueService._Query(info)
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc, id desc").Find(&list).Error
	var lukRevenue luk.LukRevenueNft
	err = lukRevenueService._Query(info).Select("sum(amount) as amount").First(&lukRevenue).Error
	if err == nil {
		summary = lukRevenue.Amount
	}
	return
}

func (lukRevenueService *LukRevenueNftService) UpdateLukRevenueNftBatch(tx *gorm.DB, data map[string]luk.LukRevenueNft) (err error) {
	for _, item := range data {
		revenue := luk.LukRevenueNft{}
		_ = tx.Where("address = ? and type = ? and token_type = ?", item.Address, item.Type, item.TokenType).First(&revenue).Error
		revenue.UpdatedAt = time.Now()
		if revenue.ID == 0 {
			revenue.Address = item.Address
			revenue.Type = item.Type
			revenue.TokenType = item.TokenType
			revenue.Amount = item.Amount
			revenue.CreatedAt = time.Now()
			err = tx.Debug().Create(&revenue).Error
			if err != nil {
				return
			}
		} else {
			revenue.Amount = revenue.Amount.Add(item.Amount)
			err = tx.Debug().Save(revenue).Error
			if err != nil {
				return
			}
		}
	}
	return
}
