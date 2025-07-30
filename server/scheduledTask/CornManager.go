package scheduledTask

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/robfig/cron"
)

// 定时任务管理模块
var cronStorage map[string]*cron.Cron
var stopCron bool // 判断是否执行定时任务

func Start() {
	// // 先初始化任务列表
	// initTaskMap()
	// 遍历taskMap，将不存在的InvokeTarget添加到数据库中
	for invokeTarget, _ := range taskMap {
		// 判断数据库是否存在
		total := 0
		global.GVA_DB.Model(luk.LukScheduledTask{}).Where("invoke_target = ?", invokeTarget).Select("count(*) as total").Pluck("total", &total)
		if total > 0 {
			continue
		}
		// 不存在该定时任务，则添加到数据库中
		task := luk.LukScheduledTask{}
		task.InvokeTarget = invokeTarget
		two := 2
		task.JobStatus = &two
		global.GVA_DB.Create(&task)
	}

	// 保存定时器
	cronStorage = map[string]*cron.Cron{}

	// 读取数据库的定时任务数据
	db := global.GVA_DB
	jobs := []luk.LukScheduledTask{}
	db.Find(&jobs)
	if len(jobs) > 0 {
		two := 2
		// 遍历数据执行定时任务
		for _, v := range jobs {
			if v.JobStatus == nil {
				v.JobStatus = &two
			}
			InitTask(v.InvokeTarget, v.CronExpression, *v.JobStatus)
		}
	}

}

// 初始化定时任务
func InitTask(taskName string, cronStr string, status int) bool {
	if !Add(taskName, cronStr) {
		// 创建任务失败，直接退出
		return false
	}
	if !ChangeStatus(taskName, status) {
		// 更改状态失败，直接退出
		return false
	}
	return true
}

// 新增定时任务
func Add(taskName string, cronStr string) bool {

	Del(taskName)
	// 创建新的定时任务
	task, ok := taskMap[taskName]
	if !ok {
		return false
	}
	// 0 0 5 31 2 ? 永远不会执行的表达式
	if stopCron {
		cronStr = "0 0 5 31 2 ?"
	}
	c := cron.New()
	c.AddFunc(cronStr, task)
	cronStorage[taskName] = c
	return true
}

// 删除定时任务
func Del(taskName string) {
	task, ok := cronStorage[taskName]
	if ok {
		task.Stop()
		// 如果任务已存在，删除任务
		delete(cronStorage, taskName)
	}
}

// 更改定时任务状态
func ChangeStatus(taskName string, status int) bool {
	task, ok := cronStorage[taskName]
	if !ok {
		return false
	}
	if status == 1 {
		task.Start()
	} else {
		task.Stop()
	}
	return true
}

// 立即执行一次
func ExecuteImmediately(taskName string) bool {
	task, ok := taskMap[taskName]
	if !ok {
		return false
	}
	go task()
	return true
}
