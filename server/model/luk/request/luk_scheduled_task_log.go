package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type LukScheduledTaskLogSearch struct {
	luk.LukScheduledTaskLog
	request.PageInfo
}
