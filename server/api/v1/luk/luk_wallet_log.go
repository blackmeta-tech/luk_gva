package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type LukWalletLogApi struct {
}

var lukWalletLogService = service.ServiceGroupApp.LukServiceGroup.LukWalletLogService

//获取提现钱包额度
func (lukWalletLogApi *LukWalletLogApi) GetWallet(c *gin.Context) {
	//获取合约地址
	tokens := map[string]string{
		"wss":       global.GVA_CONFIG.Bsc.Wss,
		"lukToken":  global.GVA_CONFIG.Bsc.Luk,
		"usdtToken": global.GVA_CONFIG.Bsc.Usdt,
		"nftToken":  global.GVA_CONFIG.Bsc.Tiger,
	}
	walletDatas := lukWalletLogService.GetWallet()
	response.OkWithDetailed(map[string]interface{}{
		"walletDatas": walletDatas,
		"tokens":      tokens,
	}, "获取成功", c)
	return
}

//获取提现钱包额度
func (lukWalletLogApi *LukWalletLogApi) GetWalletConfigure(c *gin.Context) {
	//获取合约地址
	tokens := map[string]string{
		"wss":       global.GVA_CONFIG.Bsc.Wss,
		"lukToken":  global.GVA_CONFIG.Bsc.Luk,
		"usdtToken": global.GVA_CONFIG.Bsc.Usdt,
		"nftToken":  global.GVA_CONFIG.Bsc.Tiger,
	}
	response.OkWithDetailed(tokens, "获取成功", c)
	return
}
