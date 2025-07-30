package response

import "github.com/shopspring/decimal"

type LukComboList struct {
	ComboId        uint            `json:"comboId"`        //套餐ID
	PriceMax       decimal.Decimal `json:"priceMax"`       //最大购买
	PriceMin       decimal.Decimal `json:"priceMin"`       //最小购买
	Rate           decimal.Decimal `json:"rate"`           //利率
	Limit          decimal.Decimal `json:"limit"`          //今日总额度
	Remaining      decimal.Decimal `json:"remaining"`      //今日剩余额度
	Days           int             `json:"days"`           //套餐类型
	Pic            string          `json:"pic"`            //图片
	Performance    decimal.Decimal `json:"performance"`    //达标业绩
	PerformanceNow decimal.Decimal `json:"performanceNow"` //当前业绩
	Reason         map[int]string  `json:"reason"`
}

type LukComboPerformance struct {
	Days        int             `json:"days"`
	Performance decimal.Decimal `json:"performance"`
	State       int             `json:"state"`
}
