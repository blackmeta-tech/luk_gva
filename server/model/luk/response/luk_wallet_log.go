package response

import (
	"github.com/shopspring/decimal"
)

type LukWalletList struct {
	Abbr       string          `json:"abbr"`       //钱包名称英文定义
	Name       string          `json:"name"`       //钱包名称
	Address    string          `json:"address"`    //钱包地址
	PrivateKey string          `json:"privateKey"` //私钥
	Luk        decimal.Decimal `json:"luk"`
	Usdt       decimal.Decimal `json:"usdt"`
	Bnb        decimal.Decimal `json:"bnb"`
}

type BnbRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
