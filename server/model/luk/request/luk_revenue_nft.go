package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type LukRevenueNftSearch struct {
	luk.LukRevenueNft
	request.PageInfo
}

type LukRevenueHeritageSearch struct {
	luk.LukRevenueHeritage
	request.PageInfo
}
