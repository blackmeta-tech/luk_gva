// 自动生成模板LukHeritageLp
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukHeritageLp 结构体
type LukHeritageLp struct {
	global.GVA_MODEL
	Address    string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;"`
	Balance    decimal.Decimal `json:"balance" form:"balance" gorm:"column:balance;comment:遗产余额;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Remaining  decimal.Decimal `json:"remaining" form:"remaining" gorm:"column:remaining;comment:未分红的遗产;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	IsDividend bool            `json:"isDividend" form:"isDividend" gorm:"column:is_dividend;comment:分红结束 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Remark     string          `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}

// TableName LukHeritageLp 表名
func (LukHeritageLp) TableName() string {
	return "luk_heritage_lp"
}
