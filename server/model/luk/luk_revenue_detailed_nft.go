// 自动生成模板LukRevenueDetailedNft
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukRevenueDetailedNft 结构体
type LukRevenueDetailedNft struct {
	global.GVA_MODEL
	Address    string              `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;index:idx_name_address"`
	Amount     decimal.Decimal     `json:"amount" form:"amount" gorm:"column:amount;comment:币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Relatively decimal.Decimal     `json:"relatively" form:"relatively" gorm:"column:relatively;comment:对应的其他币种数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Type       emun.RevenueTypeNft `json:"type" form:"type" gorm:"column:type;comment:类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
	TokenType  emun.TokenType      `json:"tokenType" form:"tokenType" gorm:"column:token_type;comment:代币类型 (1:USDT,2:LUK);type:tinyint NOT NULL DEFAULT '0';'"`
	Remark     string              `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	Source     string              `json:"source" form:"source" gorm:"column:source;comment:来源;type:varchar(255) NOT NULL DEFAULT '--';"`
}

// TableName LukRevenueDetailedNft 表名
func (LukRevenueDetailedNft) TableName() string {
	return "luk_revenue_detailed_nft"
}
