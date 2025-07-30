package luk

import enum "github.com/flipped-aurora/gin-vue-admin/server/emun"

func init() {
	enumMapInstance := enum.GetEnumMapInstance()
	enumMapInstance.Mount("wallet_type", new(WalletType))
	enumMapInstance.Mount("reflow_type", new(ReflowType))
}

type WalletType string

const (
	WalletTypeLuk         WalletType = "luk"         //09-合约部署钱包（NFT， LUK）
	WalletTypeComplex     WalletType = "complex"     //01-20w综合钱包
	WalletTypeContract    WalletType = "contract"    //02-80w合约钱包
	WalletTypeFormalities WalletType = "formalities" //03-DEX手续费钱包
	WalletTypeNft         WalletType = "nft"         //04-NFT钱包
	WalletTypePledge      WalletType = "pledge"      //05-套餐质押钱包
	WalletTypeDestroy     WalletType = "destroy"     //06-销毁钱包
	WalletTypeReflow      WalletType = "reflow"      //07-回流底池钱包
	WalletTypeWithdraw    WalletType = "withdraw"    //08-综合提现钱包
	WalletTypeFee         WalletType = "fee"         //10-luk转账手续费钱包
)

func (r WalletType) HiddenKeys() []string {
	return []string{}
}

func (r WalletType) KeysSort() []string {
	return []string{"01-20w综合钱包", "02-80w合约钱包", "03-DEX手续费钱包", "04-NFT钱包", "05-套餐质押钱包", "06-销毁钱包", "07-回流底池钱包", "08-综合提现钱包", "09-合约部署钱包（NFT,LUK）", "10-luk转账手续费钱包"}
}

func (r WalletType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r WalletType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		WalletTypeLuk:         "09-合约部署钱包（NFT， LUK）",
		WalletTypeComplex:     "01-20w综合钱包",
		WalletTypeContract:    "02-80w合约钱包",
		WalletTypeFormalities: "03-DEX手续费钱包",
		WalletTypeNft:         "04-NFT钱包",
		WalletTypePledge:      "05-套餐质押钱包",
		WalletTypeDestroy:     "06-销毁钱包",
		WalletTypeReflow:      "07-回流底池钱包",
		WalletTypeWithdraw:    "08-综合提现钱包",
		WalletTypeFee:         "10-luk转账手续费钱包",
	}
	return
}

type ReflowType int

const (
	ReflowTypeAlliance  ReflowType = 1 //8%联盟分红
	ReflowTypeStcPot    ReflowType = 2 //3%STC底池
	ReflowTypeStcNode   ReflowType = 3 //3%STC节点分红
	ReflowTypeOperate   ReflowType = 4 //3%运营钱包
	ReflowTypeNft       ReflowType = 5 //1%NFT加权分红
	ReflowTypeFunding   ReflowType = 6 //1%联盟经费
	ReflowTypeAllowance ReflowType = 7 //1%联盟运营岗位津贴
)

func (r ReflowType) HiddenKeys() []string {
	return []string{}
}

func (r ReflowType) KeysSort() []string {
	return []string{"联盟分红", "STC底池", "STC节点分红", "运营钱包", "NFT加权分红", "联盟经费", "联盟运营岗位津贴"}
}

func (r ReflowType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r ReflowType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		ReflowTypeAlliance:  "联盟分红",
		ReflowTypeStcPot:    "STC底池",
		ReflowTypeStcNode:   "STC节点分红",
		ReflowTypeOperate:   "运营钱包",
		ReflowTypeNft:       "NFT加权分红",
		ReflowTypeFunding:   "联盟经费",
		ReflowTypeAllowance: "联盟运营岗位津贴",
	}
	return
}

type ReflowLpType int

const (
	ReflowLpTypePool    ReflowLpType = 1 //1%回流底池
	ReflowLpTypeOperate ReflowLpType = 2 //1%运营钱包
)

func (r ReflowLpType) HiddenKeys() []string {
	return []string{}
}

func (r ReflowLpType) KeysSort() []string {
	return []string{"回流底池", "运营钱包"}
}

func (r ReflowLpType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r ReflowLpType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		ReflowLpTypePool:    "回流底池",
		ReflowLpTypeOperate: "运营钱包",
	}
	return
}
