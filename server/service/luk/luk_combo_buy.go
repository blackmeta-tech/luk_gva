package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

type LukComboBuyService struct {
}

var lukComboBuyMt sync.RWMutex

// CreateLukComboBuy 创建LukComboBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboBuyService *LukComboBuyService) CreateLukComboBuy(buy luk.LukComboBuy) (errnum int, err error) {
	lukComboBuyMt.Lock()
	defer lukComboBuyMt.Unlock()
	if buy.Address == "" || buy.ComboId == 0 || buy.Price.LessThanOrEqual(decimal.Zero) || buy.PriceUsdt.LessThanOrEqual(decimal.Zero) {
		err = errors.New("参数不足")
		return -1, err
	}
	buy.Address = strings.ToLower(buy.Address)
	if lukComboBuyService.GetByAddressHash(buy.Address) {
		err = errors.New("套餐已购买，请勿重复购买")
		return -2, err
	}
	if lukComboBuyService.GetByAddress(buy.Address) {
		err = errors.New("套餐已发起过购买，请等十五分钟后再试")
		return -2, err
	}
	combo, err := (&LukComboService{}).GetLukCombo(buy.ComboId)
	if combo.ID == 0 || err != nil {
		err = errors.New("套餐查询不到，请重试")
		return -3, err
	}
	if combo.Status != 0 {
		err = errors.New("该套餐已下架")
		return -4, err
	}
	if !(buy.PriceUsdt.GreaterThanOrEqual(combo.PriceMin) && buy.PriceUsdt.LessThanOrEqual(combo.PriceMax)) {
		err = errors.New("购买价格不再可购买区间内")
		return -5, err
	}

	buys, _ := lukComboBuyService.CountPerformanceByDate(time.Now().Format(config.FORMAT_DATE_CST), combo.ID)
	if combo.Limit.Sub(buys).LessThanOrEqual(decimal.Zero) {
		err = errors.New("该套餐当日额度已售完")
		return -6, err
	}
	if combo.Limit.Sub(buys).Sub(buy.PriceUsdt).LessThanOrEqual(decimal.Zero) {
		err = errors.New("该套餐剩余额度不足")
		return -6, err
	}
	if combo.Performance.GreaterThan(decimal.Zero) {
		_, performances := lukComboBuyService.QueryPerformance(buy.Address)
		if combo.Performance.GreaterThan(performances) {
			err = errors.New("旗下业绩未达标，不可购买")
			return -7, err
		}
	}
	buy.Rate = combo.Rate
	buy.Days = combo.Days
	buy.BuyAt = time.Now()
	buy.MaturityAt = time.Now().AddDate(0, 0, combo.Days)
	err = global.GVA_DB.Create(&buy).Error
	if err != nil {
		errnum = -8
	}
	errnum = int(buy.ID)
	return
}

// DeleteLukComboBuy 删除LukComboBuy记录
func (lukBannerService *LukComboBuyService) DeleteLukComboBuy(id uint) (err error) {
	err = global.GVA_DB.Debug().Delete(&luk.LukComboBuy{}, "id = ?", id).Error
	return err
}

// UpdateLukComboBuy 更新LukComboBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboBuyService *LukComboBuyService) UpdateLukComboBuy(lukComboBuy luk.LukComboBuy) (err error) {
	buy := luk.LukComboBuy{}
	err = global.GVA_DB.Where("id = ? and address = ?", lukComboBuy.ID, lukComboBuy.Address).First(&buy).Error
	if err != nil {
		return
	}
	if buy.ID == 0 {
		return errors.New("找不到记录")
	}
	buy.Repeat = lukComboBuy.Repeat
	err = global.GVA_DB.Save(&buy).Error
	return err
}

func (lukComboBuyService *LukComboBuyService) UpdateLukComboBuyByHash(lukComboBuy luk.LukComboBuy) (err error) {
	buy := luk.LukComboBuy{}
	err = global.GVA_DB.Where("id = ? and address = ?", lukComboBuy.ID, lukComboBuy.Address).First(&buy).Error
	if err != nil {
		return
	}
	if buy.ID == 0 {
		return errors.New("找不到记录")
	}
	buy.BuyAt = time.Now()
	buy.TxHash = lukComboBuy.TxHash
	err = global.GVA_DB.Save(&buy).Error
	return err
}

func (lukComboBuyService *LukComboBuyService) Update(tx *gorm.DB, info luk.LukComboBuy) (err error) {
	err = tx.Save(&info).Error
	return err
}

// GetLukComboBuy 根据id获取LukComboBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboBuyService *LukComboBuyService) GetLukComboBuy(id uint) (lukComboBuy luk.LukComboBuy, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukComboBuy).Error
	return
}

func (lukComboBuyService *LukComboBuyService) __Query(info lukReq.LukComboBuySearch) (db *gorm.DB) {
	// 创建db
	db = global.GVA_DB.Debug().Model(&luk.LukComboBuy{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db.Where("tx_hash != ''")
	//获取下级业绩
	if info.Subclass > 0 {
		treeAddress := (&LukUserService{}).GetTreeAddress(info.Address)
		treeAddress = append(treeAddress, info.Address)
		db.Where("address in ?", treeAddress)
	} else {
		if info.Address != "" {
			db.Where("address = ?", info.Address)
		}
	}

	if info.ComboId != 0 {
		db.Where("combo_id = ?", info.ComboId)
	}
	if info.Mode > 0 {
		t := time.Now()
		if info.Mode == 1 {
			db.Where("(maturity_at >= ? or dividend = 0)", t)
		} else {
			db.Where("(maturity_at < ? or dividend = 1)", t)
		}
	}
	if len(info.Generation) > 0 {
		db.Where("address in ?", info.Generation)
	}
	return
}

// GetLukComboBuyInfoList 分页获取LukComboBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboBuyService *LukComboBuyService) GetLukComboBuyInfoList(info lukReq.LukComboBuySearch) (list []luk.LukComboBuy, total int64, err error) {
	var lukComboBuys []luk.LukComboBuy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := lukComboBuyService.__Query(info)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukComboBuys).Error
	return lukComboBuys, total, err
}

//组装前端购买接口
func (lukComboBuyService *LukComboBuyService) QueryList(info lukReq.LukComboBuySearch) (list []response.LukComboBuyList, num int64, err error) {
	lukComboBuys, total, err := lukComboBuyService.GetLukComboBuyInfoList(info)
	if err != nil {
		return
	}
	list = make([]response.LukComboBuyList, 0)
	combo := (&LukComboService{}).QueryAll()
	for _, item := range lukComboBuys {
		info := response.LukComboBuyList{
			ID:         item.ID,
			ComboId:    item.ComboId,
			BuyAt:      item.BuyAt,
			Days:       item.Days,
			MaturityAt: item.MaturityAt,
			Price:      item.Price,
			PriceU:     item.PriceUsdt,
			Rate:       item.Rate,
			Repeat:     item.Repeat,
		}
		if c, ok := combo[info.ComboId]; ok {
			info.Pic = c.Pic
		}
		list = append(list, info)
	}
	num = total
	return
}

//统计查询条件金额
func (lukComboBuyService *LukComboBuyService) QuerySummary(info lukReq.LukComboBuySearch) (summary decimal.Decimal) {
	var lukComboBuy luk.LukComboBuy
	db := lukComboBuyService.__Query(info)
	err := db.Select("sum(price) as price").First(&lukComboBuy).Error
	if err == nil {
		summary = lukComboBuy.Price
	}
	return
}

//获取某天购买量
func (lukComboBuyService *LukComboBuyService) CountPerformanceByDate(d string, comboId uint) (priceUsdt decimal.Decimal, err error) {
	lukComboBuys := luk.LukComboBuy{}
	err = global.GVA_DB.Model(&luk.LukComboBuy{}).Select("sum(price_usdt) price_usdt").
		Where("buy_at >= ? and combo_id = ?", d, comboId).First(&lukComboBuys).Error
	if err == nil {
		priceUsdt = lukComboBuys.PriceUsdt
	}
	return
}

//统计全部套餐某天的购买量
func (lukComboBuyService *LukComboBuyService) CountPerformanceAll(d string) (comboByUsdt map[uint]decimal.Decimal, err error) {
	comboByUsdt = make(map[uint]decimal.Decimal, 0)
	lukComboBuys := []luk.LukComboBuy{}
	err = global.GVA_DB.Model(&luk.LukComboBuy{}).Select("sum(price_usdt) price_usdt, combo_id").
		Where("buy_at >= ? ", d).Group("combo_id").Find(&lukComboBuys).Error
	if err == nil {
		for _, item := range lukComboBuys {
			comboByUsdt[item.ComboId] = item.PriceUsdt
		}
	}
	return
}

//获取昨天的购买量，旗下的
func (lukComboBuyService *LukComboBuyService) CountByDateAndAddress(d string, address []string) (priceUsdt decimal.Decimal, err error) {
	lukComboBuys := luk.LukComboBuy{}
	err = global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Select("sum(price_usdt) price_usdt").
		Where("buy_at >= ? and buy_at <= ? and address in ?", d+" 00:00:00", d+" 23:59:59", address).First(&lukComboBuys).Error
	if err == nil {
		priceUsdt = lukComboBuys.PriceUsdt
	}
	return
}

//获取未到期的套餐花费价格
func (lukComboBuyService *LukComboBuyService) GetPriceByAddress(address string) (price decimal.Decimal) {
	t := time.Now()
	lukComboBuys := luk.LukComboBuy{}
	_ = global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Where("status = 0 and address = ? and (maturity_at >= ? or dividend = 0)", address, t).First(&lukComboBuys).Error
	if lukComboBuys.ID > 0 {
		price = lukComboBuys.PriceUsdt
	}
	return
}

//获取是否有购买套餐
func (lukComboBuyService *LukComboBuyService) GetByAddress(address string) (isBuy bool) {
	t := time.Now()
	lukComboBuys := luk.LukComboBuy{}
	_ = global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Where("status = 0 and address = ? and (maturity_at >= ? or dividend = 0)", address, t).First(&lukComboBuys).Error
	if lukComboBuys.ID > 0 {
		isBuy = true
	}
	return
}

//获取是否有购买套餐 排除空哈希情况
func (lukComboBuyService *LukComboBuyService) GetByAddressHash(address string) (isBuy bool) {
	t := time.Now()
	lukComboBuys := luk.LukComboBuy{}
	_ = global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Where("tx_hash != '' and status = 0 and address = ? and (maturity_at >= ? or dividend = 0)", address, t).First(&lukComboBuys).Error
	if lukComboBuys.ID > 0 {
		isBuy = true
	}
	return
}

func (lukComboBuyService *LukComboBuyService) QueryPerformance(address string) (priceLuk decimal.Decimal, priceUsdt decimal.Decimal) {
	//获取旗下全部用户地址
	treeAddress := (&LukUserService{}).GetTreeAddress(address)
	treeAddress = append(treeAddress, address)
	lukComboBuys := luk.LukComboBuy{}
	err := global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Select("sum(price) price, sum(price_usdt) price_usdt").
		Where("address in ?", treeAddress).First(&lukComboBuys).Error
	if err == nil {
		priceLuk = lukComboBuys.Price
		priceUsdt = lukComboBuys.PriceUsdt
	}
	return
}

//地址集合未过期的业绩
func (lukComboBuyService *LukComboBuyService) QueryPerformanceByaddresss(address []string) (priceUsdt decimal.Decimal) {
	t := time.Now()
	lukComboBuys := luk.LukComboBuy{}
	err := global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Select("sum(price_usdt) price_usdt").
		Where("address in ? and (maturity_at >= ? or dividend = 0)", address, t).First(&lukComboBuys).Error
	if err == nil {
		priceUsdt = lukComboBuys.PriceUsdt
	}
	return
}

func (lukComboBuyService *LukComboBuyService) GetByAddresss(addresss []string) (count int64) {
	t := time.Now()
	_ = global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Where("status = 0 and address in ? and (maturity_at >= ? or dividend = 0)", addresss, t).Group("address").Count(&count).Error
	return
}

//获取全部下级的交易记录
func (lukComboBuyService *LukComboBuyService) GetGenerationByAddresss(info lukReq.LukComboBuySearch) (list []luk.LukComboBuy, total int64, err error) {
	list = make([]luk.LukComboBuy, 0)
	if info.Address == "" {
		return
	}
	info.Generation = (&LukUserService{}).AddressAllBranch(info.Address)
	info.Mode = 1
	info.Address = ""
	var lukComboBuys []luk.LukComboBuy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := lukComboBuyService.__Query(info)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukComboBuys).Error
	return lukComboBuys, total, err
}
