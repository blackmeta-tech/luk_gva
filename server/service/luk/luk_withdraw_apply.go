package luk

import (
	"errors"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"strings"
	"sync"
)

type LukWithdrawApplyService struct {
}

var withdrawApply sync.RWMutex

// CreateLukWithdrawApply 创建LukWithdrawApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawApplyService *LukWithdrawApplyService) CreateLukWithdrawApply(apply luk.LukWithdrawApply) (msg string, err error) {
	withdrawApply.Lock()
	defer withdrawApply.Unlock()
	if apply.Address == "" || apply.AmountPrimary.IsZero() || *apply.TokenType == 0 {
		err = errors.New("参数不齐全")
		return
	}
	apply.Address = strings.ToLower(apply.Address)
	blacklist, _ := (&LukWithdrawBlacklistService{}).QueryLukWithdrawBlacklist(apply.Address)
	if blacklist.ID > 0 {
		err = errors.New("地址冻结，请联系客服")
		return
	}

	if *apply.TokenType == emun.LUK && apply.AmountPrimary.LessThan(global.GVA_LUK_CONFIG.WithdrawalQuota) {
		err = errors.New("提现额度低于" + global.GVA_LUK_CONFIG.WithdrawalQuota.String())
		return
	}
	db := global.GVA_DB
	tx := db.Begin()
	defer tx.Commit()
	user := &LukUserService{}                       //钱包相关逻辑
	withdrawHistory := &LukWithdrawHistoryService{} //提现历史相关逻辑
	uintData, _ := user.GetLukUserByAddress(tx, apply.Address)
	if uintData.ID == 0 {
		err = errors.New("查询不到该地址币种类型信息")
		tx.Rollback()
		return
	}
	apply.Address = uintData.Address
	var amount decimal.Decimal
	if *apply.TokenType == emun.USDT {
		amount = uintData.Usdt
	} else if *apply.TokenType == emun.LUK {
		amount = uintData.Luk
	}
	if apply.AmountPrimary.GreaterThan(amount) {
		err = errors.New("额度已超出限制")
		tx.Rollback()
		return
	}

	//手续费
	if global.GVA_LUK_CONFIG.WithdrawalCharge > 0 {
		charge := float64(global.GVA_LUK_CONFIG.WithdrawalCharge) / 100 //后台配置的是百分比
		apply.Procedures = apply.AmountPrimary.Mul(decimal.NewFromFloat(charge))
	}

	apply.Amount = apply.AmountPrimary.Sub(apply.Procedures) //实际提现金额=提现金额-手续费

	//查询当日最高提现额度
	data, _ := withdrawHistory.SumLukWithdrawHistory(tx, apply.Address, *apply.TokenType)
	data.AmountPrimary = data.AmountPrimary.Add(apply.AmountPrimary) //提现金额加上当日累计金额做对比

	if (*apply.TokenType == emun.USDT && data.AmountPrimary.GreaterThanOrEqual(decimal.New(global.GVA_LUK_CONFIG.WithdrawalUsdt, 0))) || (*apply.TokenType == emun.LUK && data.AmountPrimary.GreaterThanOrEqual(decimal.New(global.GVA_LUK_CONFIG.WithdrawalLuk, 0))) {
		//超过当日最高提现额度需要审核
		apply.Status = new(emun.StatusType)
	} else {
		//需要判断提现钱包是否有额度，没有则进入审核状态
		balance := withdrawHistory.GetWithdrawalBalance()
		if info, ok := balance[*apply.TokenType]; ok && info.GreaterThan(apply.Amount) {
			status := emun.StatusTypeSuccess
			apply.Status = &status
			err = withdrawHistory.CreateLukWithdrawHistory(tx, apply)
			if err != nil {
				tx.Rollback()
				return
			}
		} else {
			apply.Status = new(emun.StatusType)
		}
	}

	if err = tx.Create(&apply).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error("LukWithdrawApply创建失败!", zap.Error(err))
		return
	} else {
		//修改钱包余额
		if err = user.UpdateWalletBalance(tx, apply.Address, *apply.TokenType, apply.AmountPrimary, false); err != nil {
			global.GVA_LOG.Error("LukUserWallet更新失败!", zap.Error(err))
			tx.Rollback()
			return
		}
	}
	if *apply.Status == 0 {
		msg = "超过限额，审核中"
	} else {
		msg = "待转账"
	}
	return
}

// UpdateLukWithdrawApply 更新LukWithdrawApply记录
func (lukWithdrawApplyService *LukWithdrawApplyService) UpdateLukWithdrawApply(apply luk.LukWithdrawApply) (err error) {
	withdrawApply.Lock()
	defer withdrawApply.Unlock()
	db := global.GVA_DB
	tx := db.Begin()
	defer tx.Commit()
	withdrawHistory := &LukWithdrawHistoryService{} //提现历史相关逻辑
	//审核通过时，判断下提现钱包额度
	balance := withdrawHistory.GetWithdrawalBalance()
	if *apply.Status == emun.StatusTypeSuccess {
		if info, ok := balance[*apply.TokenType]; !ok {
			err = errors.New("获取不到提现钱包额度")
			return
		} else if info.LessThan(apply.Amount) {
			err = errors.New("提现钱包额度不足")
			return
		}
	}
	if err = tx.Save(&apply).Error; err != nil {
		tx.Rollback()
		return
	} else {
		if *apply.Status == emun.StatusTypeSuccess {
			if err = withdrawHistory.CreateLukWithdrawHistory(tx, apply); err != nil {
				tx.Rollback()
				return
			}
		} else if *apply.Status == emun.StatusTypeFail {
			if err = (&LukUserService{}).UpdateWalletBalance(tx, apply.Address, *apply.TokenType, apply.AmountPrimary, true); err != nil {
				tx.Rollback()
				global.GVA_LOG.Error("LukUser更新失败!", zap.Error(err))
				return
			}
		}
	}
	return err
}

// GetLukWithdrawApply 根据id获取LukWithdrawApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawApplyService *LukWithdrawApplyService) GetLukWithdrawApply(id uint) (lukWithdrawApply luk.LukWithdrawApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukWithdrawApply).Error
	return
}

// GetLukWithdrawApplyInfoList 分页获取LukWithdrawApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukWithdrawApplyService *LukWithdrawApplyService) GetLukWithdrawApplyInfoList(info lukReq.LukWithdrawApplySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Debug().Model(&luk.LukWithdrawApply{})
	var lukWithdrawApplys []luk.LukWithdrawApply
	if info.ID != 0 {
		db.Where("id = ?", info.ID)
	}
	if info.Address != "" {
		db.Where("address = ?", info.Address)
	}
	if info.Status != nil {
		db = db.Where("status = ? ", info.Status)
	}
	if info.TokenType != nil {
		if *info.TokenType != 0 {
			db = db.Where("token_type = ? ", info.TokenType)
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukWithdrawApplys).Error
	return lukWithdrawApplys, total, err
}
