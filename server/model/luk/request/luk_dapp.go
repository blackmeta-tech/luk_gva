package request

import (
	emun "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BlockScanRecordSearch struct {
	Type emun.MethodType `json:"type" form:"type"`
	From string          `json:"from"  form:"from"`
	request.PageInfo
}
