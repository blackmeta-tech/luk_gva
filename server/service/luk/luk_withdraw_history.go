package luk

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	tokenReq "github.com/flipped-aurora/gin-vue-admin/server/model/token/request"
	tokenService "github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
)

type LukWithdrawHistoryService struct {
}

// CreateLukWithdrawHistory 创建LukWithdrawHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (WithdrawHistoryService *LukWithdrawHistoryService) CreateLukWithdrawHistory(tx *gorm.DB, apply luk.LukWithdrawApply) (err error) {
	msg := ""
	defer func() {
		if msg != "" {
			err = errors.New(msg)
		}
	}()
	times := time.Now()
	history := luk.LukWithdrawHistory{
		Address:       apply.Address,
		TokenType:     apply.TokenType,
		AmountPrimary: apply.AmountPrimary,
		Amount:        apply.Amount,
		Procedures:    apply.Procedures,
		Time:          &times,
		Status:        new(emun.StatusType),
	}
	history.CreatedAt = apply.CreatedAt
	history.UpdatedAt = apply.UpdatedAt

	//直接调用转账接口
	info := tokenReq.TransferFrom{
		PrivateKey: global.GVA_CONFIG.Wallet.WithdrawPrivateKey.String(),
		ToAddress:  common.HexToAddress(apply.Address),
		Amount:     apply.Amount,
		TokenType:  *apply.TokenType,
	}
	signedTx, err1 := (&tokenService.BscTokenService{}).TransferFrom(info)
	if err1 != nil {
		err = err1
		return
	}
	history.SignedTx = signedTx.String()
	if err = tx.Create(&history).Error; err != nil {
		global.GVA_LOG.Error("LukWithdrawHistory创建失败!", zap.Error(err))
		return
	}

	//查询状态并自动更新数据
	WithdrawHistoryService.UpdateLukWithdrawHistory(signedTx)
	return
}

// UpdateLukWithdrawHistory 更新LukWithdrawHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (WithdrawHistoryService *LukWithdrawHistoryService) UpdateLukWithdrawHistory(txHashStr common.Hash) (err error) {
	go func(txHashStr common.Hash) {
		_, blockscan := (&tokenService.BscTokenService{}).QueryTxHash(txHashStr)
		info := luk.LukWithdrawHistory{}
		err := global.GVA_DB.Where("signed_tx = ?", txHashStr.String()).First(&info).Error
		if err == nil {
			if blockscan == 1 {
				*info.Status = emun.StatusTypeSuccess
			}
			if err = global.GVA_DB.Save(&info).Error; err != nil {
				global.GVA_LOG.Error("LukWithdrawHistory更新失败!", zap.Error(err))
			}
		} else {
			global.GVA_LOG.Error("LukWithdrawHistory更新失败!", zap.Error(err))
		}
	}(txHashStr)
	return
}

//修改提现记录
func (WithdrawHistoryService *LukWithdrawHistoryService) Update(data luk.LukWithdrawHistory) (err error) {
	if err = global.GVA_DB.Save(&data).Error; err != nil {
		return
	}
	//转账失败回退金额
	if *data.Status == emun.StatusTypeFail {
		db := global.GVA_DB
		tx := db.Begin()
		defer tx.Commit()
		if err = (&LukUserService{}).UpdateWalletBalance(tx, data.Address, *data.TokenType, data.AmountPrimary, true); err != nil {
			tx.Rollback()
			global.GVA_LOG.Error("LukUser更新失败!", zap.Error(err))
			return
		}
	}
	return
}

// GetLukWithdrawHistory 根据id获取LukWithdrawHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (WithdrawHistoryService *LukWithdrawHistoryService) GetLukWithdrawHistory(id uint) (WithdrawHistory luk.LukWithdrawHistory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&WithdrawHistory).Error
	return
}

// GetLukWithdrawHistory 获取LukWithdrawHistory 转帐status 未确认记录
// Author [piexlmax](https://github.com/piexlmax)
func (WithdrawHistoryService *LukWithdrawHistoryService) GetLukWithdrawHistoryStatusRecord() (WithdrawHistory []luk.LukWithdrawHistory, err error) {
	err = global.GVA_DB.Where("status = 0 AND signed_tx IS NOT NULL").Find(&WithdrawHistory).Error
	return
}

type QueryLukWithdrawHistory struct {
	Info lukReq.LukWithdrawHistorySearch
}

func (q QueryLukWithdrawHistory) __query() *gorm.DB {
	db := global.GVA_DB.Debug().Model(&luk.LukWithdrawHistory{})
	info := q.Info
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
	// 日期条件添加
	nilTime := time.Time{} //赋零值
	if info.GVA_MODEL.CreatedAt != nilTime {
		day := fmt.Sprintf("%d-%02d-%02d", info.GVA_MODEL.CreatedAt.Year(), info.GVA_MODEL.CreatedAt.Month(), info.GVA_MODEL.CreatedAt.Day())
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m-%d') = ? ", day)
	}

	if info.StartDate != "" {
		db = db.Where("time >= ?", info.StartDate)
	}

	if info.EndDate != "" {
		db = db.Where("time <= ?", info.EndDate)
	}
	return db
}

// GetLukWithdrawHistoryInfoList 分页获取LukWithdrawHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (q QueryLukWithdrawHistory) GetLukWithdrawHistoryInfoList() (list interface{}, total int64, err error) {
	limit := q.Info.PageSize
	offset := q.Info.PageSize * (q.Info.Page - 1)
	// 创建db
	var WithdrawHistorys []luk.LukWithdrawHistory
	db := q.__query() //查询条件
	err = db.Count(&total).Error
	if err != nil {
		return
	} else {
		db = db.Limit(limit).Offset(offset)
		if q.Info.OrderKey != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["amount_primary"] = true
			orderMap["procedures"] = true
			if orderMap[q.Info.OrderKey] {
				if q.Info.Desc {
					OrderStr = q.Info.OrderKey + " desc"
				} else {
					OrderStr = q.Info.OrderKey
				}
			} else {
				err = fmt.Errorf("非法的排序字段: %v", q.Info.OrderKey)
				return
			}
			db.Order(OrderStr)
		} else {
			db.Order("id desc")
		}
	}
	err = db.Find(&WithdrawHistorys).Error
	return WithdrawHistorys, total, err
}

//统计提现地址个数
func (q QueryLukWithdrawHistory) GetCountAddress() (num int64, err error) {
	db := q.__query() //查询条件
	err = db.Group("user_address").Count(&num).Error
	return
}

//分别统计USDT以及WHO个数
func (q QueryLukWithdrawHistory) GetCountAmount() (lukAmount decimal.Decimal, usdtAmount decimal.Decimal) {
	lukinfo := luk.LukWithdrawHistory{}
	usdtinfo := luk.LukWithdrawHistory{}
	db := q.__query() //查询条件
	db.Select("sum(amount) as amount").Where("token_type = 3").First(&lukinfo)
	lukAmount = lukinfo.Amount
	db1 := q.__query() //查询条件
	db1.Select("sum(amount) as amount").Where("token_type = 2").First(&usdtinfo)
	usdtAmount = usdtinfo.Amount
	return
}

//SumLukWithdrawHistory
func (q *LukWithdrawHistoryService) SumLukWithdrawHistory(tx *gorm.DB, address string, tokenType emun.TokenType) (WithdrawHistory luk.LukWithdrawHistory, err error) {
	err = tx.Debug().Model(&luk.LukWithdrawHistory{}).Select("sum(amount_primary) as amount_primary").
		Where("address = ?", address).
		Where("token_type = ?", tokenType).
		Where("time >= ?", time.Now().Format(config.FORMAT_DATE_CST)).
		Where("status >= ?", 0).
		First(&WithdrawHistory).Error
	return
}

//获取提现钱包额度
func (WithdrawHistoryService *LukWithdrawHistoryService) GetWithdrawalBalance() (data map[emun.TokenType]decimal.Decimal) {
	data = make(map[emun.TokenType]decimal.Decimal, 0)
	tokenService := &tokenService.BscTokenService{}
	var transfer token.Transfer
	transfer.From = global.GVA_CONFIG.Wallet.Withdraw
	transfer.Type = emun.LUK
	if reBscTokenBalanceOf, err := tokenService.GetBalanceOf(transfer); err != nil {
		global.GVA_LOG.Error("查询提现钱包失败!", zap.Error(err))
	} else {
		data[transfer.Type] = reBscTokenBalanceOf
	}
	transfer.From = global.GVA_CONFIG.Wallet.Withdraw
	transfer.Type = emun.USDT
	if reBscTokenBalanceOf, err := tokenService.GetBalanceOf(transfer); err != nil {
		global.GVA_LOG.Error("查询提现钱包失败!", zap.Error(err))
	} else {
		data[transfer.Type] = reBscTokenBalanceOf
	}
	return
}
