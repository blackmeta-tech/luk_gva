package global

import (
	"github.com/shopspring/decimal"
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id" form:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt" form:"createdAt"`     // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                 // 删除时间
}

type LUK_CONFIG struct {
	NftIssued         int64           `json:"nftIssued" form:"nftIssued"`               //Nft发行量
	WithdrawalQuota   decimal.Decimal `json:"withdrawalQuota" form:"withdrawalQuota"`   //luk最低提现额度
	WithdrawalLuk     int64           `json:"withdrawalLuk" form:"withdrawalLuk"`       //luk当日提现限额
	WithdrawalUsdt    int64           `json:"withdrawalUsdt" form:"withdrawalUsdt"`     //usdt当日提现限额`
	WithdrawalCharge  int64           `json:"withdrawalCharge" form:"withdrawalCharge"` //提现手续费
	Online            bool            `json:"online" form:"online"`                     //上线状态
	OnlineText        string          `json:"onlineText"`                               //上线显示文本
	RatioMetaverse1   decimal.Decimal `json:"ratioMetaverse1"`                          //元宇宙释放百分比 模式1
	RatioMetaverse2   decimal.Decimal `json:"ratioMetaverse2"`                          //元宇宙释放百分比 模式2
	RatioMetaverse3   decimal.Decimal `json:"ratioMetaverse3"`                          //元宇宙释放百分比 模式3
	RatioMetaverseOld decimal.Decimal `json:"ratioMetaverseOld"`                        //老元宇宙释放百分比
	RatioLp           decimal.Decimal `json:"ratioLp"`                                  //Lp释放百分比
	ComboRecommend1   decimal.Decimal `json:"comboRecommend1"`                          //套餐推荐1代分红百分比
	ComboRecommend2   decimal.Decimal `json:"comboRecommend2"`                          //套餐推荐2-5代分红百分比
	ComboRecommend6   decimal.Decimal `json:"comboRecommend6"`                          //套餐推荐6-10代分红百分比
	LpRebate          decimal.Decimal `json:"lpRebate"`                                 //LP返佣百分比
	LpRebateHighest   decimal.Decimal `json:"lpRebateHighest"`                          //Lp最高返佣额度
	RebatePerformance decimal.Decimal `json:"rebatePerformance"`                        //返佣业绩分红百分比
}
