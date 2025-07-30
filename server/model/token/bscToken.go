package token

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/shopspring/decimal"
)

// 合约转账实体类
type Transfer struct {
	From       string          `json:"from" form:"from"`             // 从（转帐地址）
	To         string          `json:"to" form:"to"`                 // 到（收款用户地址）
	For        decimal.Decimal `json:"for" form:"for"`               // 转帐金额
	Contract   string          `json:"contract" form:"contract"`     // 合约Token
	Type       emun.TokenType  `json:"type" form:"type"`             // 合约枚举(2:USDT,3:LUK)
	TxHash     string          `json:"txHash" form:"txHash"`         // 交易Hash
	StartBlock int64           `json:"startBlock" form:"startBlock"` // 开始区块（区块高度）
	EndBlocke  int64           `json:"endBlocke" form:"endBlocke"`   // 结束区块（区块高度）
}
