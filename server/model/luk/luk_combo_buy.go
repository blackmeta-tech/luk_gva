// 自动生成模板LukComboBuy
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

// LukComboBuy 结构体
type LukComboBuy struct {
	global.GVA_MODEL
	Address       string          `json:"address" form:"address" gorm:"column:address;comment:购买用户地址;size:255;index:idx_name_address"`
	BuyAt         time.Time       `json:"buyAt" form:"buyAt" gorm:"column:buy_at;comment:购买时间;"`
	ComboId       uint            `json:"comboId" form:"comboId" gorm:"column:combo_id;comment:套餐ID;size:10;"`
	Days          int             `json:"days" form:"days" gorm:"column:days;comment:套餐天数;size:10;"`
	MaturityAt    time.Time       `json:"maturityAt" form:"maturityAt" gorm:"column:maturity_at;comment:到期时间;"`
	Price         decimal.Decimal `json:"price" form:"price" gorm:"column:price;comment:购入价格;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	PriceUsdt     decimal.Decimal `json:"priceUsdt" form:"priceUsdt" gorm:"column:price_usdt;comment:购入价格等比例usdt;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Rate          decimal.Decimal `json:"rate" form:"rate" gorm:"column:rate;comment:收益率百分比;type:decimal(11,2) NOT NULL DEFAULT '0.00';"`
	Repeat        int             `json:"repeat" form:"repeat" gorm:"column:repeat;comment:复购状态 0不复购 1复购;"`
	Status        int             `json:"status" form:"status" gorm:"column:status;comment:购买状态 1已复购 0正常 -1到期 -2复购失败;"`
	StatusMsg     string          `json:"statusMsg" form:"statusMsg" gorm:"column:status_msg;comment:失败原因;size:255;"`
	ErrNum        int             `json:"errNum" form:"errNum" gorm:"column:err_num;comment:失败次数;type:smallint NOT NULL DEFAULT '0';"`
	Dividend      int             `json:"dividend" form:"dividend" gorm:"column:dividend;comment:分红状态  0还未分红 1已分红;size:10;"`
	Extend        int             `json:"extend" form:"extend" gorm:"column:extend;comment:延长状态 0还未延长 1已延长;type:smallint NOT NULL DEFAULT '0';"`
	DividendPrice decimal.Decimal `json:"dividendPrice" form:"dividendPrice" gorm:"column:dividend_price;comment:分红时币价;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Destroy       int             `json:"destroy" form:"destroy" gorm:"column:destroy;comment:销毁状态 0还未销毁 1已销毁;size:10;"`
	Type          int             `json:"type" form:"type" gorm:"column:type;comment:购买类型 0正常购买 1复购;size:10;"`
	TxHash        string          `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:购买成功哈希;size:255;"`
}

// TableName LukComboBuy 表名
func (LukComboBuy) TableName() string {
	return "luk_combo_buy"
}
