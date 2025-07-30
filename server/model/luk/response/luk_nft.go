package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type LukNftList struct {
	ID         uint            `json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	Name       string          `json:"name"`
	Pic        string          `json:"pic"`
	PicKey     string          `json:"picKey"`
	Address    string          `json:"address"`
	Price      decimal.Decimal `json:"price"`
	Status     int             `json:"status"`
	Blacklist  bool            `json:"blacklist"`
	Prohibit   bool            `json:"prohibit"`
	IsDividend bool            `json:"isDividend"`
	AreaMax    decimal.Decimal `json:"areaMax"`
	AreaMin    decimal.Decimal `json:"areaMin"`
}
