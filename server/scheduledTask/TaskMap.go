package scheduledTask

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// key 任务字符串 value 具体执行的方法
var taskMap map[string]TaskFunction

// 确保同一种定时任务不是同时执行的
var taskLockMap map[string]*sync.Mutex

func addTask(taskName string, task TaskFunction) {
	taskLockMap[taskName] = &sync.Mutex{}
	// 生成带锁和日志的任务
	taskMap[taskName] = func() {
		defer taskLockMap[taskName].Unlock() // 解锁
		defer saveLog(taskName)              // 保存日志
		taskLockMap[taskName].Lock()
		task()
	}

}

// 生成日志记录
func saveLog(taskName string) {
	status := 1
	remark := ""
	if err := recover(); err != nil {
		remark = fmt.Sprintf("%+v", err)
		status = 2
	}
	db := global.GVA_DB
	var taskLog luk.LukScheduledTaskLog
	taskLog.InvokeTarget = taskName
	taskLog.Status = &status
	taskLog.Remark = remark
	db.Create(&taskLog)
}
