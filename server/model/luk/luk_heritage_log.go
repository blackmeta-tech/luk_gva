package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type LukHeritageLog struct {
	global.GVA_MODEL
	Date string `json:"date" orm:"date" gorm:"column:date;comment:日期;type:date"`
}

// TableName DxbaKline 表名
func (LukHeritageLog) TableName() string {
	return "luk_heritage_log"
}

func (LukHeritageLog) QueryByTime(_date string) (info LukHeritageLog, err error) {
	err = global.GVA_DB.Debug().Where("date = ? ", _date).First(&info).Error
	return
}
