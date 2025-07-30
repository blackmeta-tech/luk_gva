// 自动生成模板LukNftBuy
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukNftBuy 结构体
type LukNftBuy struct {
	global.GVA_MODEL
	Address       string          `json:"address" form:"address" gorm:"column:address;comment:购买用户地址;size:255;index:idx_name_address"`
	InviteAddress string          `json:"inviteAddress" form:"inviteAddress" gorm:"column:invite_address;comment:邀请用户地址;size:255;"`
	NftId         uint            `json:"nftId" form:"nftId" gorm:"column:nft_id;comment:NFTID;size:10;"`
	NftName       string          `json:"nftName" form:"nftName" gorm:"column:nft_name;comment:NFT名称;size:255;"`
	Price         decimal.Decimal `json:"price" form:"price" gorm:"column:price;comment:购买花费金额;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	TxHash        string          `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:购买哈希;size:255;"`
}

// TableName LukNftBuy 表名
func (LukNftBuy) TableName() string {
	return "luk_nft_buy"
}
