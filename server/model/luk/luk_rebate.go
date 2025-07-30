// 自动生成模板LukRebate
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukRebate 结构体
type LukRebate struct {
	global.GVA_MODEL
	Address    string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;"`
	Balance    decimal.Decimal `json:"balance" form:"balance" gorm:"column:balance;comment:释放总余额;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Remaining  decimal.Decimal `json:"remaining" form:"remaining" gorm:"column:remaining;comment:未释放的余额;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Remark     string          `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	IsDividend bool            `json:"isDividend" form:"isDividend" gorm:"column:is_dividend;comment:分红结束 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
}

// TableName LukRebate 表名
func (LukRebate) TableName() string {
	return "luk_rebate"
}
