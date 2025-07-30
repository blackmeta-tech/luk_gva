package response

import (
	"github.com/shopspring/decimal"
	"time"
)

type LukNftTransferReponse struct {
	UpdatedAt  time.Time       `json:"updatedAt" form:"updatedAt"`
	TransferAt time.Time       `json:"transferAt" form:"transferAt"`
	From       string          `json:"from" form:"from"`
	NftId      uint            `json:"nftId" form:"nftId"`
	To         string          `json:"to" form:"to"`
	TxHash     string          `json:"txHash" form:"txHash"`
	Name       string          `json:"name" form:"name"`
	Pic        string          `json:"pic" form:"pic"`
	Price      decimal.Decimal `json:"price" form:"price"`
}
