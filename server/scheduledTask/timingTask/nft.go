package timingtask

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"sync"
	"time"
)

type NftTask struct {
}

var NftBonusComboTaskMt = new(sync.RWMutex)

func (n *NftTask) SyncNftAddress() {
	global.GVA_LOG.Info("开始同步NFT数据")
	t := time.Now().Format(config.FORMAT_DATETIME_CST)
	defer func() {
		global.GVA_LOG.Info("同步开始时间：" + t + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	nfts := []luk.LukNft{}
	err := global.GVA_DB.Find(&nfts).Error
	if err != nil {
		return
	}

	nftTokenService := service.ServiceGroupApp.TokenServiceGroup.NftTokenService
	for _, item := range nfts {
		address, err := nftTokenService.OwnerOf(item.ID)
		if err != nil {
			continue
		}
		if strings.ToLower(item.Address) != strings.ToLower(address.String()) {
			//插入一条转移记录
			transfer := luk.LukNftTransfer{
				From:  item.Address,
				To:    strings.ToLower(address.String()),
				NftId: item.ID,
			}
			err = global.GVA_DB.Create(&transfer).Error
			if err == nil {
				//更新数据库的
				item.Address = strings.ToLower(address.String())
				global.GVA_DB.Save(&item)
			}
		}

	}
}

func (n *NftTask) NftBonus(logics BonusLogic) BonusLogic {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("NFT分红同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	//查看当日是否分红过
	d := time.Now().Format(config.FORMAT_DATE_CST)
	log, _ := (luk.LukDividendsLog{}).QueryByTime(d, emun.DividendTypeNft)
	if log.ID > 0 {
		global.GVA_LOG.Error("今日已存在数据，不可再分！")
		return logics
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	var detailedWeighte, detailedValue []luk.LukRevenueDetailedNft
	var nftLogMap []luk.LukDividendsNftLog
	go func() {
		defer wg.Done()
		detailedWeighte = n.NftWeighted()
	}()
	go func() {
		defer wg.Done()
		detailedValue, nftLogMap = n.NftValue()
	}()
	wg.Wait()

	if len(detailedWeighte) > 0 {
		logics.lukRevenueDetailedNft = append(logics.lukRevenueDetailedNft, detailedWeighte...)
	}
	if len(detailedValue) > 0 {
		logics.lukRevenueDetailedNft = append(logics.lukRevenueDetailedNft, detailedValue...)
		logics.nftLogs = append(logics.nftLogs, nftLogMap...)
	}
	if len(logics.lukRevenueDetailedNft) > 0 {
		dividendsLog := luk.LukDividendsLog{
			Date:     time.Now().Format(config.FORMAT_DATE_CST),
			Category: emun.DividendTypeNft,
		}
		logics.logs = append(logics.logs, dividendsLog)
	}
	return logics
}

//nft加权分红
func (n *NftTask) NftWeighted() (detailedMap []luk.LukRevenueDetailedNft) {
	//查询昨日购买套餐
	last := time.Now().AddDate(0, 0, -1).Format(config.FORMAT_DATE_CST)
	amountReflow, _ := lukServiceGroup.LukReflowService.GetLukReflowBydate(last, emun.ReflowTypeNft)
	if amountReflow.IsZero() {
		global.GVA_LOG.Error("昨日没有购买回流数据")
		return
	}

	//统计当前nft名单
	nftList := lukServiceGroup.LukNftService.QueryNftByNum()
	if len(nftList) == 0 {
		global.GVA_LOG.Error("当前nft名单没有数据")
		return
	}

	for address, pro := range nftList {
		detailed := luk.LukRevenueDetailedNft{
			Address:   address,
			Amount:    amountReflow.Mul(pro),
			TokenType: emun.LUK,
			Type:      emun.RevenueTypeNftWeighted,
		}
		detailedMap = append(detailedMap, detailed)
	}
	return
}

//nft1500U分红
func (n *NftTask) NftValue() (detailedMap []luk.LukRevenueDetailedNft, nftLogMap []luk.LukDividendsNftLog) {
	nfts := []luk.LukNft{}
	err := global.GVA_DB.Debug().Where("blacklist = 0 and is_dividend = 0 and address != ''").Find(&nfts).Error
	if err != nil {
		return
	}
	//算出1500平均90天的U量，再根据当前币价获取得到的luk
	pool, _ := lukServiceGroup.LukPoolBasicService.GetData()
	usdtAmount := decimal.NewFromInt(1500).Div(decimal.NewFromFloat(90))
	lukAmount := pool.Luk.Div(pool.Usdt).Mul(usdtAmount)
	d := time.Now().Format(config.FORMAT_DATE_CST)
	nftlog := luk.LukDividendsNftLog{}
	for _, item := range nfts {
		//判断分红天数是否已达到90天，到的话就可以不分了
		days := nftlog.QueryByNftId(item.ID)
		if days >= 90 {
			item.IsDividend = true
			global.GVA_DB.Save(&item)
			continue
		}
		detailed := luk.LukRevenueDetailedNft{
			Address:    item.Address,
			Amount:     lukAmount,
			TokenType:  emun.LUK,
			Type:       emun.RevenueTypeNftValue,
			Relatively: usdtAmount,
			Source:     strconv.Itoa(int(item.ID)),
		}
		detailedMap = append(detailedMap, detailed)

		logInfo := luk.LukDividendsNftLog{
			NftId: item.ID,
			Date:  d,
		}
		nftLogMap = append(nftLogMap, logInfo)
	}
	return
}
