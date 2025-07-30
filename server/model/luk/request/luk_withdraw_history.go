package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type LukWithdrawHistorySearch struct {
	luk.LukWithdrawHistory
	request.PageInfo
	OrderKey  string `json:"orderKey" form:"orderKey"` // 排序
	Desc      bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
}
