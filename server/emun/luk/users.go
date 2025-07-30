package luk

import enum "github.com/flipped-aurora/gin-vue-admin/server/emun"

func init() {
	enumMapInstance := enum.GetEnumMapInstance()
	enumMapInstance.Mount("user_type", new(UserType))
	enumMapInstance.Mount("partner_type", new(PartnerType))
}

//UserType 用户类型
type UserType int

const (
	UserTypeLinkage   UserType = 1 //联盟人员
	UserTypeCommunity UserType = 2 //社区人员
)

func (r UserType) HiddenKeys() []string {
	return []string{}
}

func (r UserType) KeysSort() []string {
	return []string{"联盟人员", "社区人员"}
}

func (r UserType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r UserType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		UserTypeLinkage:   "联盟人员",
		UserTypeCommunity: "社区人员",
	}
	return
}

type PartnerType int

const (
	PartnerTypeNone     PartnerType = 0 //无等级
	PartnerTypeEarly    PartnerType = 1 //青铜
	PartnerTypeMiddle   PartnerType = 2 //白银
	PartnerTypeHigh     PartnerType = 3 //黄金
	PartnerTypeOvertake PartnerType = 4 //钻石
)

func (r PartnerType) HiddenKeys() []string {
	return []string{}
}

func (r PartnerType) KeysSort() []string {
	return []string{"无等级", "青铜", "白银", "黄金", "钻石"}
}

func (r PartnerType) IsMember(key interface{}) bool {
	members := r.AllMember()
	_, exist := members[key]
	return exist
}

func (r PartnerType) AllMember() (values map[interface{}]string) {
	values = map[interface{}]string{
		PartnerTypeNone:     "无等级",
		PartnerTypeEarly:    "青铜",
		PartnerTypeMiddle:   "白银",
		PartnerTypeHigh:     "黄金",
		PartnerTypeOvertake: "钻石",
	}
	return
}
