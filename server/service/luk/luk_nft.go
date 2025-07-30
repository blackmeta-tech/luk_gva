package luk

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/token"
	"github.com/shopspring/decimal"
)

type LukNftService struct {
}

// CreateLukNft 创建LukNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftService *LukNftService) CreateLukNft(lukNft luk.LukNft) (err error) {
	nft := luk.LukNft{}
	global.GVA_DB.Debug().Where("name = ?", lukNft.Name).First(&nft)
	if nft.ID > 0 {
		err = errors.New("名称已存在！")
		return
	}
	err = global.GVA_DB.Create(&lukNft).Error
	return err
}

// UpdateLukNft 更新LukNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftService *LukNftService) UpdateLukNft(lukNft luk.LukNft) (err error) {
	err = global.GVA_DB.Save(&lukNft).Error
	return err
}

func (lukNftService *LukNftService) UpdateLukNftBlack(id uint) (err error) {
	nft := luk.LukNft{}
	global.GVA_DB.Where("id = ?", id).First(&nft)
	if nft.ID == 0 {
		err = errors.New("找不到NFT编号")
		return
	}
	if nft.Blacklist {
		err = errors.New("该编号已加入黑名单")
		return
	}
	nft.Blacklist = true
	err = global.GVA_DB.Save(&nft).Error
	return err
}

// GetLukNft 根据id获取LukNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftService *LukNftService) GetLukNft(id uint) (lukNft luk.LukNft, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukNft).Error
	return
}

// GetLukNftInfoList 分页获取LukNft记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftService *LukNftService) GetLukNftInfoList(info lukReq.LukNftSearch) (list []response.LukNftList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table((&luk.LukNft{}).TableName()+" n").Select("n.*, user.area_max, user.area_min").
		Joins("LEFT JOIN luk_user user ON n.address = user.address", "user")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("n.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db.Where("n.name like ?", "%"+info.Name+"%")
	}
	if info.Address != "" {
		db.Where("n.address = ?", info.Address)
	}
	if info.ID != 0 {
		db.Where("n.id = ?", info.ID)
	}
	if info.Statu != nil {
		db.Where("n.status = ?", *info.Statu)
	}
	if info.IsBlacklist != nil {
		db.Where("n.blacklist = ?", *info.IsBlacklist)
	}
	if info.ByOwn != "" {
		db.Where("(n.address = ? or n.address='')", info.ByOwn)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("n.id desc").Find(&list).Error
	return
}

//购买NFT前置加密信息
func (lukNftService *LukNftService) BuyLukNftAes(pass string) (pass64 string, err error) {
	key := "aidhh451bnffg2df"
	//获取随机码拼接进去
	r, err := (&token.NftTokenService{}).GetRandom()
	if err != nil {
		return
	}
	pass = pass + r
	xpass, err := config.AesEncrypt([]byte(pass), []byte(key))
	if err != nil {
		return
	}

	pass64 = base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		return
	}

	tpass, err := config.AesDecrypt(bytesPass, []byte(key))
	if err != nil {
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
	return
}

//汇总数据
func (lukNftService *LukNftService) Summary() (circulation, hold, history int64) {
	table := (&luk.LukNft{}).TableName()
	//统计全部卖出
	global.GVA_DB.Debug().Table(table).Where("address != ''").Count(&circulation)
	//统计当前持有人量
	global.GVA_DB.Debug().Table(table).Where("address != ''").Group("address").Count(&hold)
	//统计历史持有人量
	address := make([]string, 0)
	list := []luk.LukNft{}
	err := global.GVA_DB.Debug().Table(table).Where("address != ''").Find(&list).Error
	if err != nil {
		return
	}
	for _, i := range list {
		if !config.InStringArray(i.Address, address) {
			address = append(address, i.Address)
		}
	}
	transfer := []luk.LukNftTransfer{}
	global.GVA_DB.Debug().Find(&transfer)
	for _, i := range transfer {
		if !config.InStringArray(i.To, address) {
			address = append(address, i.To)
		}
		if !config.InStringArray(i.From, address) {
			address = append(address, i.From)
		}
	}
	history = int64(len(address))
	return
}

//统计用户拥有的NFT以及算出占比
func (lukNftService *LukNftService) QueryNftByNum() (data map[string]decimal.Decimal) {
	data = make(map[string]decimal.Decimal, 0)
	type LukNftByAddress struct {
		Address string `json:"address"`
		Summary int64  `json:"summary"`
	}
	list := []LukNftByAddress{}
	db := global.GVA_DB.Debug().Table((&luk.LukNft{}).TableName()).Select("address, count(1) as summary").Where("address != '' and prohibit = 0")
	err := db.Group("address").Find(&list).Error
	if err != nil {
		return
	}
	var count int64
	global.GVA_DB.Debug().Table((&luk.LukNft{}).TableName()).Where("address != ''").Count(&count)
	for _, item := range list {
		data[item.Address] = decimal.NewFromInt(item.Summary).Div(decimal.NewFromInt(count))
	}
	return
}
