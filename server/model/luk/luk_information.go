// 自动生成模板LukInformation
package luk

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukInformation 结构体
// 如果含有time.Time 请自行import time包
type LukInformation struct {
	global.GVA_MODEL
	ArticleCh string     `json:"articleCh" form:"articleCh" gorm:"column:article_ch;comment:内容(中文);size:4294967295;"`
	ArticleEn string     `json:"articleEn" form:"articleEn" gorm:"column:article_en;comment:内容(英文);size:4294967295;"`
	Status    *int       `json:"status" form:"status" gorm:"column:status;comment:是否停用(1启用，2停用);size:10;NOT NULL DEFAULT '0'"`
	Timeline  *time.Time `json:"timeline" form:"timeline" gorm:"column:timeline;comment:时间;"`
	TitleCh   string     `json:"titleCh" form:"titleCh" gorm:"column:title_ch;comment:标题(中文);size:200;"`
	TitleEn   string     `json:"titleEn" form:"titleEn" gorm:"column:title_en;comment:标题(英文);size:200;"`
	TypeId    *int       `json:"typeId" form:"typeId" gorm:"column:type_id;comment:分类id;size:20;NOT NULL DEFAULT '0'"`
	Pic       string     `json:"pic" form:"pic" gorm:"column:pic;comment:图片路径;size:500;"`
	PicKey    string     `json:"picKey" form:"picKey" gorm:"column:pic_key;comment:图片key;size:64;"`
}

// TableName LukInformation 表名
func (LukInformation) TableName() string {
	return "luk_information"
}
