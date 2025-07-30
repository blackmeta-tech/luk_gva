package luk

import enum "github.com/flipped-aurora/gin-vue-admin/server/emun"

func init() {
	enumMapInstance := enum.GetEnumMapInstance()
	enumMapInstance.Mount("token_method_name", new(TokenMethodName))
}

//TokenMethodName 合约事件类型类型
type TokenMethodName string

const (
	TokenMethodSwapTokensForExactTokens                              TokenMethodName = "swapTokensForExactTokens"                              //兑换1
	TokenMethodSwapExactTokensForTokens                              TokenMethodName = "swapExactTokensForTokens"                              //兑换2
	TokenMethodSwapExactTokensForTokensSupportingFeeOnTransferTokens TokenMethodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens" //兑换3
	TokenMethodAddLiquidity                                          TokenMethodName = "addLiquidity"                                          //添加流动性
	TokenMethodRemoveLiquidity                                       TokenMethodName = "removeLiquidity"                                       //移除流动性1
	TokenMethodRemoveLiquidityWithPermit                             TokenMethodName = "removeLiquidityWithPermit"                             //移除流动性2
)

func (r TokenMethodName) HiddenKeys() []string {
	return []string{}
}

func (r TokenMethodName) KeysSort() []string {
	return []string{"兑换事件1", "兑换事件2", "兑换事件3", "添加流动性", "移除流动性事件1", "移除流动性事件2"}
}

func (r TokenMethodName) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r TokenMethodName) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		TokenMethodSwapTokensForExactTokens:                              "兑换事件1",
		TokenMethodSwapExactTokensForTokens:                              "兑换事件2",
		TokenMethodSwapExactTokensForTokensSupportingFeeOnTransferTokens: "兑换事件3",
		TokenMethodAddLiquidity:                                          "添加流动性",
		TokenMethodRemoveLiquidity:                                       "移除流动性事件1",
		TokenMethodRemoveLiquidityWithPermit:                             "移除流动性事件2",
	}
	return
}

type MethodType int

const (
	MethodTypeSwap            MethodType = 1 //兑换
	MethodTypeSell            MethodType = 2 //卖出
	MethodTypeAddLiquidity    MethodType = 3 //添加流动性
	MethodTypeRemoveLiquidity MethodType = 4 //移除流动性
)

func (r MethodType) HiddenKeys() []string {
	return []string{}
}

func (r MethodType) KeysSort() []string {
	return []string{"兑换", "卖出", "添加流动性", "移除流动性"}
}

func (r MethodType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r MethodType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		MethodTypeSwap:            "兑换",
		MethodTypeSell:            "卖出",
		MethodTypeAddLiquidity:    "添加流动性",
		MethodTypeRemoveLiquidity: "移除流动性",
	}
	return
}
