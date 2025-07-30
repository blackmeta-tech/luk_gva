package timingtask

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
	"time"
)

type HeritageTask struct {
}

var HeritageComboTaskMt = new(sync.RWMutex)

func (h *HeritageTask) Heritage() {
	HeritageComboTaskMt.Lock()
	defer HeritageComboTaskMt.Unlock()
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("遗产同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	pool, _ := lukServiceGroup.LukPoolBasicService.GetData()
	if pool.Usdt.Equal(decimal.Zero) {
		global.GVA_LOG.Error("USDT池子为0!")
		return
	}
	logic := HeritageLogic{
		LU: pool.Luk.Div(pool.Usdt),
	}
	logic.SyncMetaverse()

	//查看当日是否分红过
	d := time.Now().Format(config.FORMAT_DATE_CST)
	log, _ := (luk.LukHeritageLog{}).QueryByTime(d)
	if log.ID > 0 {
		global.GVA_LOG.Error("今日已存在数据，不可再分！")
		return
	}
	db := global.GVA_DB
	tx := db.Begin()
	logic.tx = tx
	logic.BonusMetaverse()
	logic.BonusMetaverseOld()
	logic.BonusLp()
	logic.BatchUpdate()
}

type HeritageLogic struct {
	tx        *gorm.DB
	LU        decimal.Decimal //luk除以usdt价格
	detaileds []luk.LukRevenueDetailedHeritage
}

//同步新元宇宙银行购买的套餐以及LP的数据
func (l *HeritageLogic) SyncMetaverse() {
	metaverses := []luk.LukHeritageMetaverse{}
	err := global.GVA_DB.Debug().Find(&metaverses).Error
	if err != nil {
		global.GVA_LOG.Error("查找记录失败!", zap.Error(err))
		return
	}
	//获取LP对应U的价值
	var price decimal.Decimal
	valueLp, _ := TokenServiceGroup.TotalSupply()
	_, valueU, _ := TokenServiceGroup.GetReserves()
	if valueLp.GreaterThan(decimal.Zero) {
		price = valueU.Div(valueLp)
	}
	for _, item := range metaverses {
		item.ComboUsdt = lukComboBuyService.GetPriceByAddress(strings.ToLower(item.Address))
		//获取LP的量
		ua, _ := lukServiceGroup.LukUserAddressService.GetOne(strings.ToLower(item.Address))
		item.LpUsdt = ua.Lp.Mul(price)
		item.UpdatedAt = time.Now()
		global.GVA_DB.Save(&item)
	}
}

//新元宇宙分红
func (l *HeritageLogic) BonusMetaverse() {
	metaverse := []luk.LukHeritageMetaverse{}
	err := global.GVA_DB.Debug().Where("is_dividend = 0").Find(&metaverse).Error
	if err != nil {
		return
	}
	_type := emun.HeritageTypeMetaverse
	metaverse1 := global.GVA_LUK_CONFIG.RatioMetaverse1.Div(decimal.NewFromInt(100)) //固定释放
	metaverse2 := global.GVA_LUK_CONFIG.RatioMetaverse2.Div(decimal.NewFromInt(100))
	metaverse3 := global.GVA_LUK_CONFIG.RatioMetaverse3.Div(decimal.NewFromInt(100))
	for _, item := range metaverse {
		if item.Remaining.IsZero() {
			item.IsDividend = true
			global.GVA_DB.Save(&item)
			continue
		}
		detailed := luk.LukRevenueDetailedHeritage{
			Address: item.Address,
			Type:    _type,
		}
		//都需要判断剩余的额度会不会小于释放的量，小于的话直接用剩余的额度释放
		//模式1 固定释放
		number := item.Balance.Mul(metaverse1)
		if item.Remaining.LessThanOrEqual(number) {
			detailed.Usdt = item.Remaining
			detailed.Luk = detailed.Usdt.Mul(l.LU)
			detailed.Percentage = metaverse1
			detailed.Remark = "固定"
			l.detaileds = append(l.detaileds, detailed)
			item.Remaining = decimal.Zero
			item.IsDividend = true
			global.GVA_DB.Save(&item)
			continue
		} else {
			item.Remaining = item.Remaining.Sub(number)
			detailed.Usdt = number
			detailed.Luk = detailed.Usdt.Mul(l.LU)
			detailed.Percentage = metaverse1
			detailed.Remark = "固定"
			l.detaileds = append(l.detaileds, detailed)
		}

		if item.ComboUsdt.GreaterThan(decimal.Zero) {
			//当前有购买套餐，则可启动释放模式二 套餐
			//需要对比套餐额度 跟自己的额度对比，哪个小用哪个比例
			if item.Balance.GreaterThan(item.ComboUsdt) {
				number = item.ComboUsdt.Mul(metaverse2)
			} else {
				number = item.Balance.Mul(metaverse2)
			}
			if item.Remaining.LessThanOrEqual(number) {
				detailed.Usdt = item.Remaining
				detailed.Luk = detailed.Usdt.Mul(l.LU)
				detailed.Percentage = metaverse2
				detailed.Remark = "套餐"
				l.detaileds = append(l.detaileds, detailed)
				item.Remaining = decimal.Zero
				item.IsDividend = true
				global.GVA_DB.Save(&item)
				continue
			} else {
				item.Remaining = item.Remaining.Sub(number)
				detailed.Usdt = number
				detailed.Luk = detailed.Usdt.Mul(l.LU)
				detailed.Percentage = metaverse2
				detailed.Remark = "套餐"
				l.detaileds = append(l.detaileds, detailed)
			}
		}

		if item.LpUsdt.GreaterThan(decimal.Zero) {
			//当前有Lp额度，则可启动释放模式三 LP
			if item.Balance.GreaterThan(item.LpUsdt) {
				number = item.LpUsdt.Mul(metaverse3)
			} else {
				number = item.Balance.Mul(metaverse3)
			}
			if item.Remaining.LessThanOrEqual(number) {
				detailed.Usdt = item.Remaining
				detailed.Luk = detailed.Usdt.Mul(l.LU)
				detailed.Percentage = metaverse3
				detailed.Remark = "LP"
				item.Remaining = decimal.Zero
				item.IsDividend = true
				global.GVA_DB.Save(&item)
			} else {
				item.Remaining = item.Remaining.Sub(number)
				detailed.Usdt = number
				detailed.Luk = detailed.Usdt.Mul(l.LU)
				detailed.Percentage = metaverse3
				detailed.Remark = "LP"
			}
			if detailed.Luk.GreaterThan(decimal.Zero) {
				l.detaileds = append(l.detaileds, detailed)
			}
		}
		item.UpdatedAt = time.Now()
		global.GVA_DB.Save(&item)
	}
}

//老元宇宙分红
func (l *HeritageLogic) BonusMetaverseOld() {
	metaverse := []luk.LukHeritageMetaverseOld{}
	err := global.GVA_DB.Debug().Where("is_dividend = 0").Find(&metaverse).Error
	if err != nil {
		return
	}
	_type := emun.HeritageTypeMetaverseOld
	bonus := global.GVA_LUK_CONFIG.RatioMetaverseOld.Div(decimal.NewFromInt(100)) //固定释放
	for _, item := range metaverse {
		if item.Remaining.IsZero() {
			item.IsDividend = true
			global.GVA_DB.Save(&item)
			continue
		}
		detailed := luk.LukRevenueDetailedHeritage{
			Address: item.Address,
			Type:    _type,
		}
		//固定释放
		number := item.Balance.Mul(bonus)
		if item.Remaining.LessThanOrEqual(number) {
			detailed.Usdt = item.Remaining
			item.Remaining = decimal.Zero
			item.IsDividend = true
		} else {
			item.Remaining = item.Remaining.Sub(number)
			detailed.Usdt = number
		}
		detailed.Luk = detailed.Usdt.Mul(l.LU)
		detailed.Percentage = bonus
		fmt.Println("====detailed", detailed)
		l.detaileds = append(l.detaileds, detailed)
		item.UpdatedAt = time.Now()
		global.GVA_DB.Save(&item)
	}
}

//LP矿机分红
func (l *HeritageLogic) BonusLp() {
	metaverse := []luk.LukHeritageLp{}
	err := global.GVA_DB.Debug().Where("is_dividend = 0").Find(&metaverse).Error
	if err != nil {
		return
	}
	_type := emun.HeritageTypeLp
	bonus := global.GVA_LUK_CONFIG.RatioLp.Div(decimal.NewFromInt(100)) //固定释放
	for _, item := range metaverse {
		if item.Remaining.IsZero() {
			item.IsDividend = true
			global.GVA_DB.Save(&item)
			continue
		}
		detailed := luk.LukRevenueDetailedHeritage{
			Address: item.Address,
			Type:    _type,
		}
		//固定释放
		number := item.Balance.Mul(bonus)
		if item.Remaining.LessThanOrEqual(number) {
			detailed.Usdt = item.Remaining
			item.Remaining = decimal.Zero
			item.IsDividend = true
		} else {
			item.Remaining = item.Remaining.Sub(number)
			detailed.Usdt = number
		}
		detailed.Luk = detailed.Usdt.Mul(l.LU)
		detailed.Percentage = bonus
		fmt.Println("====detailed", detailed)
		l.detaileds = append(l.detaileds, detailed)
		item.UpdatedAt = time.Now()
		global.GVA_DB.Save(&item)
	}
}

//批量写入处理
func (b *HeritageLogic) BatchUpdate() {
	var err error
	tx := b.tx
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	userWallets := map[string]luk.LukUser{}
	if len(b.detaileds) > 0 {
		revenues := map[string]luk.LukRevenueHeritage{}
		for _, itme := range b.detaileds {
			key := itme.Address + strconv.Itoa(int(itme.Type))
			if info, ok := revenues[key]; ok {
				info.Luk = info.Luk.Add(itme.Luk)
				info.Usdt = info.Usdt.Add(itme.Usdt)
				revenues[key] = info
			} else {
				info = luk.LukRevenueHeritage{
					Address: itme.Address,
					Type:    itme.Type,
					Luk:     itme.Luk,
					Usdt:    itme.Usdt,
				}
				revenues[key] = info
			}

			if wallet, ok := userWallets[itme.Address]; ok {
				wallet.Luk = wallet.Luk.Add(itme.Luk)
				userWallets[itme.Address] = wallet
			} else {
				wallet = luk.LukUser{}
				wallet.Address = itme.Address
				wallet.Luk = itme.Luk
				userWallets[itme.Address] = wallet
			}
		}

		err = lukServiceGroup.LukRevenueHeritageService.UpdateLukRevenueHeritageBatch(tx, revenues)
		if err != nil {
			return
		}

		err = tx.CreateInBatches(b.detaileds, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
		log := luk.LukHeritageLog{
			Date: time.Now().Format(config.FORMAT_DATE_CST),
		}
		err = tx.Create(&log).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}

	if len(userWallets) > 0 {
		err = lukServiceGroup.LukUserService.UpdateBatch(tx, userWallets)
		if err != nil {
			return
		}
	}
}
