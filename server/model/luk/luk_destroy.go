package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

//luk销毁记录
type LukDestroy struct {
	global.GVA_MODEL
	Address string           `json:"address" form:"address" gorm:"column:address;comment:购买用户地址;size:255;index:idx_name_address"`
	BuyId   uint             `json:"buyId" form:"buyId" gorm:"column:buy_id;comment:对应购买记录ID;size:11;"`
	TxHash  string           `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:tx 哈希;size:191;index:idx_name_txhash"`
	Luk     decimal.Decimal  `json:"luk" form:"luk" gorm:"column:luk;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:luk个数;"`
	Status  *emun.StatusType `json:"status" form:"status" gorm:"type:tinyint(2) NOT NULL DEFAULT '0';column:status;comment:状态 1成功 0等待 -1失败;"`
}

// TableName LukDestroy 表名
func (LukDestroy) TableName() string {
	return "luk_destroy"
}
