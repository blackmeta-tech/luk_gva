// 自动生成模板WhoScheduledTask
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WhoScheduledTask 结构体
// 如果含有time.Time 请自行import time包
type LukScheduledTask struct {
	global.GVA_MODEL
	CreatedBy      *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;size:19;NOT NULL DEFAULT '0'"`
	CronExpression string `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:cron执行表达式;size:32;"`
	DeletedBy      *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:;size:19;NOT NULL DEFAULT '0'"`
	InvokeTarget   string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:调用目标字符串;size:32;"`
	JobName        string `json:"jobName" form:"jobName" gorm:"column:job_name;comment:任务名称;size:32;"`
	JobStatus      *int   `json:"jobStatus" form:"jobStatus" gorm:"column:job_status;comment:状态（1正常 2暂停）;size:10;NOT NULL DEFAULT '0'"`
	Remark         string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`
	UpdatedBy      *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:;size:19;NOT NULL DEFAULT '0'"`
}

// TableName WhoScheduledTask 表名
func (LukScheduledTask) TableName() string {
	return "luk_scheduled_task"
}
