package timingtask

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"gorm.io/gorm"
	"strconv"
)

type BlockScanTask struct {
}

const OneScanNumber = 10000

func (b *BlockScanTask) BlockScan() {
	BlockScanMt.Lock()
	defer BlockScanMt.Unlock()
	lastBlockScan, err := lukServiceGroup.GetLukBlockHigthLast()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			firstBlock, err := TokenServiceGroup.BscTokenService.GetBscTokenInitBlockHigth()
			if err != nil {
				global.GVA_LOG.Error("获取初始高度失败")
			}
			fmt.Println("firstBlock", firstBlock)
			endBlock := firstBlock - 1
			lastBlockScan.EndBlock = &endBlock
		} else {
			global.GVA_LOG.Error(err.Error())
			return
		}
	}

	currentHeaderBlockNum, err := TokenServiceGroup.BscTokenService.GetBscTokenNewBlockHigth()
	fmt.Println("currentHeaderBlockNum:", currentHeaderBlockNum)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	// 新的区块高度 小于等于 上次扫描的高度 return
	if currentHeaderBlockNum <= *lastBlockScan.EndBlock {
		global.GVA_LOG.Sugar().Info(" currentHeaderBlockNum <=  lastBlockScan  ", currentHeaderBlockNum, "  ", *lastBlockScan.EndBlock)
		return
	}
	newStartBlock := *lastBlockScan.EndBlock + 1
	global.GVA_LOG.Sugar().Infof("开始执行扫描,扫描区块范围： %d ~ %d", newStartBlock, currentHeaderBlockNum)

	var allRecords []luk.LukBlockScanRecord
	var allBlockHigths []luk.LukTokenBlockHigth

	startBlock := newStartBlock
	blockHigh := luk.LukTokenBlockHigth{
		StartBlock: &startBlock,
		EndBlock:   &currentHeaderBlockNum,
		NowBlock:   &currentHeaderBlockNum,
	}
	records, _, err := TokenServiceGroup.BscTokenService.BlockScan(int64(newStartBlock))
	if err != nil {
		global.GVA_LOG.Sugar().Error("BlockScan  fail ", err.Error(), " StartBlock :", startBlock, " EndBlock : ", currentHeaderBlockNum)
		return
	}
	allRecords = append(allRecords, records...)
	allBlockHigths = append(allBlockHigths, blockHigh)
	b.updateBlockScanRecord(allRecords, allBlockHigths)
	global.GVA_LOG.Info("扫描完成")
}

func (b *BlockScanTask) updateBlockScanRecord(allRecords []luk.LukBlockScanRecord, allBlockHigths []luk.LukTokenBlockHigth) {
	db := global.GVA_DB
	tx := db.Begin()
	defer tx.Commit()
	err := tx.CreateInBatches(&allRecords, 100).Error
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		tx.Rollback()
		return
	}

	err = tx.CreateInBatches(&allBlockHigths, 100).Error
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		tx.Rollback()
		return
	}
}

//开定时任务24小时拉取补充数据
func (b *BlockScanTask) BlockScanSupplement() {
	BlockScanMt.Lock()
	defer BlockScanMt.Unlock()
	//获取三天前的区块高度
	lastBlockScan, err := lukServiceGroup.GetLukBlockHigthBefore()
	if err != nil {
		firstBlock, _ := TokenServiceGroup.BscTokenService.GetBscTokenInitBlockHigth()
		fmt.Println("====", firstBlock)
		endBlock := firstBlock - 1
		lastBlockScan.EndBlock = &endBlock
	}
	newStartBlock := *lastBlockScan.EndBlock + 1
	i := 0
	// 每次扫描10000个 不需要更新数据库区块信息数据
	var allRecords []luk.LukBlockScanRecord
	for i <= 0 {
		fmt.Println("======", newStartBlock)
		records, endNumber, err := TokenServiceGroup.BscTokenService.BlockScan(int64(newStartBlock))
		if err != nil {
			global.GVA_LOG.Sugar().Error("BlockScan  fail ", err.Error(), " StartBlock :", newStartBlock)
			return
		}
		fmt.Println("=====", newStartBlock, endNumber)
		allRecords = append(allRecords, records...)
		if endNumber == newStartBlock {
			i = 1
			break
		}
		newStartBlock = endNumber
	}

	err = global.GVA_DB.CreateInBatches(allRecords, 100).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	global.GVA_LOG.Info("扫描完成,已补充数据" + strconv.Itoa(len(allRecords)) + "条")
}
