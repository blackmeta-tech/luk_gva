package timingtask

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"sync"
	"time"
)

type UserTask struct {
}

var SyncComboNumComboTaskMt = new(sync.RWMutex)

//同步套餐购买交易量
func (u *UserTask) SyncComboNum() {
	SyncComboNumComboTaskMt.Lock()
	defer SyncComboNumComboTaskMt.Unlock()
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("套餐购买同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	u.SyncPerformance()
	u.SyncArea()
	u.DividendsDay() //分红处理
}

//同步每个用户旗下【含自己】20代购买业绩
func (u *UserTask) SyncPerformance() {
	var lukUsers []luk.LukUser
	err := global.GVA_DB.Model(&luk.LukUser{}).Find(&lukUsers).Error
	if err != nil {
		return
	}
	ch := make(chan luk.LukUser)
	go func() {
		defer close(ch)
		for _, item := range lukUsers {
			ch <- item
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range ch {
				//20代
				addresss := lukUserService.GetTreeAddressByGeneration(item.Address, 20)
				addresss = append(addresss, item.Address)
				item.Performance = lukComboBuyService.QueryPerformanceByaddresss(addresss)
				item.PerformanceTime = time.Now()
				global.GVA_DB.Save(item)
			}
		}()
	}
	wg.Wait()
}

//同步旗下大小区
func (u *UserTask) SyncArea() {
	var lukUsers []luk.LukUser
	err := global.GVA_DB.Model(&luk.LukUser{}).Find(&lukUsers).Error
	if err != nil {
		return
	}
	linkage, _ := lukServiceGroup.LukPartnerService.QueryList(emun.UserTypeLinkage, "desc")     //联盟配置
	community, _ := lukServiceGroup.LukPartnerService.QueryList(emun.UserTypeCommunity, "desc") //社区配置

	ch := make(chan luk.LukUser)
	go func() {
		defer close(ch)
		for _, item := range lukUsers {
			ch <- item
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range ch {
				//求出大区以及小区交易量
				areaMax, areaMin := lukUserService.QueryArea(item.Address)
				item.AreaMax = decimal.NewFromFloat(areaMax)
				item.AreaMin = decimal.NewFromFloat(areaMin)
				//判断社区的交易量
				item.CommunityLevel = emun.PartnerTypeNone
				item.LinkageLevel = emun.PartnerTypeNone
				for _, c := range community {
					if item.AreaMax.GreaterThanOrEqual(c.AreaMax) && item.AreaMin.GreaterThanOrEqual(c.AreaMin) {
						item.CommunityLevel = c.Level
						break
					}
				}
				//联盟白名单的话进行考核
				if item.IsLinkage {
					for _, c := range linkage {
						if item.AreaMax.GreaterThanOrEqual(c.AreaMax) && item.AreaMin.GreaterThanOrEqual(c.AreaMin) {
							item.LinkageLevel = c.Level
							break
						}
					}
				}
				item.UpdatedAt = time.Now()
				global.GVA_DB.Save(item)
			}
		}()
	}
	wg.Wait()
}

func (u *UserTask) DividendsDay() {
	db := global.GVA_DB
	tx := db.Begin()
	//分红逻辑
	logics := BonusLogic{
		tx: tx,
	}
	logics = u.Dividends(logics)           //社区联盟分红
	logics = (&NftTask{}).NftBonus(logics) //NFT套餐加权分红
	logics = (&LpTask{}).LpRebate(logics)  //LP返佣分红
	logics.BatchUpdate()

}
func (u *UserTask) Dividends(logics BonusLogic) BonusLogic {
	//进行分红
	wg := sync.WaitGroup{}
	wg.Add(2)
	var detailedLinkage, detailedCommunity []luk.LukRevenueDetailed
	go func() {
		defer wg.Done()
		detailedCommunity = u.DividendsCommunity()
	}()
	go func() {
		defer wg.Done()
		detailedLinkage = u.DividendsLinkage()
	}()
	wg.Wait()

	if len(detailedLinkage) > 0 {
		log := luk.LukDividendsLog{
			Date:     time.Now().Format(config.FORMAT_DATE_CST),
			Category: emun.DividendTypeLinkage,
		}
		logics.logs = append(logics.logs, log)
		logics.lukRevenueDetailed = append(logics.lukRevenueDetailed, detailedLinkage...)
	}
	if len(detailedCommunity) > 0 {
		log := luk.LukDividendsLog{
			Date:     time.Now().Format(config.FORMAT_DATE_CST),
			Category: emun.DividendTypeCommunity,
		}
		logics.logs = append(logics.logs, log)
		logics.lukRevenueDetailed = append(logics.lukRevenueDetailed, detailedCommunity...)
	}

	return logics
}

//联盟分红
func (u *UserTask) DividendsLinkage() (detailedMap []luk.LukRevenueDetailed) {
	//查看当日是否分红过
	d := time.Now().Format(config.FORMAT_DATE_CST)
	log, _ := (luk.LukDividendsLog{}).QueryByTime(d, emun.DividendTypeLinkage)
	if log.ID > 0 {
		global.GVA_LOG.Error("今日已存在数据，不可再分！")
		return
	}
	//查询昨日购买套餐，以及当前白名单用户人员
	last := time.Now().AddDate(0, 0, -1).Format(config.FORMAT_DATE_CST)
	amountReflow, _ := lukServiceGroup.LukReflowService.GetLukReflowBydate(last, emun.ReflowTypeAlliance)
	if amountReflow.IsZero() {
		global.GVA_LOG.Error("昨日没有购买回流数据")
		return
	}
	//统计当前联盟白名单
	linkages := lukUserService.QueryIsLinkage()
	if len(linkages) == 0 {
		global.GVA_LOG.Error("联盟白名单没有考核通过的数据")
		return
	}
	var lukUsers []luk.LukUser
	err := global.GVA_DB.Model(&luk.LukUser{}).Where("is_linkage = 1 and linkage_level > 0").Find(&lukUsers).Error
	if err != nil {
		return
	}
	fmt.Println("=====", amountReflow)
	fmt.Println("=====", linkages)
	for _, item := range lukUsers {
		detailed := luk.LukRevenueDetailed{
			Address:   item.Address,
			Category:  emun.DividendTypeLinkage,
			TokenType: emun.LUK,
		}
		detailed.Amount = linkages[item.LinkageLevel].Mul(amountReflow)
		switch item.LinkageLevel {
		case emun.PartnerTypeEarly:
			detailed.Type = emun.RevenueTypeLinkage1
		case emun.PartnerTypeMiddle:
			detailed.Type = emun.RevenueTypeLinkage2
		case emun.PartnerTypeHigh:
			detailed.Type = emun.RevenueTypeLinkage3
		case emun.PartnerTypeOvertake:
			detailed.Type = emun.RevenueTypeLinkage4
		}
		detailedMap = append(detailedMap, detailed)
	}
	return
}

//社区分红
func (u *UserTask) DividendsCommunity() (detailedMap []luk.LukRevenueDetailed) {
	//查看当日是否分红过
	d := time.Now().Format(config.FORMAT_DATE_CST)
	log, _ := (luk.LukDividendsLog{}).QueryByTime(d, emun.DividendTypeCommunity)
	if log.ID > 0 {
		global.GVA_LOG.Error("今日已存在数据，不可再分！")
		return
	}
	var lukUsers []luk.LukUser
	err := global.GVA_DB.Model(&luk.LukUser{}).Where("community_level > 0").Find(&lukUsers).Error
	if err != nil {
		return
	}
	pool, _ := lukServiceGroup.LukPoolBasicService.GetData()
	if pool.Usdt.Equal(decimal.Zero) {
		global.GVA_LOG.Error("USDT池子为0!", zap.Error(err))
		return
	}
	lu := pool.Luk.Div(pool.Usdt)
	detailedMap = make([]luk.LukRevenueDetailed, 0)
	earlyAmount := decimal.NewFromInt(50).Mul(lu)
	middleAmount := decimal.NewFromInt(250).Mul(lu)
	highAmount := decimal.NewFromInt(500).Mul(lu)
	overtakeAmount := decimal.NewFromInt(1000).Mul(lu)
	for _, item := range lukUsers {
		detailed := luk.LukRevenueDetailed{
			Address:   item.Address,
			Category:  emun.DividendTypeCommunity,
			TokenType: emun.LUK,
		}
		switch item.CommunityLevel {
		case emun.PartnerTypeEarly:
			detailed.Amount = earlyAmount
			detailed.Type = emun.RevenueTypeCommunity1
		case emun.PartnerTypeMiddle:
			detailed.Amount = middleAmount
			detailed.Type = emun.RevenueTypeCommunity2
		case emun.PartnerTypeHigh:
			detailed.Amount = highAmount
			detailed.Type = emun.RevenueTypeCommunity3
		case emun.PartnerTypeOvertake:
			detailed.Amount = overtakeAmount
			detailed.Type = emun.RevenueTypeCommunity4
		}
		detailedMap = append(detailedMap, detailed)
	}
	return
}
