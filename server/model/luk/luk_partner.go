// 自动生成模板LukUserPartner
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukPartner 结构体
type LukPartner struct {
	global.GVA_MODEL
	Type    emun.UserType    `json:"type" form:"type" gorm:"column:type;comment:用户类型;type:tinyint NOT NULL DEFAULT '0'"`
	Level   emun.PartnerType `json:"level" form:"level" gorm:"column:level;comment:级别;type:tinyint NOT NULL DEFAULT '0'"`
	AreaMax decimal.Decimal  `json:"areaMax" form:"areaMax" gorm:"column:area_max;comment:大区交易量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	AreaMin decimal.Decimal  `json:"areaMin" form:"areaMin" gorm:"column:area_min;comment:小区交易量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
}

// TableName LukPartner 表名
func (LukPartner) TableName() string {
	return "luk_partner"
}
