package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type LukHeritageMetaverseOldSearch struct{
    luk.LukHeritageMetaverseOld
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
