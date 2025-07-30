// 自动生成模板LukWithdrawApply
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukWithdrawApply 结构体
// 如果含有time.Time 请自行import time包
type LukWithdrawApply struct {
	global.GVA_MODEL
	Amount        decimal.Decimal  `json:"amount" form:"amount" gorm:"column:amount;comment:真实提现额度;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	AmountPrimary decimal.Decimal  `json:"amountPrimary" form:"amountPrimary" gorm:"column:amount_primary;comment:原始提现额度;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	Procedures    decimal.Decimal  `json:"procedures" form:"procedures" gorm:"column:procedures;comment:手续费;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	Status        *emun.StatusType `json:"status" form:"status" gorm:"column:status;comment:审核状态 0待审核 1成功 -1失败;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Address       string           `json:"address" form:"address" gorm:"column:address;comment:用户id;type:varchar(100) NOT NULL DEFAULT '';"`
	TokenType     *emun.TokenType  `json:"tokenType" form:"tokenType" gorm:"column:token_type;comment:币种类型;type:tinyint(2) NOT NULL DEFAULT '0';index:idx_name_user_id_token_type"`
}

// TableName LukWithdrawApply 表名
func (LukWithdrawApply) TableName() string {
	return "luk_withdraw_apply"
}
