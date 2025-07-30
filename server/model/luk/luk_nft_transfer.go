// 自动生成模板LukNftTransfer
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukNftTransfer 结构体
type LukNftTransfer struct {
	global.GVA_MODEL
	From   string `json:"from" form:"from" gorm:"column:from;comment:from用户地址;size:255;"`
	NftId  uint   `json:"nftId" form:"nftId" gorm:"column:nft_id;comment:NFTID;"`
	To     string `json:"to" form:"to" gorm:"column:to;comment:to用户地址;size:255;"`
	TxHash string `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:转移哈希;size:255;"`
}

// TableName LukNftTransfer 表名
func (LukNftTransfer) TableName() string {
	return "luk_nft_transfer"
}
