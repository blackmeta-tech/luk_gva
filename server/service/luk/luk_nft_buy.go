package luk

import (
	"errors"
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
	"strings"
	"sync"
	"time"
)

type LukNftBuyService struct {
}

var lukNftBuyMt sync.RWMutex

// CreateLukNftBuy 创建LukNftBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftBuyService *LukNftBuyService) CreateLukNftBuy(buy luk.LukNftBuy) (err error) {
	lukNftBuyMt.Lock()
	defer lukNftBuyMt.Unlock()
	//查下luk是否存在
	nft, _ := (&LukNftService{}).GetLukNft(buy.NftId)
	if nft.ID == 0 {
		err = errors.New("找不到NFT")
		return
	}
	if nft.Address != "" {
		err = errors.New("NFT已被购买")
		return
	}
	buy.Address = strings.ToLower(buy.Address)
	buy.Price = nft.Price
	buy.NftName = nft.Name
	db := global.GVA_DB
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	nft.Address = buy.Address
	nft.UpdatedAt = time.Now()
	err = tx.Save(&nft).Error
	if err != nil {
		return
	}
	//非禁止黑名单内，如果有邀请地址，则产生推荐分红
	if !(nft.Prohibit) && buy.InviteAddress != "" && nft.Price.GreaterThan(decimal.Zero) {
		buy.InviteAddress = strings.ToLower(buy.InviteAddress)
		detailed := luk.LukRevenueDetailedNft{
			Address:   buy.InviteAddress,
			Amount:    nft.Price.Mul(decimal.NewFromFloat(0.1)),
			Type:      emun.RevenueTypeRecommend,
			TokenType: emun.USDT,
			Remark:    buy.Address + "购买所得",
		}
		err = tx.Create(&detailed).Error
		if err != nil {
			return
		}
		data := make(map[string]luk.LukRevenueNft, 0)
		data[detailed.Address] = luk.LukRevenueNft{
			Address:   detailed.Address,
			Amount:    detailed.Amount,
			Type:      detailed.Type,
			TokenType: detailed.TokenType,
		}
		//到汇总表
		err = (&LukRevenueNftService{}).UpdateLukRevenueNftBatch(tx, data)
		if err != nil {
			return
		}
		//更新个人余额
		lukUser := luk.LukUser{}
		err = tx.Where("address = ?", buy.InviteAddress).First(&lukUser).Error
		if err != nil {
			return
		}
		lukUser.Usdt = lukUser.Usdt.Add(detailed.Amount)
		lukUser.UpdatedAt = time.Now()
		if tx.Save(&lukUser).Error != nil {
			return
		}
	}
	err = tx.Create(&buy).Error
	return err
}

// GetLukNftBuy 根据id获取LukNftBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftBuyService *LukNftBuyService) GetLukNftBuy(id uint) (lukNftBuy luk.LukNftBuy, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukNftBuy).Error
	return
}

// GetLukNftBuyInfoList 分页获取LukNftBuy记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftBuyService *LukNftBuyService) GetLukNftBuyInfoList(info lukReq.LukNftBuySearch) (list []luk.LukNftBuy, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukNftBuy{})
	var lukNftBuys []luk.LukNftBuy
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db.Where("address = ?", info.Address)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukNftBuys).Error
	return lukNftBuys, total, err
}
