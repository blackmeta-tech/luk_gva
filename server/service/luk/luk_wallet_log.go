package luk

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun"
	enumLuk "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	lukRes "github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	util "github.com/flipped-aurora/gin-vue-admin/server/model/utils"
	tokenService "github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"math/big"
	"net/url"
	"sort"
	"sync"
)

type LukWalletLogService struct {
}

//获取提现钱包额度
func (LukWalletLogService *LukWalletLogService) GetWallet() (walletDatas []lukRes.LukWalletList) {
	tokenService := &tokenService.BscTokenService{}
	var transfer token.Transfer

	walletDatas = make([]lukRes.LukWalletList, 0)
	wallets := global.GVA_CONFIG.Wallet
	cs := utils.StructToMap(wallets)

	enumMapInstance := enum.GetEnumMapInstance()
	data := enumMapInstance.GetData([]string{"wallet_type"})
	walletTypes := data["wallet_type"].AllMember()

	for k, v := range walletTypes {
		abbr := string(k.(enumLuk.WalletType))
		pk := abbr + "PrivateKey"
		if str, ok := cs[abbr].(string); ok {
			info := lukRes.LukWalletList{
				Abbr:    abbr,
				Name:    v,
				Address: str,
			}
			if strr, okk := cs[pk].(config.WalletEncryption); okk {
				info.PrivateKey = strr.String()
			}
			//查询钱包余额
			transfer.From = str
			var wg sync.WaitGroup
			wg.Add(3)
			go func() {
				defer wg.Done()
				transfer.Type = enumLuk.LUK
				info.Luk, _ = tokenService.GetBalanceOf(transfer)
			}()
			go func() {
				defer wg.Done()
				transfer.Type = enumLuk.USDT
				info.Usdt, _ = tokenService.GetBalanceOf(transfer)
			}()
			go func() {
				//获取BNB余额
				defer wg.Done()
				info.Bnb = LukWalletLogService.GetBnb(info.Address)
			}()
			wg.Wait()
			walletDatas = append(walletDatas, info)
		}
	}

	sort.SliceStable(walletDatas, func(i, j int) bool {
		return walletDatas[i].Name < walletDatas[j].Name
	})
	return
}

//获取BNB的值
func (LukWalletLogService *LukWalletLogService) GetBnb(address string) (bnb decimal.Decimal) {
	params := url.Values{}
	params.Add("apikey", global.GVA_CONFIG.Bsc.Apikey)
	params.Add("module", "account")
	params.Add("action", "balance")
	params.Add("address", address)
	resp, err := utils.HttpRequest(global.GVA_CONFIG.Bsc.BscAddr, "/api", params)
	if err != nil || resp.StatusCode != 200 {
		global.GVA_LOG.Sugar().Errorf("获取币安网数据失败, ststusCode:%d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	res := lukRes.BnbRes{}
	utils.Unmarshal(resp.Body, &res)
	n := new(big.Int)
	n, ok := n.SetString(res.Result, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	bnb = util.ToDecimal(n)
	return
}
