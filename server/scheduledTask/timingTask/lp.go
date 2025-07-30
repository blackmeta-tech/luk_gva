package timingtask

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/shopspring/decimal"
	"time"
)

type LpTask struct {
}

func (l *LpTask) LpRebate(logics BonusLogic) BonusLogic {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("LP返佣同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	//查看当日是否分红过
	d := time.Now().Format(config.FORMAT_DATE_CST)
	log, _ := (luk.LukDividendsLog{}).QueryByTime(d, emun.DividendTypeLp)
	if log.ID > 0 {
		global.GVA_LOG.Error("今日已存在数据，不可再分！")
		return logics
	}
	var userAddress []luk.LukUserAddress
	err := global.GVA_DB.Model(&luk.LukUserAddress{}).Where("lp > 0 and no_lp = 0 and is_charge = 0").Find(&userAddress).Error
	if err != nil {
		return logics
	}
	pool, _ := lukServiceGroup.LukPoolBasicService.GetData()
	if pool.Usdt.Equal(decimal.Zero) {
		global.GVA_LOG.Error("USDT池子为0!")
		return logics
	}
	LU := pool.Luk.Div(pool.Usdt)
	//获取LP对应U的价值
	var price decimal.Decimal
	valueLp, _ := TokenServiceGroup.TotalSupply()
	_, valueU, _ := TokenServiceGroup.GetReserves()
	if valueLp.GreaterThan(decimal.Zero) {
		price = valueU.Div(valueLp)
	}
	//释放百分比
	lpRebate := global.GVA_LUK_CONFIG.LpRebate.Div(decimal.NewFromInt(100))
	for _, item := range userAddress {
		//统计改地址分红的U的数量，达到一千U即可停止
		countUsdt := lukServiceGroup.LukRevenueDetailedService.QueryByAddress(item.Address)
		if countUsdt.GreaterThanOrEqual(global.GVA_LUK_CONFIG.LpRebateHighest) {
			item.NoLp = 1
			logics.tx.Save(&item)
			continue
		}
		detailed := luk.LukRevenueDetailed{
			Address:   item.Address,
			Category:  emun.DividendTypeLp,
			TokenType: emun.LUK,
			Type:      emun.RevenueTypeLpRebate,
		}
		lpUsdt := item.Lp.Mul(price)
		if global.GVA_LUK_CONFIG.LpRebateHighest.LessThanOrEqual(lpUsdt) {
			detailed.Relatively = global.GVA_LUK_CONFIG.LpRebateHighest
		} else {
			detailed.Relatively = lpUsdt
		}
		relatively := global.GVA_LUK_CONFIG.LpRebateHighest.Sub(countUsdt).Div(lpRebate)
		if relatively.LessThan(detailed.Relatively) {
			detailed.Relatively = relatively
		}
		detailed.Amount = detailed.Relatively.Mul(lpRebate).Mul(LU)
		logics.lukRevenueDetailed = append(logics.lukRevenueDetailed, detailed)
	}
	if len(logics.lukRevenueDetailed) > 0 {
		dividendsLog := luk.LukDividendsLog{
			Date:     time.Now().Format(config.FORMAT_DATE_CST),
			Category: emun.DividendTypeLp,
		}
		logics.logs = append(logics.logs, dividendsLog)
	}
	return logics
}
