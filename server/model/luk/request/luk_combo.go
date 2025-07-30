package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"time"
)

type LukComboSearch struct {
	luk.LukCombo
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Statu          *int       `json:"statu" form:"statu"`
	request.PageInfo
}

type LukComboSearchByOne struct {
	Address string `json:"address" form:"address"`
}
