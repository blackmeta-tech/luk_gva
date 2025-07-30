// 自动生成模板LukBanner
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukBanner 结构体
type LukBanner struct {
	global.GVA_MODEL
	Pic    string `json:"pic" form:"pic" gorm:"column:pic;comment:图片路径;size:500;"`
	Link   string `json:"link" form:"link" gorm:"column:link;comment:链接;size:500;"`
	Status *int   `json:"status" form:"status" gorm:"column:status;comment:是否停用(1启用，2停用);size:10;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`
	PicKey string `json:"picKey" form:"picKey" gorm:"column:pic_key;comment:图片key;size:64;"`
}

// TableName LukBanner 表名
func (LukBanner) TableName() string {
	return "luk_banner"
}
