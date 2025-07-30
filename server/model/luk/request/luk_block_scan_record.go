package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type LukBlockScanRecordSearch struct {
	Address string `json:"address" form:"address"`
	request.PageInfo
}
