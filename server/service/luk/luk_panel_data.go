package luk

import (
	"fmt"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"sync"
)

type LukPanelDataService struct {
}

//面板数据统计
func (l LukPanelDataService) QueryData() (list interface{}) {
	pondData := response.PondData{}
	precipitationData := response.PrecipitationData{}
	dividendsData := response.DividendsData{}
	heritageData := response.HeritageData{}
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		pondData = l.PondData()
	}()
	go func() {
		defer wg.Done()
		precipitationData = l.PrecipitationData()
	}()
	go func() {
		defer wg.Done()
		dividendsData = l.DividendsData()
	}()
	go func() {
		defer wg.Done()
		heritageData = l.HeritageData()
	}()
	otherData := response.OtherData{
		Wss:            global.GVA_CONFIG.Bsc.Wss,
		LukToken:       global.GVA_CONFIG.Bsc.Luk,
		UsdtToken:      global.GVA_CONFIG.Bsc.Usdt,
		NftToken:       global.GVA_CONFIG.Bsc.Tiger,
		LiquidityToken: global.GVA_CONFIG.Bsc.Liquidity,
	}
	query := QueryLukWithdrawHistory{}
	otherData.WithdrawLuk, otherData.WithdrawUsdt = query.GetCountAmount()
	wg.Wait()

	data := map[string]interface{}{
		"pond":          pondData,
		"precipitation": precipitationData,
		"dividends":     dividendsData,
		"other":         otherData,
		"heritage":      heritageData,
	}
	list = data
	return
}

//池子数据
func (LukPanelDataService) PondData() (data response.PondData) {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		pool, _ := (LukPoolBasicService{}).GetData()
		data.Luk = pool.Luk
		data.Usdt = pool.Usdt
	}()
	go func() {
		defer wg.Done()
		data.DestructionLuk = (&LukDestroyService{}).QuerySummary("")
	}()
	go func() {
		defer wg.Done()
		data.Address = (&LukUserService{}).CountAll()
	}()
	go func() {
		defer wg.Done()
		lukBlockScanRecord := (&LukBlockScanRecordService{}).GetRemoveLiquidity()
		for _, item := range lukBlockScanRecord {
			switch item.Type {
			case emun.MethodTypeSwap:
				data.SwapLuk = item.LukAmount
				data.SwapUsdt = item.UsdtAmount
				data.SwapFee = item.LukServiceFee
			case emun.MethodTypeSell:
				data.SellLuk = item.LukAmount
				data.SellUsdt = item.UsdtAmount
				data.SellFee = item.LukServiceFee
			case emun.MethodTypeAddLiquidity:
				data.AddLiquidity = item.LpAmount
			case emun.MethodTypeRemoveLiquidity:
				data.RemoveLiquidity = item.LpAmount
				data.RemoveLiquidityFee = item.LukServiceFee
			}
		}
	}()
	global.GVA_DB.Debug().Table((&luk.LukNft{}).TableName()).Where("address != ''").Count(&data.Nft)
	wg.Wait()
	return
}

//回流数据
func (LukPanelDataService) PrecipitationData() (data response.PrecipitationData) {
	reflow := []luk.LukReflow{}
	err := global.GVA_DB.Model(&luk.LukReflow{}).Select("type, sum(amount) amount").Group("type").Find(&reflow).Error
	if err == nil {
		for _, item := range reflow {
			switch item.Type {
			case emun.ReflowTypeAlliance:
				data.Alliance = item.Amount
			case emun.ReflowTypeStcPot:
				data.StcPot = item.Amount
			case emun.ReflowTypeStcNode:
				data.StcNode = item.Amount
			case emun.ReflowTypeOperate:
				data.Operate = item.Amount
			case emun.ReflowTypeNft:
				data.Nft = item.Amount
			case emun.ReflowTypeFunding:
				data.Funding = item.Amount
			case emun.ReflowTypeAllowance:
				data.Allowance = item.Amount
			}
		}
	}
	reflowLp := []luk.LukReflowLp{}
	err1 := global.GVA_DB.Model(&luk.LukReflowLp{}).Select("type, sum(amount) amount").Group("type").Find(&reflowLp).Error
	if err1 == nil {
		for _, item := range reflowLp {
			switch item.Type {
			case emun.ReflowLpTypePool:
				data.LpPool = item.Amount
			case emun.ReflowLpTypeOperate:
				data.LpOperate = item.Amount
			}
		}
	}
	return
}

//分红数据
func (LukPanelDataService) DividendsData() (data response.DividendsData) {
	revenueNft := []luk.LukRevenueNft{}
	err := global.GVA_DB.Model(&luk.LukRevenueNft{}).Select("type, sum(amount) amount").Group("type").Find(&revenueNft).Error
	if err == nil {
		for _, item := range revenueNft {
			switch item.Type {
			case emun.RevenueTypeRecommend:
				data.Recommend = item.Amount
			case emun.RevenueTypeNftSwap:
				data.NftSwap = item.Amount
			case emun.RevenueTypeNftSell:
				data.NftSell = item.Amount
			case emun.RevenueTypeNftRemoveLiquidity:
				data.NftRemoveLiquidity = item.Amount
			case emun.RevenueTypeNftWeighted:
				data.NftWeighted = item.Amount
			case emun.RevenueTypeNftValue:
				data.NftValue = item.Amount
			}
		}
	}
	revenue := []luk.LukRevenue{}
	err1 := global.GVA_DB.Debug().Model(&luk.LukRevenue{}).Select("type, sum(amount) amount").Group("type").Find(&revenue).Error
	fmt.Println("====", err1)
	if err1 == nil {
		for _, item := range revenue {
			switch item.Type {
			case emun.RevenueTypeCombo7:
				data.Combo7 = item.Amount
			case emun.RevenueTypeCombo15:
				data.Combo15 = item.Amount
			case emun.RevenueTypeCombo30:
				data.Combo30 = item.Amount
			case emun.RevenueTypeCombo60:
				data.Combo60 = item.Amount
			case emun.RevenueTypeComboRecommend:
				data.ComboRecommend = item.Amount
			case emun.RevenueTypeLinkage1:
				data.Linkage1 = item.Amount
			case emun.RevenueTypeLinkage2:
				data.Linkage2 = item.Amount
			case emun.RevenueTypeLinkage3:
				data.Linkage3 = item.Amount
			case emun.RevenueTypeLinkage4:
				data.Linkage4 = item.Amount
			case emun.RevenueTypeCommunity1:
				data.Community1 = item.Amount
			case emun.RevenueTypeCommunity2:
				data.Community2 = item.Amount
			case emun.RevenueTypeCommunity3:
				data.Community3 = item.Amount
			case emun.RevenueTypeCommunity4:
				data.Community4 = item.Amount
			case emun.RevenueTypeLpSwap:
				data.LpSwap = item.Amount
			case emun.RevenueTypeLpSell:
				data.LpSell = item.Amount
			case emun.RevenueTypeLpRebate:
				data.LpRebate = item.Amount
			}
		}
	}
	return
}

//遗产数据
func (LukPanelDataService) HeritageData() (data response.HeritageData) {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		heritageMetaverse := luk.LukHeritageMetaverse{}
		lukHeritageMetaverse := global.GVA_DB.Model(&luk.LukHeritageMetaverse{}).Select("sum(balance) balance, sum(remaining) as remaining")
		lukHeritageMetaverse.Count(&data.MetaverseNum)
		err := lukHeritageMetaverse.First(&heritageMetaverse).Error
		if err == nil {
			data.MetaverseBalance = heritageMetaverse.Balance
			data.MetaverseRemaining = heritageMetaverse.Remaining
		}
	}()
	go func() {
		defer wg.Done()
		heritageMetaverseOld := luk.LukHeritageMetaverseOld{}
		lukHeritageMetaverseOld := global.GVA_DB.Model(&luk.LukHeritageMetaverseOld{}).Select("sum(balance) balance, sum(remaining) as remaining")
		lukHeritageMetaverseOld.Count(&data.MetaverseOldNum)
		err := lukHeritageMetaverseOld.First(&heritageMetaverseOld).Error
		if err == nil {
			data.MetaverseOldBalance = heritageMetaverseOld.Balance
			data.MetaverseOldRemaining = heritageMetaverseOld.Remaining
		}
	}()
	go func() {
		defer wg.Done()
		heritageLp := luk.LukHeritageLp{}
		lukHeritageLp := global.GVA_DB.Model(&luk.LukHeritageLp{}).Select("sum(balance) balance, sum(remaining) as remaining")
		lukHeritageLp.Count(&data.LpNum)
		err := lukHeritageLp.First(&heritageLp).Error
		if err == nil {
			data.LpBalance = heritageLp.Balance
			data.LpRemaining = heritageLp.Remaining
		}
	}()
	go func() {
		defer wg.Done()
		revenue := []luk.LukRevenueHeritage{}
		err := global.GVA_DB.Model(&luk.LukRevenueHeritage{}).Select("type, sum(luk) luk, sum(usdt) usdt").Group("type").Find(&revenue).Error
		if err == nil {
			for _, item := range revenue {
				switch item.Type {
				case emun.HeritageTypeMetaverse:
					data.Metaverse = item.Luk
					data.MetaverseUsdt = item.Usdt
				case emun.HeritageTypeMetaverseOld:
					data.MetaverseOld = item.Luk
					data.MetaverseOldUsdt = item.Usdt
				case emun.HeritageTypeLp:
					data.Lp = item.Luk
					data.LpUsdt = item.Usdt
				}
			}
		}
	}()
	wg.Wait()
	return
}
