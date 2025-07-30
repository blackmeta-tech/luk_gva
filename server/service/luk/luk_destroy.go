package luk

import (
	"github.com/ethereum/go-ethereum/common"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	tokenReq "github.com/flipped-aurora/gin-vue-admin/server/model/token/request"
	tokenService "github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

type LukDestroyService struct {
}

//luk销毁记录
func (d *LukDestroyService) Destroy() (err error) {
	list := []luk.LukDestroy{}
	err = global.GVA_DB.Debug().Where("tx_hash = ''").Find(&list).Error
	if err != nil {
		return
	}
	for _, item := range list {
		d.DoDestroy(item)
	}
	//二次销毁因哈希出错问题
	t := time.Now()
	err = global.GVA_DB.Debug().Where("tx_hash != '' and status = 0 and created_at < ?", t.Add(-time.Minute*10)).Find(&list).Error
	bt := &tokenService.BscTokenService{}
	for _, item := range list {
		_, blockscan := bt.QueryTxHash(common.HexToHash(item.TxHash))
		if blockscan == 0 {
			d.DoDestroy(item)
		} else {
			updates := map[string]interface{}{
				"updated_at": time.Now(),
				"status":     emun.StatusTypeSuccess,
			}
			global.GVA_DB.Debug().Model(&luk.LukDestroy{}).Where("tx_hash = ?", item.TxHash).Updates(updates)
		}
	}
	return
}

//销毁操作
func (d *LukDestroyService) DoDestroy(item luk.LukDestroy) {
	time.Sleep(3 * time.Second)
	if item.Luk.Equal(decimal.Zero) {
		return
	}
	info := tokenReq.TransferFrom{
		PrivateKey: global.GVA_CONFIG.Wallet.DestroyPrivateKey.String(),
		ToAddress:  common.HexToAddress("0x0000000000000000000000000000000000000001"),
		Amount:     item.Luk,
		TokenType:  emun.LUK,
	}
	hash, err := (&tokenService.BscTokenService{}).TransferFrom(info)
	if err == nil {
		item.TxHash = hash.Hex()
		item.UpdatedAt = time.Now()
		if global.GVA_DB.Save(&item).Error == nil {
			d.Update(hash)
		}
	}
	return
}
func (d *LukDestroyService) Update(txHashStr common.Hash) {
	go func(txHashStr common.Hash) {
		_, blockscan := (&tokenService.BscTokenService{}).QueryTxHash(txHashStr)
		if blockscan == 1 {
			updates := map[string]interface{}{
				"updated_at": time.Now(),
				"status":     emun.StatusTypeSuccess,
			}
			err := global.GVA_DB.Debug().Model(&luk.LukDestroy{}).Where("tx_hash = ?", txHashStr.String()).Updates(updates).Error
			if err != nil {
				global.GVA_LOG.Error("LukDestroy更新失败!", zap.Error(err))
			}
		}
	}(txHashStr)
	return
}

func (d *LukDestroyService) GetLukDestroyInfoList(info lukReq.LukDestroySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukDestroy{})
	var lukDestroys []luk.LukDestroy
	if info.Address != "" {
		db.Where("address = ?", info.Address)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukDestroys).Error
	return lukDestroys, total, err
}

//统计查询条件金额
func (d *LukDestroyService) QuerySummary(address string) (summary decimal.Decimal) {
	db := global.GVA_DB.Model(&luk.LukDestroy{})
	var lukDestroy luk.LukDestroy
	if address != "" {
		db.Where("address = ?", address)
	}
	err := db.Select("sum(luk) as luk").First(&lukDestroy).Error
	if err == nil {
		summary = lukDestroy.Luk
	}
	return
}
