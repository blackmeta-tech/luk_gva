package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type LukRevenueDetailedService struct {
}

// CreateLukRevenueDetailed 创建LukRevenueDetailed记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedService) CreateLukRevenueDetailed(lukRevenueDetailed luk.LukRevenueDetailed) (err error) {
	err = global.GVA_DB.Create(&lukRevenueDetailed).Error
	return err
}

// GetLukRevenueDetailed 根据id获取LukRevenueDetailed记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedService) GetLukRevenueDetailed(id uint) (lukRevenueDetailed luk.LukRevenueDetailed, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukRevenueDetailed).Error
	return
}

func (lukRevenueDetailedService *LukRevenueDetailedService) _Query(info lukReq.LukRevenueDetailedSearch) (db *gorm.DB) {
	db = global.GVA_DB.Debug().Model(&luk.LukRevenueDetailed{})
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
	if info.Category > 0 {
		db.Where("category = ? ", info.Category)
	}
	return
}

// GetLukRevenueDetailedInfoList 分页获取LukRevenueDetailed记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukRevenueDetailedService *LukRevenueDetailedService) GetLukRevenueDetailedInfoList(info lukReq.LukRevenueDetailedSearch) (list []luk.LukRevenueDetailed, summary decimal.Decimal, total int64, err error) {
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

	var lukRevenueDetailed luk.LukRevenueDetailed
	err = lukRevenueDetailedService._Query(info).Select("sum(amount) as amount").First(&lukRevenueDetailed).Error
	if err == nil {
		summary = lukRevenueDetailed.Amount
	}
	return
}

//统计地址LP分佣分红累计U量
func (lukRevenueDetailedService *LukRevenueDetailedService) QueryByAddress(address string) (relatively decimal.Decimal) {
	info := luk.LukRevenueDetailed{}
	err := global.GVA_DB.Debug().Model(&luk.LukRevenueDetailed{}).Select("sum(relatively) as relatively").
		Where("address = ? and type = ?", address, emun.RevenueTypeLpRebate).First(&info).Error
	if err != nil {
		return
	}
	relatively = info.Relatively.Mul(global.GVA_LUK_CONFIG.LpRebate.Div(decimal.NewFromInt(100)))
	return
}
