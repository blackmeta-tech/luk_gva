// 自动生成模板LukInformationType
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukInformationType 结构体
// 如果含有time.Time 请自行import time包
type LukInformationType struct {
	global.GVA_MODEL
	NameCh  string `json:"nameCh" form:"nameCh" gorm:"column:name_ch;comment:分类(中文);size:64;"`
	NameEn  string `json:"nameEn" form:"nameEn" gorm:"column:name_en;comment:分类(英文);size:64;"`
	SortNum *int   `json:"sortNum" form:"sortNum" gorm:"column:sort_num;comment:排序序号;size:10;NOT NULL DEFAULT '0'"`
	Status  *int   `json:"status" form:"status" gorm:"column:status;comment:是否停用(1启用，2停用);size:10;NOT NULL DEFAULT '0'"`
}

// TableName LukInformationType 表名
func (LukInformationType) TableName() string {
	return "luk_information_type"
}
