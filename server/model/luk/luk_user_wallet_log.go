// 自动生成模板LukUserWalletLog
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukUserWalletLog 结构体
type LukUserWalletLog struct {
	global.GVA_MODEL
	Address string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;index:idx_name_address"`
	ErrMsg  string          `json:"errMsg" form:"errMsg" gorm:"column:err_msg;comment:修改报错;size:255;"`
	OldLuk  decimal.Decimal `json:"oldLuk" form:"oldLuk" gorm:"column:old_luk;comment:修改前LUK;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	NewLuk  decimal.Decimal `json:"newLuk" form:"newLuk" gorm:"column:new_luk;comment:修改后LUK;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	OldUsdt decimal.Decimal `json:"oldUsdt" form:"oldUsdt" gorm:"column:old_usdt;comment:修改前USDT;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	NewUsdt decimal.Decimal `json:"newUsdt" form:"newUsdt" gorm:"column:new_usdt;comment:修改后USDT;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
}

// TableName LukUserWalletLog 表名
func (LukUserWalletLog) TableName() string {
	return "luk_user_wallet_log"
}
