package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/shopspring/decimal"
)

type LukReflowService struct {
}

func (d *LukReflowService) GetLukReflowInfoList(info lukReq.LukReflowSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&luk.LukReflow{})
	var lukReflows []luk.LukReflow
	if info.Type > 0 {
		db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&lukReflows).Error
	return lukReflows, total, err
}

//查询某一天的回流的量
func (d *LukReflowService) GetLukReflowBydate(date string, _type emun.ReflowType) (amount decimal.Decimal, err error) {
	reflow := luk.LukReflow{}
	err = global.GVA_DB.Model(&luk.LukReflow{}).Select("sum(amount) amount").
		Where("created_at >= ? and created_at <= ? and type = ?", date+" 00:00:00", date+" 23:59:59", _type).First(&reflow).Error
	if err == nil {
		amount = reflow.Amount
	}
	return
}
