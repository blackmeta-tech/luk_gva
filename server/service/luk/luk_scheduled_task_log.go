package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
)

type LukScheduledTaskLogService struct {
}

// CreateLukamiScheduledTaskLog 创建LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) CreateLukScheduledTaskLog(lukScheduledTaskLog luk.LukScheduledTaskLog) (err error) {
	err = global.GVA_DB.Create(&lukScheduledTaskLog).Error
	return err
}

// DeleteLukScheduledTaskLog 删除LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) DeleteLukScheduledTaskLog(lukScheduledTaskLog luk.LukScheduledTaskLog) (err error) {
	err = global.GVA_DB.Delete(&lukScheduledTaskLog).Error
	return err
}

// DeleteLukScheduledTaskLogByIds 批量删除LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) DeleteLukScheduledTaskLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukScheduledTaskLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukScheduledTaskLog 更新LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) UpdateLukScheduledTaskLog(lukScheduledTaskLog luk.LukScheduledTaskLog) (err error) {
	err = global.GVA_DB.Save(&lukScheduledTaskLog).Error
	return err
}

// GetLukScheduledTaskLog 根据id获取LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) GetLukScheduledTaskLog(id uint) (lukScheduledTaskLog luk.LukScheduledTaskLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukScheduledTaskLog).Error
	return
}

// GetLukScheduledTaskLogInfoList 分页获取LukScheduledTaskLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukScheduledTaskLogService *LukScheduledTaskLogService) GetLukScheduledTaskLogInfoList(info lukReq.LukScheduledTaskLogSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukScheduledTaskLog{})
	var lukScheduledTaskLogs []luk.LukScheduledTaskLog
	if info.InvokeTarget != "" {
		db.Where("invoke_target like ?", "%"+info.InvokeTarget+"%")
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Order("id desc").Find(&lukScheduledTaskLogs).Error
	return lukScheduledTaskLogs, total, err
}
