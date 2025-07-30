package luk

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/scheduledTask"
)

type LukScheduledTaskService struct {
}

// CreateLukScheduledTask 创建LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) CreateLukScheduledTask(lukScheduledTask luk.LukScheduledTask) (err error) {
	// 先判断任务目标字符串是否已存在
	if lukScheduledTask.InvokeTarget == "" {
		err = errors.New("任务目标字符串不能为空")
		return
	}
	two := 2
	if lukScheduledTask.JobStatus == nil {
		lukScheduledTask.JobStatus = &two
	}
	// 判断数据库中是否已有重名的目标字符串
	total := 0
	global.GVA_DB.Model(&luk.LukScheduledTask{}).Where("invoke_target = ?", lukScheduledTask.InvokeTarget).Select("count(*) as total").Pluck("total", &total)
	if total > 0 {
		err = errors.New("创建失败，任务目标字符串已存在")
		return
	}
	// 将任务添加到定时器中
	if !scheduledTask.InitTask(lukScheduledTask.InvokeTarget, lukScheduledTask.CronExpression, *lukScheduledTask.JobStatus) {
		lukScheduledTask.JobStatus = &two
	}
	err = global.GVA_DB.Create(&lukScheduledTask).Error
	return
}

// DeleteLukScheduledTask 删除LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) DeleteLukScheduledTask(lukScheduledTask luk.LukScheduledTask) (err error) {
	// 先根据id查询出目标调用字符串
	task := luk.LukScheduledTask{}
	global.GVA_DB.Where("id = ?", lukScheduledTask.ID).Take(&task)
	if task.ID != 0 {
		// 删除对应的定时任务
		scheduledTask.Del(task.InvokeTarget)
	}
	err = global.GVA_DB.Delete(&lukScheduledTask).Error
	return err
}

// DeleteLukScheduledTaskByIds 批量删除LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) DeleteLukScheduledTaskByIds(ids request.IdsReq) (err error) {
	// 先查出所有要删除的定时任务的invokeTarget，停止定时任务
	invokeTargetList := []string{}
	global.GVA_DB.Model(&luk.LukScheduledTask{}).Where("id in (?)", ids.Ids).Pluck("invoke_target", &invokeTargetList)
	if len(invokeTargetList) > 0 {
		for _, v := range invokeTargetList {
			scheduledTask.Del(v)
		}
	}
	err = global.GVA_DB.Delete(&[]luk.LukScheduledTask{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukScheduledTask 更新LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) UpdateLukScheduledTask(lukScheduledTask luk.LukScheduledTask) (err error) {
	// 判断id是否存在
	if lukScheduledTask.ID == 0 {
		err = errors.New("传入id不能为空")
		return
	}
	// 判断任务目标字符串是否已存在
	if lukScheduledTask.InvokeTarget == "" {
		err = errors.New("任务目标字符串不能为空")
		return
	}
	// 判断是否有重复的invokeTarget
	total := 0
	global.GVA_DB.Where("id != ? and invoke_target = ?", lukScheduledTask.ID, lukScheduledTask.InvokeTarget).Select("count(*) as total").Pluck("total", &total)
	if total > 0 {
		err = errors.New("修改失败，任务目标字符串已存在")
		return
	}
	// 根据id查询
	target := luk.LukScheduledTask{}
	err = global.GVA_DB.Where("id = ?", lukScheduledTask.ID).Take(&target).Error
	if err != nil {
		return
	}

	two := 2
	// 停掉原来的invokeTarget的定时任务并删除
	scheduledTask.Del(lukScheduledTask.InvokeTarget)
	if lukScheduledTask.JobStatus == nil {
		lukScheduledTask.JobStatus = &two
	}
	// 创建新的定时任务
	if scheduledTask.InitTask(lukScheduledTask.InvokeTarget, lukScheduledTask.CronExpression, *lukScheduledTask.JobStatus) {
		target.JobStatus = lukScheduledTask.JobStatus
	} else {
		// 如果创建定时任务状态失败，就是停用状态
		target.JobStatus = &two
	}
	// 修改对应字段
	target.JobName = lukScheduledTask.JobName
	target.InvokeTarget = lukScheduledTask.InvokeTarget
	target.CronExpression = lukScheduledTask.CronExpression
	target.Remark = lukScheduledTask.Remark
	err = global.GVA_DB.Save(&target).Error
	return err
}

// GetLukScheduledTask 根据id获取LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) GetLukScheduledTask(id uint) (lukScheduledTask luk.LukScheduledTask, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukScheduledTask).Error
	return
}

// GetLukScheduledTaskInfoList 分页获取LukScheduledTask记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskService *LukScheduledTaskService) GetLukScheduledTaskInfoList(info lukReq.LukScheduledTaskSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukScheduledTask{})
	var lukScheduledTasks []luk.LukScheduledTask
	if info.JobName != "" {
		db.Where("job_name like ?", "%"+info.JobName+"%")
	}
	if info.InvokeTarget != "" {
		db.Where("invoke_target like ?", "%"+info.InvokeTarget+"%")
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukScheduledTasks).Error
	return lukScheduledTasks, total, err
}

func (lukScheduledTaskService *LukScheduledTaskService) ExecuteImmediately(info luk.LukScheduledTask) (err error) {
	if info.ID == 0 {
		err = errors.New("传入id不存在")
		return
	}
	// 根据id查询出对应的执行字符串
	db := global.GVA_DB.Model(&luk.LukScheduledTask{})
	target := luk.LukScheduledTask{}
	db.Where("id = ?", info.ID).Take(&target)
	if target.ID == 0 {
		err = errors.New("查无该id任务")
		return
	}
	if !scheduledTask.ExecuteImmediately(target.InvokeTarget) {
		err = errors.New("任务不存在，开始任务失败")
		return
	}
	return
}
