package timingtask

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"sync"
)

type TimingTaskGroup struct {
	NftTask
	ComboTask
	LukTask
	UserTask
	BlockScanTask
	LpTask
	HeritageTask
}

var TimingTaskGroupApp = new(TimingTaskGroup)
var (
	lukServiceGroup    = service.ServiceGroupApp.LukServiceGroup
	TokenServiceGroup  = service.ServiceGroupApp.TokenServiceGroup
	lukUserService     = lukServiceGroup.LukUserService
	lukComboBuyService = lukServiceGroup.LukComboBuyService
	lukComboService    = lukServiceGroup.LukComboService
	BlockScanMt        = new(sync.RWMutex) //区块扫描锁
)
