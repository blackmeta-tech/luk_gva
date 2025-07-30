// 自动生成模板LukBlockScanRecord
package luk

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

// LukBlockScanRecord 结构体
type LukBlockScanRecord struct {
	global.GVA_MODEL
	Block          *uint64              `json:"block" form:"block" gorm:"column:block;comment:区块编号;"`
	TxHash         string               `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:tx 哈希;size:191;index:idx_name_txhash"`
	MethodName     emun.TokenMethodName `json:"methodName" form:"methodName" gorm:"column:method_name;comment:调用ABI方法名;size:255;"`
	From           string               `json:"from" form:"from" gorm:"column:from;comment:From 地址;size:191;"`
	To             string               `json:"to" form:"to" gorm:"column:to;comment:to地址;size:191;"`
	LukAmount      decimal.Decimal      `json:"lukAmount" form:"lukAmount" gorm:"column:luk_amount;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:lukAmount 金额;"`
	UsdtAmount     decimal.Decimal      `json:"usdtAmount" form:"usdtAmount" gorm:"column:usdt_amount;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:usdtAmount 金额;"`
	LpAmount       decimal.Decimal      `json:"lpAmount" form:"lpAmount" gorm:"column:lp_amount;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:lpAmount 金额;"`
	LukServiceFee  decimal.Decimal      `json:"lukServiceFee" form:"lukServiceFee" gorm:"column:luk_service_fee;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:luk手续费;"`
	UsdtServiceFee decimal.Decimal      `json:"usdtServiceFee" form:"usdtServiceFee" gorm:"column:usdt_service_fee;type:decimal(32,8) NOT NULL DEFAULT '0.000000000000000000';comment:usdt手续费;"`
	Status         *int                 `json:"status" form:"status" gorm:"type:tinyint(2) NOT NULL DEFAULT '0';column:status;comment:状态 1成功 0失败;"`
	TimeStamp      *int64               `json:"timeStamp" form:"timeStamp" gorm:"column:time_stamp;comment:时间戳;"`
	Time           time.Time            `json:"time" form:"time" gorm:"comment:时间"`
	Type           emun.MethodType      `json:"type" form:"type" gorm:"type:tinyint(2) NOT NULL DEFAULT '0';column:type;comment:类型;"`
}

// TableName LukBlockScanRecord 表名
func (LukBlockScanRecord) TableName() string {
	return "luk_block_scan_record"
}
