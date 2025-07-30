package request

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"time"
)

type LukUserSearch struct {
	luk.LukUser
	StartCreatedAt *time.Time        `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time        `json:"endCreatedAt" form:"endCreatedAt"`
	Communitylevel *emun.PartnerType `json:"communitylevel" form:"communitylevel"`
	Linkagelevel   *emun.PartnerType `json:"linkagelevel" form:"linkagelevel"`
	Islinkage      *bool             `json:"islinkage" form:"islinkage"`
	request.PageInfo
}

type QuerryUser struct {
	Address string `json:"address" form:"address"`
}
