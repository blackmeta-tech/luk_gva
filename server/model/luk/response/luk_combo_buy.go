package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type LukComboBuyList struct {
	ID         uint            `json:"id"`
	ComboId    uint            `json:"comboId"` //套餐ID
	Pic        string          `json:"pic"`     //图片
	BuyAt      time.Time       `json:"buyAt"`
	Days       int             `json:"days"`
	MaturityAt time.Time       `json:"maturityAt"`
	Price      decimal.Decimal `json:"price"`
	PriceU     decimal.Decimal `json:"priceU"`
	Rate       decimal.Decimal `json:"rate"`
	Repeat     int             `json:"repeat" form:"repeat"`
}
