package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type LukRevenueHeritageService struct {
}

// CreateLukRevenueHeritage 创建LukRevenueHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueHeritageService) CreateLukRevenueHeritage(lukRevenue luk.LukRevenueHeritage) (err error) {
	err = global.GVA_DB.Create(&lukRevenue).Error
	return err
}

// GetLukRevenueHeritage 根据id获取LukRevenueHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueHeritageService) GetLukRevenueHeritage(id uint) (lukRevenue luk.LukRevenueHeritage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenue).Error
	return
}

func (lukRevenueService *LukRevenueHeritageService) _Query(info lukReq.LukRevenueHeritageSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenueHeritage{})
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

// GetLukRevenueHeritageInfoList 分页获取LukRevenueHeritage记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueService *LukRevenueHeritageService) GetLukRevenueHeritageInfoList(info lukReq.LukRevenueHeritageSearch) (list []luk.LukRevenueHeritage, summaryLuk, summaryUsdt decimal.Decimal, total int64, err error) {
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
	var lukRevenue luk.LukRevenueHeritage
	err = lukRevenueService._Query(info).Select("sum(luk) as luk, sum(usdt) as usdt").First(&lukRevenue).Error
	if err == nil {
		summaryLuk = lukRevenue.Luk
		summaryUsdt = lukRevenue.Usdt
	}
	return
}

func (lukRevenueService *LukRevenueHeritageService) UpdateLukRevenueHeritageBatch(tx *gorm.DB, data map[string]luk.LukRevenueHeritage) (err error) {
	for _, item := range data {
		revenue := luk.LukRevenueHeritage{}
		_ = tx.Where("address = ? and type = ?", item.Address, item.Type).First(&revenue).Error
		revenue.UpdatedAt = time.Now()
		if revenue.ID == 0 {
			revenue.Address = item.Address
			revenue.Type = item.Type
			revenue.Luk = item.Luk
			revenue.Usdt = item.Usdt
			revenue.CreatedAt = time.Now()
			err = tx.Debug().Create(&revenue).Error
			if err != nil {
				return
			}
		} else {
			revenue.Luk = revenue.Luk.Add(item.Luk)
			revenue.Usdt = revenue.Usdt.Add(item.Usdt)
			err = tx.Debug().Save(revenue).Error
			if err != nil {
				return
			}
		}
	}
	return
}

//查找汇总数量
func (lukRevenueService *LukRevenueHeritageService) GetbyAddress(address string, _type emun.HeritageType) (amount decimal.Decimal) {
	revenue := luk.LukRevenueHeritage{}
	err := global.GVA_DB.Debug().Where("address = ? and type = ?", address, _type).First(&revenue).Error
	if err == nil {
		amount = revenue.Usdt
	}
	return
}
