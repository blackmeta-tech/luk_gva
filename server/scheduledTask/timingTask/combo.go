package timingtask

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	tokenService "github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ComboTask struct {
}

var BonusComboTaskMt = new(sync.RWMutex)

func (c *ComboTask) AllBonus() {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("分红同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	BonusComboTaskMt.Lock()
	defer BonusComboTaskMt.Unlock()
	pool, err := lukServiceGroup.LukPoolBasicService.GetData()
	if err != nil {
		global.GVA_LOG.Error("获取底池数据失败!", zap.Error(err))
		return
	}
	if pool.Usdt.Equal(decimal.Zero) {
		global.GVA_LOG.Error("USDT池子为0!", zap.Error(err))
		return
	}
	db := global.GVA_DB
	tx := db.Begin()
	//分红逻辑
	logics := BonusLogic{
		tx: tx,
		LU: pool.Luk.Div(pool.Usdt),
	}
	logics.DestroyReflow() //销毁与回流
	logics.Bonus()         //套餐质押分红
	logics.LpBonus()       //LP分红
	logics.BatchUpdate()
	//进行销毁
	go func() {
		_ = service.ServiceGroupApp.LukServiceGroup.LukDestroyService.Destroy()
	}()
	logics.repeat() //复购操作
}

//套餐质押分红
func (c *BonusLogic) Bonus() {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("套餐质押分红同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	buys := []luk.LukComboBuy{}
	err := global.GVA_DB.Debug().Where("maturity_at < ? and dividend = 0 and destroy = 1", time.Now()).Find(&buys).Error
	if err != nil {
		global.GVA_LOG.Error("查找记录失败!", zap.Error(err))
		return
	}

	for _, item := range buys {
		//算出套餐购买所得到的利率
		rate := item.Rate.Div(decimal.NewFromInt(100)).Mul(decimal.NewFromInt(int64(item.Days)))
		interest := item.PriceUsdt.Mul(c.LU).Mul(rate) //得到的利息 用购买的USDT乘汇率得到利息
		relatively := item.PriceUsdt.Mul(rate)         //得到的利息 Usdt
		c.data = item
		c.interest = interest
		c.relatively = relatively
		c.ownBonus()
		c.recommend()
		item.Status = -1
		item.Dividend = 1
		item.DividendPrice = c.LU
		_ = lukComboBuyService.Update(c.tx, item)
		//需要复购的加进复购逻辑
		if item.Repeat == 1 {
			c.repeats = append(c.repeats, item)
		}
	}
}

type BonusLogic struct {
	data                  luk.LukComboBuy
	interest              decimal.Decimal             //得到的利息
	relatively            decimal.Decimal             //得到的利息USDT
	lukRevenueDetailed    []luk.LukRevenueDetailed    //分红记录
	lukRevenueDetailedNft []luk.LukRevenueDetailedNft //分红记录Nft
	repeats               []luk.LukComboBuy           //复购记录
	tx                    *gorm.DB
	LU                    decimal.Decimal          //luk除以usdt价格
	logs                  []luk.LukDividendsLog    //分红日期记录
	reflowLps             []luk.LukReflowLp        //lp回流记录
	nftLogs               []luk.LukDividendsNftLog //nft一千五90天分红记录
	destroys              []luk.LukDestroy         //销毁
	reflows               []luk.LukReflow          //回流
}

//套餐质押给自己的分红
func (b *BonusLogic) ownBonus() {
	detailed := luk.LukRevenueDetailed{
		Address:    b.data.Address,
		Amount:     b.interest,
		Category:   emun.DividendTypeCombo,
		TokenType:  emun.LUK,
		Remark:     "利息",
		Source:     strconv.Itoa(int(b.data.ID)),
		Relatively: b.relatively,
	}
	switch b.data.Days {
	case 7:
		detailed.Type = emun.RevenueTypeCombo7
	case 15:
		detailed.Type = emun.RevenueTypeCombo15
	case 30:
		detailed.Type = emun.RevenueTypeCombo30
	case 60:
		detailed.Type = emun.RevenueTypeCombo60
	}
	b.lukRevenueDetailed = append(b.lukRevenueDetailed, detailed)
	//归还本金
	detailed.Amount = b.data.PriceUsdt.Mul(b.LU)
	detailed.Relatively = b.data.PriceUsdt
	detailed.Remark = "本金"
	b.lukRevenueDetailed = append(b.lukRevenueDetailed, detailed)
	return
}

//套餐代数分红
func (b *BonusLogic) recommend() {
	//十代分红利率
	recommend1 := global.GVA_LUK_CONFIG.ComboRecommend1.Div(decimal.NewFromInt(100))
	recommend2 := global.GVA_LUK_CONFIG.ComboRecommend2.Div(decimal.NewFromInt(100))
	recommend6 := global.GVA_LUK_CONFIG.ComboRecommend6.Div(decimal.NewFromInt(100))
	rate := map[int]decimal.Decimal{
		1:  recommend1,
		2:  recommend2,
		3:  recommend2,
		4:  recommend2,
		5:  recommend2,
		6:  recommend6,
		7:  recommend6,
		8:  recommend6,
		9:  recommend6,
		10: recommend6,
	}
	var keys []int
	for k := range rate {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("=====", keys)
	address := b.data.Address
	for _, i := range keys {
		//查找父级
		user, _ := lukUserService.GetLukUserByAddress(b.tx, address)
		if user.PAddress == "" {
			continue
		}
		address = user.PAddress
		//父级要购买套餐
		if !lukComboBuyService.GetByAddressHash(address) {
			continue
		}
		//获取父级直推地址要大于等于代数
		subAddress := lukUserService.QueryDirectPush(address)
		if len(subAddress) < i {
			continue
		}
		//父级直推有效用户要大于等于
		buyCount := lukComboBuyService.GetByAddresss(subAddress)
		if i == 1 {
			buyCount = buyCount + 1
		}
		if int(buyCount) < i {
			continue
		}

		fmt.Println("=====", i)
		detailed := luk.LukRevenueDetailed{
			Address:   address,
			Amount:    b.interest.Mul(rate[i]),
			Category:  emun.DividendTypeComboRecommend,
			TokenType: emun.LUK,
			Type:      emun.RevenueTypeComboRecommend,
			Remark:    strconv.Itoa(i) + "代推荐分红",
			Source:    strconv.Itoa(int(b.data.ID)),
		}
		b.lukRevenueDetailed = append(b.lukRevenueDetailed, detailed)
	}
	return
}

//批量写入处理
func (b *BonusLogic) BatchUpdate() {
	var err error
	tx := b.tx
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	userWallets := map[string]luk.LukUser{}
	if len(b.lukRevenueDetailed) > 0 {
		revenues := map[string]luk.LukRevenue{}
		for _, itme := range b.lukRevenueDetailed {
			if itme.Remark != "本金" {
				key := itme.Address + strconv.Itoa(int(itme.Type)) + strconv.Itoa(int(itme.TokenType))
				if info, ok := revenues[key]; ok {
					info.Amount = info.Amount.Add(itme.Amount)
					revenues[key] = info
				} else {
					info = luk.LukRevenue{
						Address:   itme.Address,
						Category:  itme.Category,
						Type:      itme.Type,
						TokenType: itme.TokenType,
						Amount:    itme.Amount,
					}
					revenues[key] = info
				}
			}

			if wallet, ok := userWallets[itme.Address]; ok {
				switch itme.TokenType {
				case emun.LUK:
					wallet.Luk = wallet.Luk.Add(itme.Amount)
				case emun.USDT:
					wallet.Usdt = wallet.Usdt.Add(itme.Amount)
				}
				userWallets[itme.Address] = wallet
			} else {
				wallet = luk.LukUser{}
				wallet.Address = itme.Address
				switch itme.TokenType {
				case emun.LUK:
					wallet.Luk = itme.Amount
				case emun.USDT:
					wallet.Usdt = itme.Amount
				}
				userWallets[itme.Address] = wallet
			}
		}

		err = service.ServiceGroupApp.LukServiceGroup.LukRevenueService.UpdateLukRevenueBatch(tx, revenues)
		if err != nil {
			return
		}

		err = tx.CreateInBatches(b.lukRevenueDetailed, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
	if len(b.lukRevenueDetailedNft) > 0 {
		revenues := map[string]luk.LukRevenueNft{}

		for _, itme := range b.lukRevenueDetailedNft {
			key := itme.Address + strconv.Itoa(int(itme.Type)) + strconv.Itoa(int(itme.TokenType))
			if info, ok := revenues[key]; ok {
				info.Amount = info.Amount.Add(itme.Amount)
				revenues[key] = info
			} else {
				info = luk.LukRevenueNft{
					Address:   itme.Address,
					Type:      itme.Type,
					TokenType: itme.TokenType,
					Amount:    itme.Amount,
				}
				revenues[key] = info
			}

			if wallet, ok := userWallets[itme.Address]; ok {
				switch itme.TokenType {
				case emun.LUK:
					wallet.Luk = wallet.Luk.Add(itme.Amount)
				case emun.USDT:
					wallet.Usdt = wallet.Usdt.Add(itme.Amount)
				}
				userWallets[itme.Address] = wallet
			} else {
				wallet = luk.LukUser{}
				wallet.Address = itme.Address
				switch itme.TokenType {
				case emun.LUK:
					wallet.Luk = itme.Amount
				case emun.USDT:
					wallet.Usdt = itme.Amount
				}
				userWallets[itme.Address] = wallet
			}
		}

		err = service.ServiceGroupApp.LukServiceGroup.LukRevenueNftService.UpdateLukRevenueNftBatch(tx, revenues)
		if err != nil {
			return
		}

		err = tx.CreateInBatches(b.lukRevenueDetailedNft, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
	if len(userWallets) > 0 {
		err = service.ServiceGroupApp.LukServiceGroup.LukUserService.UpdateBatch(tx, userWallets)
		if err != nil {
			return
		}
	}
	if len(b.logs) > 0 {
		err = tx.CreateInBatches(b.logs, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
	if len(b.reflowLps) > 0 {
		err = tx.CreateInBatches(b.reflowLps, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
	if len(b.nftLogs) > 0 {
		err = tx.CreateInBatches(b.nftLogs, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}

	if len(b.reflows) > 0 {
		err = tx.CreateInBatches(b.reflows, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
	if len(b.destroys) > 0 {
		err = tx.CreateInBatches(b.destroys, 100).Error
		if err != nil {
			fmt.Println("=====", err)
			return
		}
	}
}

//套餐复购
func (b *BonusLogic) repeat() {
	var wd sync.WaitGroup
	buysAll, _ := lukComboBuyService.CountPerformanceAll(time.Now().Format(config.FORMAT_DATE_CST))
	for _, item := range b.repeats {
		wd.Add(1)
		go func(info luk.LukComboBuy) {
			var err error
			db := global.GVA_DB
			tx := db.Begin()
			defer func() {
				if err != nil {
					info.Status = -2
					info.StatusMsg = err.Error()
				} else {
					info.Status = 1
				}
				if tx.Save(&info).Error != nil {
					tx.Rollback()
				}
				tx.Commit()
				wd.Done()
			}()
			//查看套餐购买价格
			luks, err := lukComboService.GetLukCombo(info.ComboId)
			if err != nil {
				return
			}
			if luks.Status < 0 {
				err = errors.New("套餐已下架")
				return
			}

			buy := luk.LukComboBuy{
				Address:    info.Address,
				ComboId:    luks.ID,
				PriceUsdt:  info.PriceUsdt,
				Rate:       luks.Rate,
				Repeat:     info.Repeat,
				Days:       luks.Days,
				BuyAt:      time.Now(),
				MaturityAt: time.Now().AddDate(0, 0, luks.Days),
				Type:       1,
				TxHash:     info.TxHash,
			}
			//算出购买的套餐需要花费的LUK
			buy.Price = buy.PriceUsdt.Mul(b.LU)
			//算出复购的luk对应的usdt
			if !(buy.PriceUsdt.GreaterThanOrEqual(luks.PriceMin) && buy.PriceUsdt.LessThanOrEqual(luks.PriceMax)) {
				err = errors.New("购买价格不再可购买区间内")
				return
			}

			if num, ok := buysAll[luks.ID]; ok {
				if luks.Limit.Sub(num).Sub(buy.PriceUsdt).LessThanOrEqual(decimal.Zero) {
					err = errors.New("该套餐当日额度已售完")
					return
				} else {
					buysAll[luks.ID] = num.Add(buy.PriceUsdt)
				}
			}

			if tx.Create(&buy).Error != nil {
				tx.Rollback()
				return
			} else {
				//扣除钱包地址LUK数量
				data := luk.LukUser{}
				_ = tx.Model(&data).Where("address = ?", buy.Address).First(&data).Error
				if data.ID == 0 {
					err = errors.New("找不到对应用户地址")
					tx.Rollback()
					return
				}
				if data.Luk.LessThan(buy.Price) {
					err = errors.New("LUK额度不足")
					tx.Rollback()
					return
				}
				data.Luk = data.Luk.Sub(buy.Price)
				data.UpdatedAt = time.Now()
				err = tx.Save(&data).Error
			}
		}(item)
	}
	wd.Wait()
}

//销毁与回流 判断写入记录超过十分钟，还未传回哈希的进行删除
func (b *BonusLogic) DestroyReflow() {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("销毁与回流同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	//业绩返佣白名单
	rebates := lukServiceGroup.LukRebateService.QueryAll()
	buys := []luk.LukComboBuy{}
	err := global.GVA_DB.Debug().Where("buy_at < ? and destroy = 0", t.Add(-time.Minute*10)).Find(&buys).Error
	if err != nil {
		global.GVA_LOG.Error("查找记录失败!", zap.Error(err))
		return
	}
	wg := sync.WaitGroup{}
	status := emun.StatusTypeWait
	for _, item := range buys {
		wg.Add(1)
		go func(r luk.LukComboBuy) {
			var err error
			defer func() {
				//修改再删除
				if err != nil {
					r.StatusMsg = err.Error()
					r.ErrNum = r.ErrNum + 1
					_ = global.GVA_DB.Debug().Save(&r).Error
					if r.ErrNum > 10 || r.TxHash == "" {
						_ = lukComboBuyService.DeleteLukComboBuy(r.ID)
					}
				} else {
					r.Destroy = 1
					r.StatusMsg = ""
					_ = global.GVA_DB.Debug().Save(&r).Error
				}
				defer wg.Done()
			}()
			err = b.verifyHash(r)
			if err != nil {
				return
			}
			//80%销毁
			destroy := luk.LukDestroy{
				Address: r.Address,
				BuyId:   r.ID,
				Luk:     r.Price.Mul(decimal.NewFromFloat(0.8)),
				Status:  &status,
			}
			b.destroys = append(b.destroys, destroy)
			//20回流
			b.reflow(r.ID, r.Price)

			//业绩返佣
			b.performanceRebate(r, rebates)
		}(item)
		wg.Wait()
	}
}

func (b *BonusLogic) verifyHash(r luk.LukComboBuy) (err error) {
	//复购的不需要验证哈希
	if r.Type == 1 {
		return
	}
	//增加睡眠时间
	time.Sleep(3 * time.Second)
	if r.TxHash == "" {
		err = errors.New("十分钟后哈希还为空")
		return
	}
	//哈希解密处理
	bytesPass, err := base64.StdEncoding.DecodeString(r.TxHash)
	if err != nil {
		return
	}
	tpass, err := config.AesDecrypt(bytesPass, []byte("aidhh451bnffg2df"))
	if err != nil {
		return
	}
	//判断哈希状态
	fmt.Println("====", string(tpass))
	amounts, address, blockscan := (&tokenService.BscTokenService{}).DirectQueryTxHash(common.HexToHash(string(tpass)))
	fmt.Println("=====", amounts, address, blockscan)
	if blockscan == 0 {
		err = errors.New("哈希链上状态不对")
		return
	}

	fmt.Println("====", r.Price)
	if !amounts.Equal(r.Price) {
		err = errors.New("哈希链上金额不对")
		return
	}
	if strings.ToLower(address) != r.Address {
		err = errors.New("哈希链上地址不对")
		return
	}

	return
}

//回流数据组装
func (d *BonusLogic) reflow(id uint, price decimal.Decimal) {
	//联盟
	info := luk.LukReflow{
		BuyId:  id,
		Amount: price.Mul(decimal.NewFromFloat(0.08)),
		Type:   emun.ReflowTypeAlliance,
	}
	d.reflows = append(d.reflows, info)
	//STC底池
	info.Amount = price.Mul(decimal.NewFromFloat(0.03))
	info.Type = emun.ReflowTypeStcPot
	d.reflows = append(d.reflows, info)
	//STC节点分红
	info.Type = emun.ReflowTypeStcNode
	d.reflows = append(d.reflows, info)
	//运营钱包
	info.Type = emun.ReflowTypeOperate
	d.reflows = append(d.reflows, info)
	//NFT加权分红
	info.Amount = price.Mul(decimal.NewFromFloat(0.01))
	info.Type = emun.ReflowTypeNft
	d.reflows = append(d.reflows, info)
	//联盟经费
	info.Type = emun.ReflowTypeFunding
	d.reflows = append(d.reflows, info)
	//联盟运营岗位津贴
	info.Type = emun.ReflowTypeAllowance
	d.reflows = append(d.reflows, info)
}

//业绩返佣
func (d *BonusLogic) performanceRebate(buy luk.LukComboBuy, rebates []string) {
	//查询上级
	var PAddress string
	bonu := 0
	address := buy.Address
	for bonu <= 0 {
		user, _ := lukUserService.GetLukUserByAddress(d.tx, address)
		if user.PAddress == "" {
			break
		}
		//在白名单则可以跳出来
		if config.InStringArray(user.PAddress, rebates) {
			PAddress = user.PAddress
			bonu = 1
			break
		}
		address = user.PAddress
	}

	if PAddress != "" {
		var lukRebate luk.LukRebate
		err := d.tx.Model(&luk.LukRebate{}).Where("is_dividend = 0 and address = ?", PAddress).First(&lukRebate).Error
		if err != nil {
			return
		}
		if lukRebate.Remaining.IsZero() {
			lukRebate.IsDividend = true
			d.tx.Save(&lukRebate)
			return
		}
		rebatePerformance := global.GVA_LUK_CONFIG.RebatePerformance.Div(decimal.NewFromInt(100))
		detailed := luk.LukRevenueDetailed{
			Type:      emun.RevenueTypeAchievement,
			Category:  emun.DividendTypeRebate,
			TokenType: emun.LUK,
			Address:   PAddress,
			Remark:    rebatePerformance.String(),
			Source:    strconv.Itoa(int(buy.ID)),
		}
		number := buy.PriceUsdt.Mul(rebatePerformance)
		if lukRebate.Remaining.LessThanOrEqual(number) {
			detailed.Relatively = lukRebate.Remaining
			detailed.Amount = detailed.Relatively.Mul(d.LU)
			lukRebate.Remaining = decimal.Zero
			lukRebate.IsDividend = true
			d.tx.Save(&lukRebate)
		} else {
			lukRebate.Remaining = lukRebate.Remaining.Sub(number)
			detailed.Relatively = number
			detailed.Amount = detailed.Relatively.Mul(d.LU)
		}
		if detailed.Amount.GreaterThan(decimal.Zero) {
			d.lukRevenueDetailed = append(d.lukRevenueDetailed, detailed)
		}
		lukRebate.UpdatedAt = time.Now()
		d.tx.Save(&lukRebate)
	}
}

func (l *BonusLogic) LpBonus() {
	t := time.Now()
	defer func() {
		global.GVA_LOG.Info("LP分红同步开始时间：" + t.Format(config.FORMAT_DATETIME_CST) + ",同步结束时间：" + time.Now().Format(config.FORMAT_DATETIME_CST))
	}()
	var LukExchangeRecords []luk.LukExchangeRecords
	err := l.tx.Model(&luk.LukExchangeRecords{}).Where("status = 0").Find(&LukExchangeRecords).Error
	if err != nil {
		return
	}
	//分红逻辑
	lpLogics := LpBonusLogics{}
	//查出NFT的人数以及名单
	lpLogics.nftList = lukServiceGroup.LukNftService.QueryNftByNum()
	//算出LP地址以及占比
	lpLogics.lpList = lukServiceGroup.LukUserAddressService.QueryLpPro(l.tx)
	//分红占比
	lpLogics.pro = decimal.NewFromFloat(1).Div(decimal.NewFromInt(4)) //0.25
	var wd sync.WaitGroup
	for _, item := range LukExchangeRecords {
		wd.Add(1)
		go func(info luk.LukExchangeRecords) {
			defer func() {
				if err != nil {
					info.Errmsg = err.Error()
				}
				l.tx.Save(&info)
				wd.Done()
			}()
			if info.ChargeLuk.IsZero() {
				info.Status = -1
				info.Errmsg = "零手续费"
				return
			}
			lpLogics.data = info
			lpLogics.nftBonus()
			lpLogics.lppBonus()
			lpLogics.reflow()
			info.Status = 1
		}(item)
		wd.Wait()
	}

	l.lukRevenueDetailed = append(l.lukRevenueDetailed, lpLogics.lukRevenueDetailed...)
	l.lukRevenueDetailedNft = append(l.lukRevenueDetailedNft, lpLogics.lukRevenueDetailedNft...)
	l.reflowLps = append(l.reflowLps, lpLogics.reflowLps...)
}

type LpBonusLogics struct {
	nftList               map[string]decimal.Decimal
	lpList                map[string]decimal.Decimal
	pro                   decimal.Decimal
	lukRevenueDetailed    []luk.LukRevenueDetailed
	lukRevenueDetailedNft []luk.LukRevenueDetailedNft
	data                  luk.LukExchangeRecords
	reflowLps             []luk.LukReflowLp
}

//1%拥有nft
func (l *LpBonusLogics) nftBonus() {
	money := l.pro.Mul(l.data.ChargeLuk)
	for address, pro := range l.nftList {
		detailed := luk.LukRevenueDetailedNft{
			Address:   address,
			Amount:    money.Mul(pro),
			TokenType: emun.LUK,
			Source:    l.data.TxHash,
		}
		switch l.data.Type {
		case emun.MethodTypeSwap:
			detailed.Type = emun.RevenueTypeNftSwap
		case emun.MethodTypeSell:
			detailed.Type = emun.RevenueTypeNftSell
		case emun.MethodTypeRemoveLiquidity:
			detailed.Type = emun.RevenueTypeNftRemoveLiquidity
		}
		l.lukRevenueDetailedNft = append(l.lukRevenueDetailedNft, detailed)
	}
}

//%1lp占比分红
func (l *LpBonusLogics) lppBonus() {
	money := l.pro.Mul(l.data.ChargeLuk)
	for address, pro := range l.lpList {
		detailed := luk.LukRevenueDetailed{
			Address:   address,
			Amount:    money.Mul(pro),
			Category:  emun.DividendTypeLp,
			TokenType: emun.LUK,
			Source:    l.data.TxHash,
		}
		switch l.data.Type {
		case emun.MethodTypeSwap:
			detailed.Type = emun.RevenueTypeLpSwap
		case emun.MethodTypeSell:
			detailed.Type = emun.RevenueTypeLpSell
		case emun.MethodTypeRemoveLiquidity:
			detailed.Type = emun.RevenueTypeLpRemoveLiquidity
		}
		l.lukRevenueDetailed = append(l.lukRevenueDetailed, detailed)
	}
}

//回流
func (l *LpBonusLogics) reflow() {
	info := luk.LukReflowLp{
		TxHash: l.data.TxHash,
		Amount: l.pro.Mul(l.data.ChargeLuk),
		Type:   emun.ReflowLpTypePool,
	}
	l.reflowLps = append(l.reflowLps, info)
	info.Type = emun.ReflowLpTypeOperate
	l.reflowLps = append(l.reflowLps, info)
}

//将当前未过期套餐延长过期时间一个月
func (c *ComboTask) ExtendMaturityAt() {
	buys := []luk.LukComboBuy{}
	err := global.GVA_DB.Debug().Where("dividend = 0 and destroy = 1 and extend = 0").Find(&buys).Error
	if err != nil {
		global.GVA_LOG.Error("查找记录失败!", zap.Error(err))
		return
	}
	for _, item := range buys {
		item.Extend = 1
		item.MaturityAt = item.MaturityAt.AddDate(0, 0, 30)
		global.GVA_DB.Save(&item)
	}
}
