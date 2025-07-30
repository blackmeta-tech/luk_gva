package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// LukUserAddress 结构体
type LukUserAddress struct {
	global.GVA_MODEL
	Address  string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:200;index:idx_name_user_address"`
	IsCharge int             `json:"isCharge" form:"isCharge" gorm:"column:is_charge;comment:零手续费白名单 0否 1是;type:tinyint NOT NULL DEFAULT '0'"`
	Lp       decimal.Decimal `json:"lp" form:"lp" gorm:"column:lp;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:流动的量;"`
	NoLp     int             `json:"noLp" form:"noLp" gorm:"column:no_lp;type:int(11) NOT NULL DEFAULT '0';comment:LP分红标记 1标识停止分红;"`
	Remark   string          `json:"remark" form:"remark" gorm:"column:remark;comment:用户地址;size:200;"`
}

// TableName LukUserAddress 表名
func (LukUserAddress) TableName() string {
	return "luk_user_address"
}

//修改添加或移除流动性的对应值
func (d *LukUserAddress) UpdateDataLp(address string, lp decimal.Decimal, _type bool) (err error) {
	address = strings.ToLower(address)
	info := LukUserAddress{}
	global.GVA_DB.Where("address = ?", address).First(&info)
	if _type {
		//添加流动性
		info.Lp = info.Lp.Add(lp)
	} else {
		//移除流动性时可能会出现移除的多，直接致为零
		if Sublp := info.Lp.Sub(lp); Sublp.LessThanOrEqual(decimal.Zero) {
			info.Lp = decimal.Zero
		} else {
			info.Lp = Sublp
		}
	}
	if info.ID > 0 {
		info.UpdatedAt = time.Now()
		err = global.GVA_DB.Save(&info).Error
	} else {
		info.Address = address
		err = global.GVA_DB.Create(&info).Error
	}
	return
}
