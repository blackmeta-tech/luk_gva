package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

//  Wei 工具 转 Wei(big.Int) Decimal
func ToDecimal(bal *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(bal, -18)
}

//  Wei 工具 Decimal 转 Wei (big.Int)
func ToWei(bal decimal.Decimal) *big.Int {
	return bal.Mul(decimal.NewFromFloat(1000000000000000000)).BigInt()
}
