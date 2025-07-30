package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type LukPoolBasic struct {
	global.GVA_MODEL
	Luk            decimal.Decimal `json:"luk" form:"luk" gorm:"column:luk;comment:luk数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Usdt           decimal.Decimal `json:"usdt" form:"usdt" gorm:"column:usdt;comment:usdt数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	DestructionLuk decimal.Decimal `json:"destructionLuk" form:"destructionLuk" gorm:"column:destruction_luk;comment:销毁的luk总量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
}

// TableName LukPoolBasic 表名
func (LukPoolBasic) TableName() string {
	return "luk_pool_basic"
}
