package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
)

type LukPartnerService struct {
}

// UpdateLukPartner 更新LukPartner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukPartnerService *LukPartnerService) UpdateLukPartner(lukPartner luk.LukPartner) (err error) {
	err = global.GVA_DB.Save(&lukPartner).Error
	return err
}

// GetLukPartner 根据id获取LukPartner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukPartnerService *LukPartnerService) GetLukPartner(id uint) (lukPartner luk.LukPartner, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukPartner).Error
	return
}

// GetLukPartnerInfoList 分页获取LukPartner记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukPartnerService *LukPartnerService) GetLukPartnerInfoList(info lukReq.LukPartnerSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukPartner{})
	var lukPartners []luk.LukPartner
	if info.Type > 0 {
		db = db.Where("type = ? ", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukPartners).Error
	return lukPartners, total, err
}

func (lukPartnerService *LukPartnerService) QueryList(_type emun.UserType, _order string) (list []response.PartnerList, err error) {
	err = global.GVA_DB.Model(&luk.LukPartner{}).Where("type = ? ", _type).Order("id " + _order).Find(&list).Error
	return
}
