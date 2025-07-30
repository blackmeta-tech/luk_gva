package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type LukDividendsLog struct {
	global.GVA_MODEL
	Date     string            `json:"date" orm:"date" gorm:"column:date;comment:日期;type:date"`
	Category emun.DividendType `json:"category" form:"category" gorm:"column:category;comment:分红大类类型,定义在代码里;type:tinyint NOT NULL DEFAULT '0'; "`
}

// TableName DxbaKline 表名
func (LukDividendsLog) TableName() string {
	return "luk_dividends_log"
}

func (LukDividendsLog) QueryByTime(_date string, category emun.DividendType) (info LukDividendsLog, err error) {
	err = global.GVA_DB.Debug().Where("date = ? and category = ?", _date, category).First(&info).Error
	return
}
