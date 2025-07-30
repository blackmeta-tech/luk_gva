package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type LukUserAddressSearch struct {
	Address string `form:"address"`
	request.PageInfo
}

type LukUserAddressUpdate struct {
	Address string `json:"address"`
	Remark  string `json:"remark"`
	Type    int    `json:"type"`
}
