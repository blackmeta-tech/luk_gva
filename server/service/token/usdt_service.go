package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	"github.com/flipped-aurora/gin-vue-admin/server/model/utils"
	"github.com/shopspring/decimal"
	"math/big"
)

type UsdtTokenService struct {
}

func (d *UsdtTokenService) GetBalanceOf(tokenAddress, from common.Address) (valueD decimal.Decimal, err error) {
	instance, err := token.NewUsdtToken(tokenAddress, client)
	if err != nil {
		return
	}
	address := from
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err == nil {
		// 转 Wei
		valueD = utils.ToDecimal(bal)
	}
	return
}

func (d *UsdtTokenService) TransferFrom(opts *bind.TransactOpts, tokenAddress, to common.Address, amount *big.Int) (txHash common.Hash, err error) {
	instance, err := token.NewUsdtToken(tokenAddress, client)
	if err != nil {
		return
	}
	trans, err := instance.Transfer(opts, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	txHash = trans.Hash()
	return
}
