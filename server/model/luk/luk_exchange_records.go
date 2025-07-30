package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukExchangeRecords 兑换记录
type LukExchangeRecords struct {
	global.GVA_MODEL
	Address   string          `json:"address" form:"address" gorm:"column:address;comment:用户地址;size:200;"`
	TxHash    string          `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:tx 哈希;size:191;index:idx_name_txhash"`
	Type      emun.MethodType `json:"type" form:"type" gorm:"column:type;type:tinyint(0) NOT NULL DEFAULT '0';comment:兑换类型 默认0luk兑换usdt 1usdt兑换luk;"`
	Luk       decimal.Decimal `json:"luk" form:"luk" gorm:"column:luk;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:luk个数;"`
	Usdt      decimal.Decimal `json:"usdt" form:"usdt" gorm:"column:usdt;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:usdt个数;"`
	ChargeLuk decimal.Decimal `json:"chargeLuk" form:"chargeLuk" gorm:"column:charge_luk;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:手续费，换算成luk;"`
	Status    int             `json:"status" form:"status" gorm:"type:tinyint(2) NOT NULL DEFAULT '0';column:status;comment:状态 1成功 0等待 -1失败;"`
	Errmsg    string          `json:"errmsg" form:"errmsg" gorm:"type:varchar(255) NOT NULL DEFAULT '';column:errmsg;comment:报错提示;"`
}

// TableName LukBlockScanRecord 表名
func (LukExchangeRecords) TableName() string {
	return "luk_exchange_records"
}

func (LukExchangeRecords) GetByTxHash(txHash string) (info LukExchangeRecords, err error) {
	err = global.GVA_DB.Where("tx_hash = ?", txHash).First(&info).Error
	return
}
