package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/scheduledTask"
	timingtask "github.com/flipped-aurora/gin-vue-admin/server/scheduledTask/timingTask"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	decimal.DivisionPrecision = 8 //设置8位精度 与数据库同步
	global.GVA_VP = core.Viper()  // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
		startTimer()
	}
	core.RunWindowsServer()
}

func startTimer() {
	scheduledTask.InitTaskMap() // 初始化任务列表
	/*scheduledTask.AddTask("ScheduledTaskTest", scheduledTask.Test)*/
	scheduledTask.AddTask("SyncNftAddress", timingtask.TimingTaskGroupApp.SyncNftAddress)           //同步NFT地址
	scheduledTask.AddTask("SyncChainData", timingtask.TimingTaskGroupApp.SyncChainData)             //同步链上数据
	scheduledTask.AddTask("AllBonus", timingtask.TimingTaskGroupApp.AllBonus)                       //分红
	scheduledTask.AddTask("SyncComboNum", timingtask.TimingTaskGroupApp.SyncComboNum)               //每天分红
	scheduledTask.AddTask("BlockScan", timingtask.TimingTaskGroupApp.BlockScan)                     // 区块扫描
	scheduledTask.AddTask("BlockScanSupplement", timingtask.TimingTaskGroupApp.BlockScanSupplement) // 24小时区块扫描
	scheduledTask.AddTask("Heritage", timingtask.TimingTaskGroupApp.Heritage)                       // 移产分红
	scheduledTask.AddTask("ExtendMaturityAt", timingtask.TimingTaskGroupApp.ExtendMaturityAt)       // 延长分红时间
	scheduledTask.Start()                                                                           // 初始化定时任务
}
