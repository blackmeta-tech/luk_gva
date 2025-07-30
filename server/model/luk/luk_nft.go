// 自动生成模板LukNft
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukNft 结构体
type LukNft struct {
	global.GVA_MODEL
	Name       string          `json:"name" form:"name" gorm:"column:name;comment:NFT名称;size:64;"`
	Pic        string          `json:"pic" form:"pic" gorm:"column:pic;comment:图片路径;size:500;"`
	PicKey     string          `json:"picKey" form:"picKey" gorm:"column:pic_key;comment:图片key;size:64;"`
	Address    string          `json:"address" form:"address" gorm:"column:address;comment:归属用户地址;size:255;index:idx_name_address"`
	Price      decimal.Decimal `json:"price" form:"price" gorm:"column:price;comment:上架价格;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Status     int             `json:"status" form:"status" gorm:"column:status;comment:上架状态 0正常 -1下架;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Blacklist  bool            `json:"blacklist" form:"blacklist" gorm:"column:blacklist;comment:1.5Ku黑名单 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
	Prohibit   bool            `json:"prohibit" form:"prohibit" gorm:"column:prohibit;comment:禁止分红 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
	IsDividend bool            `json:"isDividend" form:"isDividend" gorm:"column:is_dividend;comment:分红结束 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"`
}

// TableName LukNft 表名
func (LukNft) TableName() string {
	return "luk_nft"
}
