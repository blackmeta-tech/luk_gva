package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type LukDividendsNftLog struct {
	global.GVA_MODEL
	NftId uint   `json:"nftId" form:"nftId" gorm:"column:nft_id;comment:NFTID;size:10;"`
	Date  string `json:"date" orm:"date" gorm:"column:date;comment:日期;type:date"`
}

// TableName DxbaKline 表名
func (LukDividendsNftLog) TableName() string {
	return "luk_dividends_nft_log"
}

func (LukDividendsNftLog) QueryByNftId(nftId uint) (count int64) {
	global.GVA_DB.Debug().Where("nft_id = ? ", nftId).Count(&count)
	return
}
