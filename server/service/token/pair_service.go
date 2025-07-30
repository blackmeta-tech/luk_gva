package token

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	"github.com/flipped-aurora/gin-vue-admin/server/model/utils"
	"github.com/shopspring/decimal"
	"strings"
)

type PairTokenService struct {
}

//获取储备量
func (pairTokenService *PairTokenService) GetReserves() (valueD decimal.Decimal, valueU decimal.Decimal, err error) {
	if client == nil {
		client = NewWss()
	}
	// 创建调用合约实例
	tokenHexAddress := common.HexToAddress(global.GVA_CONFIG.Bsc.Liquidity)
	instance, err := token.NewPairToken(tokenHexAddress, client)
	if err != nil {
		return
	}
	ss, err := instance.GetReserves(&bind.CallOpts{})
	fmt.Println("====", ss, err)
	if err != nil {
		fmt.Println("获取储备量错误", err)
		return
	}

	token0, err := instance.Token0(&bind.CallOpts{})
	if strings.ToLower(token0.String()) == strings.ToLower(global.GVA_CONFIG.Bsc.Luk) {
		//reserve0是Luk的储备量 reserve1是USDT的储备量
		valueD = utils.ToDecimal(ss.Reserve0)
		valueU = utils.ToDecimal(ss.Reserve1)
	} else {
		//reserve0是Luk的储备量 reserve1是Luk的储备量
		valueD = utils.ToDecimal(ss.Reserve1)
		valueU = utils.ToDecimal(ss.Reserve0)
	}
	return
}

//获取LP的总量
func (pairTokenService *PairTokenService) TotalSupply() (Lp decimal.Decimal, err error) {
	if client == nil {
		client = NewWss()
	}
	// 创建调用合约实例
	tokenHexAddress := common.HexToAddress(global.GVA_CONFIG.Bsc.Liquidity)
	instance, err := token.NewPairToken(tokenHexAddress, client)
	if err != nil {
		return
	}
	ss, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		fmt.Println("获取Lp量错误", err)
		return
	}
	Lp = utils.ToDecimal(ss)
	return
}
