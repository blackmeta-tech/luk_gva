package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/shopspring/decimal"
	"time"
)

type LukComboService struct {
}

// CreateLukCombo 创建LukCombo记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboService *LukComboService) CreateLukCombo(lukCombo luk.LukCombo) (err error) {
	combo := luk.LukCombo{}
	global.GVA_DB.Where("days = ? and status = 0", lukCombo.Days).First(&combo)
	if combo.ID > 0 {
		err = errors.New("已存在对应天数套餐")
		return
	}
	err = global.GVA_DB.Create(&lukCombo).Error
	return err
}

// UpdateLukCombo 更新LukCombo记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboService *LukComboService) UpdateLukCombo(lukCombo luk.LukCombo) (err error) {
	err = global.GVA_DB.Save(&lukCombo).Error
	return err
}

// GetLukCombo 根据id获取LukCombo记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboService *LukComboService) GetLukCombo(id uint) (lukCombo luk.LukCombo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukCombo).Error
	return
}

// GetLukComboInfoList 分页获取LukCombo记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukComboService *LukComboService) GetLukComboInfoList(info lukReq.LukComboSearch) (list []luk.LukCombo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukCombo{})
	var lukCombos []luk.LukCombo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Days != 0 {
		db = db.Where("days = ?", info.Days)
	}
	if info.Statu != nil {
		db = db.Where("status = ?", info.Statu)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("days asc").Find(&lukCombos).Error
	return lukCombos, total, err
}

func (lukComboService *LukComboService) GetLukComboInfoListByAddress(address string) (list []response.LukComboList, err error) {
	list = make([]response.LukComboList, 0)
	var lukCombos []luk.LukCombo
	err = global.GVA_DB.Model(&luk.LukCombo{}).Where("status = 0").Order("days asc").Find(&lukCombos).Error
	if err != nil {
		return
	}
	t := time.Now().Format(config.FORMAT_DATE_CST)
	buyService := (&LukComboBuyService{})
	//统计旗下业绩
	_, performances := buyService.QueryPerformance(address)
	for _, itme := range lukCombos {
		info := response.LukComboList{
			ComboId:        itme.ID,
			PriceMax:       itme.PriceMax,
			PriceMin:       itme.PriceMin,
			Rate:           itme.Rate,
			Limit:          itme.Limit,
			Days:           itme.Days,
			Pic:            itme.Pic,
			Performance:    itme.Performance,
			PerformanceNow: performances,
		}
		reason := make(map[int]string, 0)
		if info.Days >= 100 {
			info.Remaining = info.Limit
			reason[4] = "您暂未有购买权限"
		} else {
			//统计购买额度
			buyNum, _ := buyService.CountPerformanceByDate(t, itme.ID)
			info.Remaining = info.Limit.Sub(buyNum)
			if info.Remaining.LessThanOrEqual(decimal.Zero) {
				reason[1] = "今日额度已售完，请明日再来"
			}
			//判断当前套餐自己是否有买
			if buyService.GetByAddressHash(address) {
				reason[2] = "您的套餐未到期，不可购买"
			}
			//判断业绩
			if info.Performance.GreaterThan(decimal.Zero) && info.Performance.GreaterThan(performances) {
				reason[3] = "旗下业绩未达标，不可购买"
			}
		}
		info.Reason = reason
		list = append(list, info)
	}
	return
}

//获取全部未下架的键值对
func (lukComboService *LukComboService) QueryAll() (list map[uint]luk.LukCombo) {
	list = make(map[uint]luk.LukCombo, 0)
	var lukCombos []luk.LukCombo
	err := global.GVA_DB.Model(&luk.LukCombo{}).Where("status = 0").Find(&lukCombos).Error
	if err != nil {
		return
	}
	for _, t := range lukCombos {
		list[t.ID] = t
	}
	return
}

//获取业绩以及对应套餐所需业绩
func (lukComboService *LukComboService) QueryPerformance(address string) (list []response.LukComboPerformance, performances, performancesU decimal.Decimal, err error) {
	list = make([]response.LukComboPerformance, 0)
	var lukCombos []luk.LukCombo
	err = global.GVA_DB.Model(&luk.LukCombo{}).Where("status = 0").Order("days asc").Find(&lukCombos).Error
	if err != nil {
		return
	}
	buyService := (&LukComboBuyService{})
	//统计旗下业绩
	performances, performancesU = buyService.QueryPerformance(address)
	for _, itme := range lukCombos {
		info := response.LukComboPerformance{
			Days:        itme.Days,
			Performance: itme.Performance,
		}
		if info.Performance.GreaterThan(decimal.Zero) && info.Performance.GreaterThan(performancesU) {
			info.State = 0
		} else {
			info.State = 1
		}
		list = append(list, info)
	}
	return
}
