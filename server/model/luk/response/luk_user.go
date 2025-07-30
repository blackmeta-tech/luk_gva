package response

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/shopspring/decimal"
	"time"
)

type SubclassList struct {
	Address         string           `json:"address" form:"address"`
	IsLinkage       bool             `json:"isLinkage" form:"isLinkage"`
	LinkageLevel    emun.PartnerType `json:"linkageLevel" form:"linkageLevel"`
	PerformanceTime time.Time        `json:"performanceTime" form:"performanceTime"`
	Performance     decimal.Decimal  `json:"performance" form:"performance"`
}

type LinkageList struct {
	LinkageLevel emun.PartnerType `json:"linkageLevel"` //等级
	Proportion   int              `json:"proportion"`   //占比
	Count        int64            `json:"count"`        //人数
}
