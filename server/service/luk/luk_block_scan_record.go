package luk

import (
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"strings"
)

type LukBlockScanRecordService struct {
}

// CreateLukBlockScanRecord 创建LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) CreateLukBlockScanRecord(lukBlockScanRecord luk.LukBlockScanRecord) (err error) {
	err = global.GVA_DB.Create(&lukBlockScanRecord).Error
	return err
}

// DeleteLukBlockScanRecord 删除LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) DeleteLukBlockScanRecord(lukBlockScanRecord luk.LukBlockScanRecord) (err error) {
	err = global.GVA_DB.Delete(&lukBlockScanRecord).Error
	return err
}

// DeleteLukBlockScanRecordByIds 批量删除LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) DeleteLukBlockScanRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukBlockScanRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukBlockScanRecord 更新LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) UpdateLukBlockScanRecord(lukBlockScanRecord luk.LukBlockScanRecord) (err error) {
	err = global.GVA_DB.Save(&lukBlockScanRecord).Error
	return err
}

// GetLukBlockScanRecord 根据id获取LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) GetLukBlockScanRecord(id uint) (lukBlockScanRecord luk.LukBlockScanRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukBlockScanRecord).Error
	return
}

// GetLukBlockScanRecordInfoList 分页获取LukBlockScanRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukBlockScanRecordService *LukBlockScanRecordService) GetLukBlockScanRecordInfoList(info lukReq.LukBlockScanRecordSearch) (list []luk.LukBlockScanRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukBlockScanRecord{})
	var lukBlockScanRecords []luk.LukBlockScanRecord
	db.Where("methodName = ? or methodName = ?", enum.TokenMethodAddLiquidity, enum.TokenMethodRemoveLiquidity)
	if info.Address != "" {
		db.Where("`from` = ? or `to` = ?", strings.ToLower(info.Address), strings.ToLower(info.Address))
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("time desc").Find(&lukBlockScanRecords).Error
	return lukBlockScanRecords, total, err
}

func (lukBlockScanRecordService *LukBlockScanRecordService) GetRemoveLiquidity() (lukBlockScanRecord []luk.LukBlockScanRecord) {
	global.GVA_DB.Debug().Model(&luk.LukBlockScanRecord{}).Select("sum(lp_amount) lp_amount, sum(luk_amount) as luk_amount, sum(usdt_amount) as usdt_amount, sum(luk_service_fee) as luk_service_fee, type").Group("type").Find(&lukBlockScanRecord)
	return
}
