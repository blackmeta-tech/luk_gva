package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

//技术钱包记录
type LukReflow struct {
	global.GVA_MODEL
	BuyId  uint            `json:"buyId" form:"buyId" gorm:"column:buy_id;comment:对应购买记录ID;size:11;"`
	Amount decimal.Decimal `json:"amount" form:"amount" gorm:"column:amount;comment:币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Type   emun.ReflowType `json:"type" form:"type" gorm:"column:type;comment:类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
}

// TableName LukDepositNode 表名
func (LukReflow) TableName() string {
	return "luk_reflow"
}
