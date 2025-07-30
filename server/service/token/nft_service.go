package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	"math/big"
)

type NftTokenService struct {
}

func (n *NftTokenService) OwnerOf(tokenId uint) (address common.Address, err error) {
	tokenHexAddress := NewToken(enum.NFT)
	instance, err := token.NewTigerToken(tokenHexAddress, client)
	if err != nil {
		return
	}
	address, err = instance.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	return
}

func (n *NftTokenService) GetRandom() (random string, err error) {
	tokenHexAddress := NewToken(enum.NFT)
	instance, err := token.NewTigerToken(tokenHexAddress, client)
	if err != nil {
		return
	}

	random, err = instance.GetRandom(&bind.CallOpts{
		From: common.HexToAddress(global.GVA_CONFIG.Wallet.Luk),
	})
	return
}
