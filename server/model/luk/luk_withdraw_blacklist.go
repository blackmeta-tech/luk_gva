// 自动生成模板LukWithdrawBlacklist
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukWithdrawBlacklist 结构体
type LukWithdrawBlacklist struct {
	global.GVA_MODEL
	Address string `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:200;index:idx_name_user_address"`
	Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:200;"`
}

// TableName LukWithdrawBlacklist 表名
func (LukWithdrawBlacklist) TableName() string {
	return "luk_withdraw_blacklist"
}
