// 自动生成模板LukRevenueDetailedHeritage
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukRevenueDetailedHeritage 结构体
type LukRevenueDetailedHeritage struct {
	global.GVA_MODEL
	Address    string            `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;index:idx_name_address"`
	Luk        decimal.Decimal   `json:"luk" form:"luk" gorm:"column:luk;comment:luk币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Usdt       decimal.Decimal   `json:"usdt" form:"usdt" gorm:"column:usdt;comment:币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Type       emun.HeritageType `json:"type" form:"type" gorm:"column:type;comment:类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
	Percentage decimal.Decimal   `json:"percentage" form:"percentage" gorm:"column:percentage;comment:分配比例;type:decimal(32,4) NOT NULL DEFAULT '0.000000000000000000';"`
	Remark     string            `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}

// TableName LukRevenueDetailedHeritage 表名
func (LukRevenueDetailedHeritage) TableName() string {
	return "luk_revenue_detailed_heritage"
}
