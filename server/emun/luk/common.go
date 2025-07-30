package luk

import enum "github.com/flipped-aurora/gin-vue-admin/server/emun"

func init() {
	enumMapInstance := enum.GetEnumMapInstance()
	enumMapInstance.Mount("token_type", new(TokenType))
	enumMapInstance.Mount("revenue_type", new(RevenueType))
	enumMapInstance.Mount("dividend_type", new(DividendType))
	enumMapInstance.Mount("status_type", new(StatusType))
	enumMapInstance.Mount("revenue_type_nft", new(RevenueTypeNft))
	enumMapInstance.Mount("heritage_type", new(HeritageType))
}

//TokenType 币种类型
type TokenType int

const (
	NFT TokenType = iota + 1
	USDT
	LUK
)

func (r TokenType) HiddenKeys() []string {
	return []string{}
}

func (r TokenType) KeysSort() []string {
	return []string{"USDT", "NFT", "LUK"}
}

func (r TokenType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r TokenType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		NFT:  "NFT",
		USDT: "USDT",
		LUK:  "LUK",
	}
	return
}

//RevenueType 分红类型
type RevenueType int

const (
	RevenueTypeCombo7            RevenueType = 2  //套餐质押 7天
	RevenueTypeCombo15           RevenueType = 3  //套餐质押 15天
	RevenueTypeCombo30           RevenueType = 4  //套餐质押 30天
	RevenueTypeCombo60           RevenueType = 5  //套餐质押 60天
	RevenueTypeComboRecommend    RevenueType = 6  //套餐推荐
	RevenueTypeLinkage1          RevenueType = 7  //联盟分红-青铜
	RevenueTypeLinkage2          RevenueType = 8  //联盟分红-白银
	RevenueTypeLinkage3          RevenueType = 9  //联盟分红-黄金
	RevenueTypeLinkage4          RevenueType = 10 //联盟分红-钻石
	RevenueTypeCommunity1        RevenueType = 11 //社区分红-青铜
	RevenueTypeCommunity2        RevenueType = 12 //社区分红-白银
	RevenueTypeCommunity3        RevenueType = 13 //社区分红-黄金
	RevenueTypeCommunity4        RevenueType = 14 //社区分红-钻石
	RevenueTypeLpSwap            RevenueType = 18 //LP-兑换买入
	RevenueTypeLpSell            RevenueType = 19 //LP-兑换卖出
	RevenueTypeLpRemoveLiquidity RevenueType = 20 //LP-移除流动性
	RevenueTypeLpRebate          RevenueType = 23 //LP-返佣
	RevenueTypeAchievement       RevenueType = 24 //业绩返佣
)

func (r RevenueType) HiddenKeys() []string {
	return []string{}
}

func (r RevenueType) KeysSort() []string {
	return []string{"套餐质押分红-7天", "套餐质押分红-15天", "套餐质押分红-30天", "套餐质押分红-60天", "套餐推荐分红",
		"联盟分红-青铜", "联盟分红-白银", "联盟分红-黄金", "联盟分红-钻石", "社区分红-青铜", "社区分红-白银", "社区分红-黄金", "社区分红-钻石",
		"LP-兑换买入", "LP-兑换卖出", "LP-移除流动性", "LP-返佣", "业绩返佣",
	}
}

func (r RevenueType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r RevenueType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		RevenueTypeCombo7:            "套餐质押分红-7天",
		RevenueTypeCombo15:           "套餐质押分红-15天",
		RevenueTypeCombo30:           "套餐质押分红-30天",
		RevenueTypeCombo60:           "套餐质押分红-60天",
		RevenueTypeComboRecommend:    "套餐推荐分红",
		RevenueTypeLinkage1:          "联盟分红-青铜",
		RevenueTypeLinkage2:          "联盟分红-白银",
		RevenueTypeLinkage3:          "联盟分红-黄金",
		RevenueTypeLinkage4:          "联盟分红-钻石",
		RevenueTypeCommunity1:        "社区分红-青铜",
		RevenueTypeCommunity2:        "社区分红-白银",
		RevenueTypeCommunity3:        "社区分红-黄金",
		RevenueTypeCommunity4:        "社区分红-钻石",
		RevenueTypeLpSwap:            "LP-兑换买入",
		RevenueTypeLpSell:            "LP-兑换卖出",
		RevenueTypeLpRemoveLiquidity: "LP-移除流动性",
		RevenueTypeLpRebate:          "LP-返佣",
		RevenueTypeAchievement:       "业绩返佣",
	}
	return
}

//RevenueTypeNft Nft分红类型
type RevenueTypeNft int

const (
	RevenueTypeRecommend          RevenueTypeNft = 1  //NFT 推荐分红
	RevenueTypeNftSwap            RevenueTypeNft = 15 //NFT-兑换买入
	RevenueTypeNftSell            RevenueTypeNft = 16 //NFT-兑换卖出
	RevenueTypeNftRemoveLiquidity RevenueTypeNft = 17 //NFT-移除流动性
	RevenueTypeNftWeighted        RevenueTypeNft = 21 //NFT-加权套餐
	RevenueTypeNftValue           RevenueTypeNft = 22 //NFT-1.5kU价值
)

func (r RevenueTypeNft) HiddenKeys() []string {
	return []string{}
}

func (r RevenueTypeNft) KeysSort() []string {
	return []string{"NFT-推荐分红", "NFT-兑换买入", "NFT-兑换卖出", "NFT-移除流动性", "NFT-加权套餐", "NFT-1.5kU价值"}
}

func (r RevenueTypeNft) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r RevenueTypeNft) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		RevenueTypeRecommend:          "NFT-推荐分红",
		RevenueTypeNftSwap:            "NFT-兑换买入",
		RevenueTypeNftSell:            "NFT-兑换卖出",
		RevenueTypeNftRemoveLiquidity: "NFT-移除流动性",
		RevenueTypeNftWeighted:        "NFT-加权套餐",
		RevenueTypeNftValue:           "NFT-1.5kU价值",
	}
	return
}

//DividendType 分红大类
type DividendType int

const (
	DividendTypeNft            DividendType = 1 //NFT
	DividendTypeCombo          DividendType = 2 //套餐质押
	DividendTypeComboRecommend DividendType = 3 //套餐推荐
	DividendTypeLinkage        DividendType = 4 //联盟分红
	DividendTypeCommunity      DividendType = 5 //社区分红
	DividendTypeLp             DividendType = 6 //LP分红
	DividendTypeRebate         DividendType = 7 //业绩返佣
)

func (r DividendType) HiddenKeys() []string {
	return []string{}
}

func (r DividendType) KeysSort() []string {
	return []string{"NFT分红", "套餐质押分红", "套餐推荐分红", "联盟分红", "社区分红", "LP分红", "业绩返佣"}
}

func (r DividendType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r DividendType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		DividendTypeNft:            "NFT分红",
		DividendTypeCombo:          "套餐质押分红",
		DividendTypeComboRecommend: "套餐推荐分红",
		DividendTypeLinkage:        "联盟分红",
		DividendTypeCommunity:      "社区分红",
		DividendTypeLp:             "LP分红",
		DividendTypeRebate:         "业绩返佣",
	}
	return
}

//StatusType 状态类型
type StatusType int

const (
	StatusTypeWait    StatusType = 0  //等待
	StatusTypeSuccess StatusType = 1  //成功
	StatusTypeFail    StatusType = -1 //失败
)

func (r StatusType) HiddenKeys() []string {
	return []string{}
}

func (r StatusType) KeysSort() []string {
	return []string{"等待", "成功", "失败"}
}

func (r StatusType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r StatusType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		StatusTypeWait:    "等待",
		StatusTypeSuccess: "成功",
		StatusTypeFail:    "失败",
	}
	return
}

//HeritageType 移产类型
type HeritageType int

const (
	HeritageTypeMetaverse    HeritageType = 1 //新元宇宙银行
	HeritageTypeMetaverseOld HeritageType = 2 //老元宇宙银行
	HeritageTypeLp           HeritageType = 3 //LP矿机
)

func (r HeritageType) HiddenKeys() []string {
	return []string{}
}

func (r HeritageType) KeysSort() []string {
	return []string{"新元宇宙银行", "老元宇宙银行", "LP矿机"}
}

func (r HeritageType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r HeritageType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		HeritageTypeMetaverse:    "新元宇宙银行",
		HeritageTypeMetaverseOld: "老元宇宙银行",
		HeritageTypeLp:           "LP矿机",
	}
	return
}
