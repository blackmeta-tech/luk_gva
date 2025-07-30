package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"time"
)

type LukComboBuySearch struct {
	luk.LukComboBuy
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Mode           int        `json:"mode" form:"mode"`
	Subclass       int        `json:"subclass" form:"subclass"`
	Generation     []string   `json:"generation" form:"generation"`
	request.PageInfo
}
