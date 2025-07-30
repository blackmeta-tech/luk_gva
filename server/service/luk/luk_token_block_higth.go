package luk

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"time"
)

type LukTokenBlockHigthService struct {
}

// CreateLukTokenBlockHigth 创建LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) CreateLukTokenBlockHigth(lukTokenBlockHigth luk.LukTokenBlockHigth) (err error) {
	err = global.GVA_DB.Create(&lukTokenBlockHigth).Error
	return err
}

// DeleteLukTokenBlockHigth 删除LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) DeleteLukTokenBlockHigth(lukTokenBlockHigth luk.LukTokenBlockHigth) (err error) {
	err = global.GVA_DB.Delete(&lukTokenBlockHigth).Error
	return err
}

// DeleteLukTokenBlockHigthByIds 批量删除LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) DeleteLukTokenBlockHigthByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]luk.LukTokenBlockHigth{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLukTokenBlockHigth 更新LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) UpdateLukTokenBlockHigth(lukTokenBlockHigth luk.LukTokenBlockHigth) (err error) {
	err = global.GVA_DB.Save(&lukTokenBlockHigth).Error
	return err
}

// GetLukTokenBlockHigth 根据id获取LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) GetLukTokenBlockHigth(id uint) (lukTokenBlockHigth luk.LukTokenBlockHigth, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukTokenBlockHigth).Error
	return
}

func (lukTokenBlockHigthService *LukTokenBlockHigthService) GetLukBlockHigthLast() (lukTokenBlockHigth luk.LukTokenBlockHigth, err error) {
	err = global.GVA_DB.Last(&lukTokenBlockHigth).Error
	return
}

//获取24小时前的区块高度
func (lukTokenBlockHigthService *LukTokenBlockHigthService) GetLukBlockHigthBefore() (lukTokenBlockHigth luk.LukTokenBlockHigth, err error) {
	t := time.Now().AddDate(0, 0, -1)
	fmt.Println("====", t)
	err = global.GVA_DB.Where("created_at <= ?", t).Order("id desc").First(&lukTokenBlockHigth).Error
	return
}

// GetLukTokenBlockHigthInfoList 分页获取LukTokenBlockHigth记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukTokenBlockHigthService *LukTokenBlockHigthService) GetLukTokenBlockHigthInfoList(info lukReq.LukTokenBlockHigthSearch) (list []luk.LukTokenBlockHigth, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukTokenBlockHigth{})
	var lukTokenBlockHigths []luk.LukTokenBlockHigth
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&lukTokenBlockHigths).Error
	return lukTokenBlockHigths, total, err
}
