// 自动生成模板LukHeritageMetaverse
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukHeritageMetaverse 结构体
type LukHeritageMetaverse struct {
	global.GVA_MODEL
	Address    string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;"`
	Balance    decimal.Decimal `json:"balance" form:"balance" gorm:"column:balance;comment:遗产余额;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Remaining  decimal.Decimal `json:"remaining" form:"remaining" gorm:"column:remaining;comment:未分红的遗产;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	ComboUsdt  decimal.Decimal `json:"comboUsdt" form:"comboUsdt" gorm:"column:combo_usdt;comment:购买套餐花费Usdt;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	LpUsdt     decimal.Decimal `json:"lpUsdt" form:"lpUsdt" gorm:"column:lp_usdt;comment:lp对应Usdt的价值;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Remark     string          `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	IsDividend bool            `json:"isDividend" form:"isDividend" gorm:"column:is_dividend;comment:分红结束 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
}

// TableName LukHeritageMetaverse 表名
func (LukHeritageMetaverse) TableName() string {
	return "luk_heritage_metaverse"
}
