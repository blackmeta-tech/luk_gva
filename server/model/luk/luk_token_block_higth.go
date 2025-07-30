// 自动生成模板LukTokenBlockHigth
package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LukTokenBlockHigth 结构体
type LukTokenBlockHigth struct {
	global.GVA_MODEL
	StartBlock *uint64 `json:"startBlock" form:"startBlock" gorm:"column:startBlock;comment:开始区块高度;"`
	EndBlock   *uint64 `json:"endBlock" form:"endBlock" gorm:"column:endBlock;comment:结束区块高度;"`
	NowBlock   *uint64 `json:"nowBlock" form:"nowBlock" gorm:"column:nowBlock;comment:当前区块高度;"`
}

// TableName LukTokenBlockHigth 表名
func (LukTokenBlockHigth) TableName() string {
	return "luk_token_block_higth"
}
