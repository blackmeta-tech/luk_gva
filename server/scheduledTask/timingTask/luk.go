package timingtask

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"sync"
)

type LukTask struct {
}

//拉取链上数据更新到数据库
func (LukTask) SyncChainData() {
	var wg sync.WaitGroup
	wg.Add(2)
	data := luk.LukPoolBasic{}
	go func() {
		defer wg.Done()
		data.DestructionLuk = decimal.Zero
	}()
	go func() {
		defer wg.Done()
		totalLuk, totalUsdt, err := service.ServiceGroupApp.TokenServiceGroup.PairTokenService.GetReserves()
		fmt.Println("=====luk以及usdt总量", totalLuk, totalUsdt, err)
		if err == nil {
			data.Luk = totalLuk
			data.Usdt = totalUsdt
		}
	}()
	wg.Wait()
	db := global.GVA_DB
	tx := db.Begin()
	defer tx.Commit()

	err := lukServiceGroup.LukPoolBasicService.Update(tx, data)
	if err != nil {
		global.GVA_LOG.Error("更新链上失败!", zap.Error(err))
		return
	}
}
