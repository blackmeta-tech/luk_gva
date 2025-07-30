// 自动生成模板LukCombo
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukCombo 结构体
type LukCombo struct {
	global.GVA_MODEL
	Days        int             `json:"days" form:"days" gorm:"column:days;comment:套餐天数;type:int(11) NOT NULL DEFAULT '0'"`
	Limit       decimal.Decimal `json:"limit" form:"limit" gorm:"column:limit;comment:当日限购;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Pic         string          `json:"pic" form:"pic" gorm:"column:pic;comment:图片路径;size:500;"`
	PicKey      string          `json:"picKey" form:"picKey" gorm:"column:pic_key;comment:图片key;size:64;"`
	PriceMax    decimal.Decimal `json:"priceMax" form:"priceMax" gorm:"column:price_max;comment:最高购入价格;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	PriceMin    decimal.Decimal `json:"priceMin" form:"priceMin" gorm:"column:price_min;comment:最低购入价格;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Rate        decimal.Decimal `json:"rate" form:"rate" gorm:"column:rate;comment:收益率百分比;type:decimal(11,2) NOT NULL DEFAULT '0.00';"`
	Status      int             `json:"status" form:"status" gorm:"column:status;comment:上架状态 0正常 -1下架;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Performance decimal.Decimal `json:"performance" form:"performance" gorm:"column:performance;comment:达标业绩;type:decimal(11,2) NOT NULL DEFAULT '0.00';"`
}

// TableName LukCombo 表名
func (LukCombo) TableName() string {
	return "luk_combo"
}
