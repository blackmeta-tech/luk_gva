// 自动生成模板LukUser
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

// LukUser 结构体
type LukUser struct {
	global.GVA_MODEL
	Address         string           `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:255;index:idx_name_address"`
	Pid             uint             `json:"pid" form:"pid" gorm:"column:pid;comment:上级ID;type:int(11) NOT NULL DEFAULT '0';index:idx_name_pid"`
	PAddress        string           `json:"pAddress" form:"pAddress" gorm:"column:p_address;comment:上级地址;size:255;"`
	Path            string           `json:"path" form:"path" gorm:"column:path;comment:上下级关系;"`
	Level           int              `json:"level" form:"level" gorm:"column:level;comment:层数;type:int(11) NOT NULL DEFAULT '0';"`
	Invit           string           `json:"invit" form:"invit" gorm:"column:invit;comment:推荐码;size:10;"`
	LastLoginTime   time.Time        `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:最后登录时间;"`
	PerformanceTime time.Time        `json:"performanceTime" form:"performanceTime" gorm:"column:performance_time;comment:业绩更新时间;"`
	Usdt            decimal.Decimal  `json:"usdt" form:"usdt" gorm:"column:usdt;comment:USDT数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Luk             decimal.Decimal  `json:"luk" form:"luk" gorm:"column:luk;comment:LUK数量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	AreaMax         decimal.Decimal  `json:"areaMax" form:"areaMax" gorm:"column:area_max;comment:大区交易量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	AreaMin         decimal.Decimal  `json:"areaMin" form:"areaMin" gorm:"column:area_min;comment:小区交易量;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	Performance     decimal.Decimal  `json:"performance" form:"performance" gorm:"column:performance;comment:含自己20代业绩;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';"`
	CommunityLevel  emun.PartnerType `json:"communityLevel" form:"communityLevel" gorm:"community_level;comment:社区等级;type:int(11) NOT NULL DEFAULT '0';"`
	IsLinkage       bool             `json:"isLinkage" form:"isLinkage" gorm:"is_linkage;comment:联盟白名单;type:tinyint(2) NOT NULL DEFAULT '0';"`
	LinkageLevel    emun.PartnerType `json:"linkageLevel" form:"linkageLevel" gorm:"linkage_level;comment:联盟等级;type:int(11) NOT NULL DEFAULT '0';"`
}

// TableName LukUser 表名
func (LukUser) TableName() string {
	return "luk_user"
}
