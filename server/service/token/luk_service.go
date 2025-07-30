package token

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	"github.com/flipped-aurora/gin-vue-admin/server/model/utils"
	"github.com/shopspring/decimal"
	"math/big"
)

type LukTokenService struct {
}

func (d *LukTokenService) NewlukToken() (instance *token.LukToken, err error) {
	tokenHexAddress := NewToken(enum.LUK)
	instance, err = token.NewLukToken(tokenHexAddress, client)
	return
}

func (d *LukTokenService) GetBalanceOf(tokenAddress, from common.Address) (valueD decimal.Decimal, err error) {
	instance, err := token.NewLukToken(tokenAddress, client)
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

func (d *LukTokenService) TransferFrom(opts *bind.TransactOpts, tokenAddress, to common.Address, amount *big.Int) (txHash common.Hash, err error) {
	instance, err := token.NewLukToken(tokenAddress, client)
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

//零手续费白名单
func (d *LukTokenService) SetWhiteList(addres common.Address, isCharge int) (err error) {
	instance, err := d.NewlukToken()
	if err != nil {
		return
	}
	opts, err := GetTransactOpts(global.GVA_CONFIG.Wallet.LukPrivateKey.String())
	if err != nil {
		return
	}
	isTrade := false
	if isCharge == 1 {
		isTrade = true
	}
	_, err = instance.SetWhiteList(opts, addres, isTrade)
	return
}

//查看某人是否在零手续费白名单
func (d *LukTokenService) IsWhiteList(address string) (status bool, err error) {
	instance, err := d.NewlukToken()
	opts := bind.CallOpts{}
	status, err = instance.GetWhiteList(&opts, common.HexToAddress(address))
	fmt.Println("=====IsWhiteList", status, err)
	return
}

func (d *LukTokenService) Burn(amount decimal.Decimal) (hash common.Hash, err error) {
	instance, err := d.NewlukToken()
	if err != nil {
		return
	}
	opts, err := GetTransactOpts(global.GVA_CONFIG.Wallet.LukPrivateKey.String())
	if err != nil {
		return
	}
	trans, err := instance.Burn(opts, common.HexToAddress(global.GVA_CONFIG.Wallet.Destroy), utils.ToWei(amount))
	if err != nil {
		return
	}
	hash = trans.Hash()
	return
}
