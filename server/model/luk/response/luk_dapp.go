package response

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/shopspring/decimal"
	"time"
)

type PoolBasicInfo struct {
	Luk  decimal.Decimal `json:"luk"`  //dba数量
	Usdt decimal.Decimal `json:"usdt"` //usdt数量
}

type TransactionRecordList struct {
	CreatedAt time.Time `json:"createdAt"`
	Address   string    `json:"address"`
	Content   string    `json:"content"`
}

type PartnerList struct {
	Level   emun.PartnerType `json:"level" form:"level"`
	AreaMax decimal.Decimal  `json:"areaMax" form:"areaMax"`
	AreaMin decimal.Decimal  `json:"areaMin" form:"areaMin"`
}

type BlockScanRecord struct {
	LukAmount  decimal.Decimal `json:"lukAmount"`
	UsdtAmount decimal.Decimal `json:"usdtAmount"`
	LpAmount   decimal.Decimal `json:"lpAmount"`
	Time       time.Time       `json:"time"`
	From       string          `json:"from"`
}

type HeritageInfo struct {
	State                  bool            `json:"state"`                  //移产状态
	MetaverseBalance       decimal.Decimal `json:"metaverseBalance"`       //新元宇宙余额
	MetaverseFreed         decimal.Decimal `json:"metaverseFreed"`         //新元宇宙释放余额
	MetaversePercentage    decimal.Decimal `json:"metaversePercentage"`    //新元宇宙分配比例
	MetaverseOldBalance    decimal.Decimal `json:"metaverseOldBalance"`    //老元宇宙余额
	MetaverseOldFreed      decimal.Decimal `json:"metaverseOldFreed"`      //老元宇宙释放余额
	MetaverseOldPercentage decimal.Decimal `json:"metaverseOldPercentage"` //老元宇宙分配比例
	LpBalance              decimal.Decimal `json:"lpBalance"`              //lp余额
	LpFreed                decimal.Decimal `json:"lpFreed"`                //lp释放余额
	LpPercentage           decimal.Decimal `json:"lpPercentage"`           //lp分配比例
}
