package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
)

type LukDappService struct {
}

//获取首页交易记录，组装NFT以及套餐购买
func (l *LukDappService) TransactionRecord() (datas []response.TransactionRecordList) {
	stamps := make([]int, 0)
	datas = make([]response.TransactionRecordList, 0)
	list := make(map[int]response.TransactionRecordList, 0)
	//取购买套餐前面8条记录
	var lukComboBuys []luk.LukComboBuy
	global.GVA_DB.Debug().Model(&luk.LukComboBuy{}).Limit(8).Order("id desc").Find(&lukComboBuys)
	for _, item := range lukComboBuys {
		item.CreatedAt.Unix()
		stamp := int(item.CreatedAt.Unix())
		info := response.TransactionRecordList{
			CreatedAt: item.CreatedAt,
			Content:   strconv.Itoa(item.Days) + "天套餐",
			Address:   item.Address,
		}
		list[stamp] = info
		stamps = append(stamps, stamp)
	}
	var lukNftBuys []luk.LukNftBuy
	global.GVA_DB.Debug().Model(&luk.LukNftBuy{}).Limit(8).Order("id desc").Find(&lukNftBuys)
	for _, item := range lukNftBuys {
		stamp := int(item.CreatedAt.Unix())
		info := response.TransactionRecordList{
			CreatedAt: item.CreatedAt,
			Content:   item.NftName,
			Address:   item.Address,
		}
		list[stamp] = info
		stamps = append(stamps, stamp)
	}
	sort.Ints(stamps)
	sort.Sort(sort.Reverse(sort.IntSlice(stamps)))

	for _, k := range stamps {
		datas = append(datas, list[k])
	}
	return
}

//DEX操作记录
func (l *LukDappService) _blockScanRecordQuery(req lukReq.BlockScanRecordSearch) *gorm.DB {
	// 创建db
	db := global.GVA_DB.Debug().Model(&luk.LukBlockScanRecord{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if req.Type > 0 {
		db.Where("type = ? ", req.Type)
	}
	if req.From != "" {
		db.Where("`from` = ?", req.From)
	}
	return db
}
func (l *LukDappService) BlockScanRecordList(req lukReq.BlockScanRecordSearch) (datas []response.BlockScanRecord, summary decimal.Decimal, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := l._blockScanRecordQuery(req)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("time_stamp desc").Find(&datas).Error

	r := luk.LukBlockScanRecord{}
	err = l._blockScanRecordQuery(req).Select("sum(lp_amount) as lp_amount, sum(luk_amount) as luk_amount, sum(usdt_amount) as usdt_amount").First(&r).Error
	if err == nil {
		if req.Type == emun.MethodTypeSwap {
			summary = r.UsdtAmount
		} else if req.Type == emun.MethodTypeSell {
			summary = r.LukAmount
		} else {
			summary = r.LpAmount
		}
	}
	return
}

func (l *LukDappService) QueryHeritage(address string) (datas response.HeritageInfo) {
	if address == "" {
		return
	}
	address = strings.ToLower(address)
	infoLp := luk.LukHeritageLp{}
	global.GVA_DB.Where("address = ?", address).First(&infoLp)
	if infoLp.ID > 0 {
		datas.LpBalance = infoLp.Balance
		datas.LpPercentage = global.GVA_LUK_CONFIG.RatioLp
		datas.LpFreed = infoLp.Balance.Sub(infoLp.Remaining)
	}
	infoMetaverseOld := luk.LukHeritageMetaverseOld{}
	global.GVA_DB.Where("address = ?", address).First(&infoMetaverseOld)
	if infoMetaverseOld.ID > 0 {
		datas.MetaverseOldBalance = infoMetaverseOld.Balance
		datas.MetaverseOldPercentage = global.GVA_LUK_CONFIG.RatioMetaverseOld
		datas.MetaverseOldFreed = infoMetaverseOld.Balance.Sub(infoMetaverseOld.Remaining)
	}
	infoMetaverse := luk.LukHeritageMetaverse{}
	global.GVA_DB.Where("address = ?", address).First(&infoMetaverse)
	if infoMetaverse.ID > 0 {
		datas.MetaverseBalance = infoMetaverse.Balance
		datas.MetaversePercentage = global.GVA_LUK_CONFIG.RatioMetaverse1
		if infoMetaverse.LpUsdt.GreaterThan(decimal.Zero) {
			datas.MetaversePercentage = datas.MetaversePercentage.Add(global.GVA_LUK_CONFIG.RatioMetaverse3)
		}
		if infoMetaverse.ComboUsdt.GreaterThan(decimal.Zero) {
			datas.MetaversePercentage = datas.MetaversePercentage.Add(global.GVA_LUK_CONFIG.RatioMetaverse2)
		}
		datas.MetaverseFreed = infoMetaverse.Balance.Sub(infoMetaverse.Remaining)
	}
	if datas.LpBalance.GreaterThan(decimal.Zero) || datas.MetaverseBalance.GreaterThan(decimal.Zero) || datas.MetaverseOldBalance.GreaterThan(decimal.Zero) {
		datas.State = true
	}
	return
}
