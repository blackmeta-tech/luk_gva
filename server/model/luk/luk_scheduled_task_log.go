// 自动生成模板WhoScheduledTaskLog
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WhoScheduledTaskLog 结构体
// 如果含有time.Time 请自行import time包
type LukScheduledTaskLog struct {
	global.GVA_MODEL
	CreatedBy    *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;size:19;NOT NULL DEFAULT '0'"`
	DeletedBy    *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:;size:19;NOT NULL DEFAULT '0'"`
	InvokeTarget string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:定时任务执行字符串;size:32;"`
	Remark       string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:4294967295;"`
	Status       *int   `json:"status" form:"status" gorm:"column:status;comment:执行状态(1成功，2失败);size:10;NOT NULL DEFAULT '0'"`
	UpdatedBy    *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:;size:19;NOT NULL DEFAULT '0'"`
}

// TableName WhoScheduledTaskLog 表名
func (LukScheduledTaskLog) TableName() string {
	return "luk_scheduled_task_log"
}
