package luk

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type LukUserService struct {
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CreateLukUser 创建LukUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukUserService *LukUserService) CreateLukUser(lukUser luk.LukUser) (user luk.LukUser, errNum int, err error) {
	//如果地址存在则不需要创建用户信息
	lukUser.Address = strings.ToLower(lukUser.Address)
	global.GVA_DB.Where("address = ?", lukUser.Address).First(&user)
	if user.ID > 0 {
		return
	}
	//邀请码核实
	pUser := luk.LukUser{}
	global.GVA_DB.Where("invit = ?", lukUser.Invit).First(&pUser)
	if pUser.ID == 0 {
		err = errors.New("邀请码不存在")
		errNum = -1
		return
	}
	user.Address = lukUser.Address
	user.Pid = pUser.ID
	user.PAddress = pUser.Address
	user.Level = pUser.Level + 1
	user.Path = pUser.Path + strconv.Itoa(int(pUser.ID)) + "-"
	user.LastLoginTime = time.Now()
	user.PerformanceTime = time.Now()
	//生成自己的推荐码
	code := RandStringRunes(8)
	u := luk.LukUser{}
	global.GVA_DB.Where("invit = ?", code).First(&u)
	for u.ID > 0 {
		code = RandStringRunes(8)
		global.GVA_DB.Where("invit = ?", code).First(&u)
	}
	user.Invit = code
	err = global.GVA_DB.Create(&user).Error
	if err != nil {
		errNum = -2
	}
	return
}

func (lukUserService *LukUserService) UpdateLukUser(user luk.LukUser) (err error) {
	err = global.GVA_DB.Save(&user).Error
	return err
}

func (lukUserService *LukUserService) UpdateLinkage(address string) (err error) {
	address = strings.ToLower(address)
	user := luk.LukUser{}
	_ = global.GVA_DB.Where("address = ?", address).First(&user).Error
	if user.ID == 0 {
		err = errors.New("找不到该用户地址")
		return
	}
	if user.IsLinkage {
		err = errors.New("该用户已加入联盟白名单")
		return
	}
	user.IsLinkage = true
	err = global.GVA_DB.Save(&user).Error
	return err
}

// GetLukUser 根据id获取LukUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukUserService *LukUserService) GetLukUser(id uint) (lukUser luk.LukUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukUser).Error
	return
}

func (lukUserService *LukUserService) GetLukUserByAddress(tx *gorm.DB, address string) (lukUser luk.LukUser, err error) {
	err = tx.Where("address = ?", address).First(&lukUser).Error
	return
}

// GetLukUserInfoList 分页获取LukUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukUserService *LukUserService) GetLukUserInfoList(info lukReq.LukUserSearch) (list []luk.LukUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukUser{})
	var lukUsers []luk.LukUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db.Where("address = ?", strings.ToLower(info.Address))
	}
	if info.ID != 0 {
		db.Where("id = ?", info.ID)
	}
	if info.Islinkage != nil {
		db.Where("is_linkage = ?", *info.Islinkage)
	}
	if info.Communitylevel != nil {
		db.Where("community_level = ?", *info.Communitylevel)
	}
	if info.Linkagelevel != nil {
		db.Where("linkage_level = ?", *info.Linkagelevel)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukUsers).Error
	return lukUsers, total, err
}

//根据用户地址获取用户信息并更新登录事件
func (lukUserService *LukUserService) QueryUserByAddress(address string) (lukUser luk.LukUser, err error) {
	address = strings.ToLower(address)
	err = global.GVA_DB.Where("address = ?", address).First(&lukUser).Error
	if lukUser.ID > 0 {
		lukUser.LastLoginTime = time.Now()
		global.GVA_DB.Save(&lukUser)
	}
	return
}

func (lukUserService *LukUserService) UpdateWalletBalance(tx *gorm.DB, address string, tokenType enum.TokenType, amount decimal.Decimal, addOrsub bool) (err error) {
	data := luk.LukUser{}
	err = tx.Model(&data).Where("address = ?", address).First(&data).Error
	if err != nil || data.ID == 0 {
		return
	}
	switch tokenType {
	case enum.USDT:
		if addOrsub {
			//增加
			data.Usdt = data.Usdt.Add(amount)
		} else {
			//减少
			data.Usdt = data.Usdt.Sub(amount)
		}
	case enum.LUK:
		if addOrsub {
			//增加
			data.Luk = data.Luk.Add(amount)
		} else {
			//减少
			data.Luk = data.Luk.Sub(amount)
		}
	}
	data.UpdatedAt = time.Now()
	err = tx.Save(&data).Error
	return
}

func (lukUserService *LukUserService) QuerySubclass(address string) (lukUser []response.SubclassList, err error) {
	address = strings.ToLower(address)
	err = global.GVA_DB.Debug().Model(&luk.LukUser{}).Select("address, is_linkage, linkage_level, performance, performance_time").Where("p_address = ?", address).Find(&lukUser).Error
	return
}

//实现获取推荐树全部用户地址
func (lukUserService *LukUserService) GetTreeAddress(address string) (treeAddress []string) {
	tabel := (&luk.LukUser{}).TableName()
	fathers := []string{address}
	treeAddress = make([]string, 0)
	for len(fathers) > 0 {
		data := []luk.LukUser{}
		_ = global.GVA_DB.Debug().Table(tabel).Select("address").Where("p_address in ? ", fathers).Find(&data).Error
		fathers = make([]string, 0)
		for _, item := range data {
			fathers = append(fathers, item.Address)
			if !config.InStringArray(item.Address, treeAddress) {
				treeAddress = append(treeAddress, item.Address)
			}
		}
	}
	return
}

//实现获取推荐树全部用户地址 20代
func (lukUserService *LukUserService) GetTreeAddressByGeneration(address string, generation int) (treeAddress []string) {
	tabel := (&luk.LukUser{}).TableName()
	fathers := []string{address}
	treeAddress = make([]string, 0)
	for generation > 0 {
		if len(fathers) == 0 {
			break
		}
		data := []luk.LukUser{}
		_ = global.GVA_DB.Debug().Table(tabel).Select("address").Where("p_address in ? ", fathers).Find(&data).Error
		fathers = make([]string, 0)
		for _, item := range data {
			fathers = append(fathers, item.Address)
			if !config.InStringArray(item.Address, treeAddress) {
				treeAddress = append(treeAddress, item.Address)
			}
		}
		generation--
	}
	return
}

func (lukUserService *LukUserService) UpdateBatch(tx *gorm.DB, list map[string]luk.LukUser) (err error) {
	time.Sleep(3 * time.Second)
	for _, item := range list {
		data := luk.LukUser{}
		_ = tx.Where("address = ?", item.Address).First(&data).Error
		if data.ID == 0 {
			continue
		}
		//增加日志
		log := luk.LukUserWalletLog{
			Address: data.Address,
			OldLuk:  data.Luk,
			OldUsdt: data.Usdt,
		}

		data.Usdt = data.Usdt.Add(item.Usdt)
		data.Luk = data.Luk.Add(item.Luk)
		data.UpdatedAt = time.Now()
		err = tx.Save(&data).Error

		if err != nil {
			log.ErrMsg = err.Error()
		}
		log.NewUsdt = data.Usdt
		log.NewLuk = data.Luk
		err1 := tx.Create(&log).Error
		fmt.Println("=====", err1)

		if err != nil {
			return
		}

	}
	return
}

func (lukUserService *LukUserService) AddressAllBranch(address string) (list []string) {
	sub := []response.SubclassList{}
	err := global.GVA_DB.Debug().Model(&luk.LukUser{}).Select("address").Where("p_address = ?", address).Find(&sub).Error
	if err != nil || len(sub) == 0 {
		return
	}
	list = make([]string, 0)
	for _, item := range sub {
		list = append(list, item.Address)
		addresss := lukUserService.GetTreeAddressByGeneration(item.Address, 20)
		list = append(list, addresss...)
	}
	return
}

//统计联盟白名单用户数据
func (lukUserService *LukUserService) QueryIsLinkage() (data map[enum.PartnerType]decimal.Decimal) {
	sub := make([]response.LinkageList, 0)
	err := global.GVA_DB.Debug().Model(&luk.LukUser{}).Select("linkage_level, linkage_level as proportion, count(1) as count").Where("is_linkage = 1 and linkage_level > 0").Group("linkage_level").Find(&sub).Error
	if err != nil {
		return
	}

	if len(sub) == 0 {
		return
	}
	//计算出总额占比
	proportionTotal := 0
	for _, item := range sub {
		proportionTotal = proportionTotal + item.Proportion
	}
	data = make(map[enum.PartnerType]decimal.Decimal, 0)
	for _, item := range sub {
		//用自己的占比除总占比
		pro := decimal.NewFromInt(int64(item.Proportion)).Div(decimal.NewFromInt(int64(proportionTotal)))
		data[item.LinkageLevel] = pro.Div(decimal.NewFromInt(item.Count))
	}
	return
}

//全部用户地址
func (lukUserService *LukUserService) CountAll() (count int64) {
	_ = global.GVA_DB.Debug().Model(&luk.LukUser{}).Count(&count).Error
	return
}

//获取父级直推地址
func (lukUserService *LukUserService) QueryDirectPush(address string) (addressMap []string) {
	addressMap = make([]string, 0)
	lukUser := []luk.LukUser{}
	address = strings.ToLower(address)
	err := global.GVA_DB.Debug().Model(&luk.LukUser{}).Select("address").Where("p_address = ?", address).Find(&lukUser).Error
	if err == nil {
		for _, item := range lukUser {
			addressMap = append(addressMap, item.Address)
		}
	}
	return
}

//获取直推地址大小区
func (lukUserService *LukUserService) QueryArea(address string) (areaMax, areaMin float64) {
	performance := make([]float64, 0)
	lukUser := []luk.LukUser{}
	address = strings.ToLower(address)
	err := global.GVA_DB.Debug().Model(&luk.LukUser{}).Select("performance").Where("p_address = ?", address).Find(&lukUser).Error
	if err == nil {
		for _, item := range lukUser {
			performance = append(performance, item.Performance.InexactFloat64())
		}
		//金额排序，取最大的，其他的全部加起来
		if len(performance) > 0 {
			sort.Float64s(performance)
			sort.Sort(sort.Reverse(sort.Float64Slice(performance)))
			areaMax = performance[0]
			performance = performance[1:]
			for _, item := range performance {
				areaMin = areaMin + item
			}
		}
	}
	return
}
