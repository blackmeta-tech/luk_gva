// 自动生成模板LukRevenueHeritage
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukRevenueHeritage 结构体
type LukRevenueHeritage struct {
	global.GVA_MODEL
	Address string            `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;index:idx_name_address"`
	Luk     decimal.Decimal   `json:"luk" form:"luk" gorm:"column:luk;comment:luk币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Usdt    decimal.Decimal   `json:"usdt" form:"usdt" gorm:"column:usdt;comment:币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Type    emun.HeritageType `json:"type" form:"type" gorm:"column:type;comment:类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
}

// TableName LukRevenueHeritage 表名
func (LukRevenueHeritage) TableName() string {
	return "luk_revenue_heritage"
}
