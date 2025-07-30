package luk

import (
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type LukRevenueService struct {
}

// CreateLukRevenue 创建LukRevenue记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueService) CreateLukRevenue(lukRevenue luk.LukRevenue) (err error) {
	err = global.GVA_DB.Create(&lukRevenue).Error
	return err
}

// GetLukRevenue 根据id获取LukRevenue记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueService) GetLukRevenue(id uint) (lukRevenue luk.LukRevenue, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenue).Error
	return
}

func (lukRevenueService *LukRevenueService) _Query(info lukReq.LukRevenueSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenue{})
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
	if info.Category > 0 {
		db.Where("category = ? ", info.Category)
	}
	return
}

// GetLukRevenueInfoList 分页获取LukRevenue记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueService) GetLukRevenueInfoList(info lukReq.LukRevenueSearch) (list []luk.LukRevenue, summary decimal.Decimal, total int64, err error) {
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
	var lukRevenue luk.LukRevenue
	err = lukRevenueService._Query(info).Select("sum(amount) as amount").First(&lukRevenue).Error
	if err == nil {
		summary = lukRevenue.Amount
	}
	return
}

func (lukRevenueService *LukRevenueService) UpdateLukRevenueBatch(tx *gorm.DB, data map[string]luk.LukRevenue) (err error) {
	for _, item := range data {
		revenue := luk.LukRevenue{}
		_ = tx.Where("address = ? and type = ? and token_type = ?", item.Address, item.Type, item.TokenType).First(&revenue).Error
		revenue.UpdatedAt = time.Now()
		if revenue.ID == 0 {
			revenue.Address = item.Address
			revenue.Type = item.Type
			revenue.Category = item.Category
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

func (lukRevenueService *LukRevenueService) GetAll() (data map[enum.RevenueType]decimal.Decimal) {
	data = make(map[enum.RevenueType]decimal.Decimal, 0)
	list := []luk.LukRevenue{}
	err := global.GVA_DB.Select("sum(amount) as amount, type").Group("type").Find(&list).Error
	if err != nil {
		return
	}
	for _, item := range list {
		data[item.Type] = item.Amount
	}
	return
}
