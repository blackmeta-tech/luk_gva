package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type LukReflowLp struct {
	global.GVA_MODEL
	TxHash string            `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:对应产生的哈希;size:255;"`
	Amount decimal.Decimal   `json:"amount" form:"amount" gorm:"column:amount;comment:币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Type   emun.ReflowLpType `json:"type" form:"type" gorm:"column:type;comment:类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
}

// TableName LukDepositNode 表名
func (LukReflowLp) TableName() string {
	return "luk_reflow_lp"
}
