package request

import (
	"github.com/ethereum/go-ethereum/common"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/shopspring/decimal"
)

//TransferFrom 转账事件
type TransferFrom struct {
	TokenType  enum.TokenType  `json:"tokenType"`
	PrivateKey string          `json:"privateKey"` //地址私钥
	ToAddress  common.Address  `json:"toAddress"`  //转账到
	Amount     decimal.Decimal `json:"amount"`
	TxHash     common.Hash     `json:"txHash,omitempty"`
}
