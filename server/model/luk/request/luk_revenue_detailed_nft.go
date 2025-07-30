package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type LukRevenueDetailedNftSearch struct {
	luk.LukRevenueDetailedNft
	request.PageInfo
}

type LukRevenueDetailedHeritageSearch struct {
	luk.LukRevenueDetailedHeritage
	request.PageInfo
}
