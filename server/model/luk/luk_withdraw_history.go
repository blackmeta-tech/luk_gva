// 自动生成模板LukWithdrawHistory
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukWithdrawHistory 结构体
// 如果含有time.Time 请自行import time包
type LukWithdrawHistory struct {
	global.GVA_MODEL
	Address       string           `json:"address" form:"address" gorm:"column:address;comment:用户地址;type:varchar(100) NOT NULL DEFAULT '';index:idx_name_address;"`
	TokenType     *emun.TokenType  `json:"tokenType" form:"tokenType" gorm:"column:token_type;comment:币种类型;index:idx_name_token_type;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Amount        decimal.Decimal  `json:"amount" form:"amount" gorm:"column:amount;comment:真实提现额度;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	AmountPrimary decimal.Decimal  `json:"amountPrimary" form:"amountPrimary" gorm:"column:amount_primary;comment:原始提现额度;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	Procedures    decimal.Decimal  `json:"procedures" form:"procedures" gorm:"column:procedures;comment:手续费;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000'"`
	Status        *emun.StatusType `json:"status" form:"status" gorm:"column:status;comment:交易状态  0等待转账 1成功 -1失败;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Time          *time.Time       `json:"time" form:"time" gorm:"column:time;comment:提现时间;"`
	SignedTx      string           `json:"signedTx"  form:"signedTx" gorm:"column:signed_tx;comment:交易哈希;type:varchar(100) NOT NULL DEFAULT '';index:idx_name_signed_tx"`
	Number        int              `json:"number" form:"number" gorm:"column:number;type:int;not null;default:0;comment:确认次数;"`
}

// TableName LukWithdrawHistory 表名
func (LukWithdrawHistory) TableName() string {
	return "luk_withdraw_history"
}
