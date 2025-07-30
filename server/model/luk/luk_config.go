package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

type LukConfig struct {
	global.GVA_MODEL
	NftIssued         int64           `json:"nftIssued" form:"nftIssued" gorm:"column:nft_issued;comment:Nft发行量;size:11;"`                     //Nft发行量
	WithdrawalQuota   decimal.Decimal `json:"withdrawalQuota" form:"withdrawalQuota" gorm:"column:withdrawal_quota;comment:提现限制;size:11;"`     //luk最低提现额度
	WithdrawalLuk     int64           `json:"withdrawalLuk" form:"withdrawalLuk" gorm:"column:withdrawal_luk;comment:luk当日提现限额;size:11;"`      //luk当日提现限额
	WithdrawalUsdt    int64           `json:"withdrawalUsdt" form:"withdrawalUsdt" gorm:"column:withdrawal_usdt;comment:usdt当日提现限额;size:11;"`  //usdt当日提现限额
	WithdrawalCharge  int64           `json:"withdrawalCharge" form:"withdrawalCharge" gorm:"column:withdrawal_charge;comment:提现手续费;size:11;"` //提现手续费
	Online            bool            `json:"online" form:"online" gorm:"column:online;comment:上线 0否 1是;type:tinyint(2) NOT NULL DEFAULT '0'"` //上线状态
	OnlineText        string          `json:"onlineText" form:"onlinTexte" gorm:"column:online_text;comment:上线文本;size:255;"`                   //上线文本
	RatioMetaverse1   decimal.Decimal `json:"ratioMetaverse1" form:"ratioMetaverse1" gorm:"column:ratio_metaverse1;comment:元宇宙释放百分比 模式1;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	RatioMetaverse2   decimal.Decimal `json:"ratioMetaverse2" form:"ratioMetaverse2" gorm:"column:ratio_metaverse2;comment:元宇宙释放百分比 模式2;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	RatioMetaverse3   decimal.Decimal `json:"ratioMetaverse3" form:"ratioMetaverse3" gorm:"column:ratio_metaverse3;comment:元宇宙释放百分比 模式3;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	RatioMetaverseOld decimal.Decimal `json:"ratioMetaverseOld" form:"ratioMetaverseOld" gorm:"column:ratio_metaverse_old;comment:老元宇宙释放百分比type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	RatioLp           decimal.Decimal `json:"ratioLp" form:"ratioLp" gorm:"column:ratio_lp;comment:Lp释放百分比;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	ComboRecommend1   decimal.Decimal `json:"comboRecommend1" form:"comboRecommend1" gorm:"column:combo_recommend1;comment:套餐推荐1代;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	ComboRecommend2   decimal.Decimal `json:"comboRecommend2" form:"comboRecommend2" gorm:"column:combo_recommend2;comment:套餐推荐2代;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	ComboRecommend6   decimal.Decimal `json:"comboRecommend6" form:"comboRecommend6" gorm:"column:combo_recommend6;comment:套餐推荐6代;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	LpRebate          decimal.Decimal `json:"lpRebate" form:"lpRebate" gorm:"column:lp_rebate;comment:Lp返佣百分比;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	LpRebateHighest   decimal.Decimal `json:"lpRebateHighest" form:"lpRebateHighest" gorm:"column:lp_rebate_highest;comment:Lp最高返佣额度;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
	RebatePerformance decimal.Decimal `json:"rebatePerformance" form:"rebatePerformance" gorm:"column:rebate_performance;comment:返佣业绩分红百分比;type:decimal(32,2) NOT NULL DEFAULT '0.00';"`
}

// TableName LukConfig 表名
func (LukConfig) TableName() string {
	return "luk_config"
}

func (LukConfig) GetData() (data LukConfig, err error) {
	err = global.GVA_DB.Where("id = 1").First(&data).Error
	return
}

// 更新or创建数据
func (LukConfig) Update(data LukConfig) (err error) {
	data.UpdatedAt = time.Now()
	err = global.GVA_DB.Debug().Save(&data).Error
	return
}
