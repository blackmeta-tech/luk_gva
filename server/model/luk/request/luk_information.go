package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type LukInformationSearch struct {
	luk.LukInformation
	request.PageInfo
}
