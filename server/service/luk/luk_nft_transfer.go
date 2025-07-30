package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukReq "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"strings"
)

type LukNftTransferService struct {
}

// CreateLukNftTransfer 创建LukNftTransfer记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftTransferService *LukNftTransferService) CreateLukNftTransfer(transfer luk.LukNftTransfer) (err error) {
	if transfer.From == "" || transfer.TxHash == "" || transfer.To == "" || transfer.NftId == 0 {
		err = errors.New("参数不足")
		return
	}
	transfer.To = strings.ToLower(transfer.To)
	transfer.From = strings.ToLower(transfer.From)
	//查下luk是否存在
	nft, _ := (&LukNftService{}).GetLukNft(transfer.NftId)
	if nft.ID == 0 {
		err = errors.New("找不到NFT")
		return
	}
	if nft.Address != transfer.From {
		err = errors.New("NFT不属于你")
		return
	}
	db := global.GVA_DB
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()
	nft.Address = strings.ToLower(transfer.To)
	err = tx.Save(&nft).Error
	if err != nil {
		return
	}
	err = tx.Create(&transfer).Error
	return err
}

// GetLukNftTransfer 根据id获取LukNftTransfer记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftTransferService *LukNftTransferService) GetLukNftTransfer(id uint) (lukNftTransfer luk.LukNftTransfer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lukNftTransfer).Error
	return
}

// GetLukNftTransferInfoList 分页获取LukNftTransfer记录
// Author [piexlmax](https://github.com/piexlmax)
func (lukNftTransferService *LukNftTransferService) GetLukNftTransferInfoList(info lukReq.LukNftTransferSearch) (list []response.LukNftTransferReponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Debug().Table((&luk.LukNftTransfer{}).TableName()+" transfer").
		Select("transfer.from, transfer.nft_id, transfer.to, transfer.tx_hash, transfer.created_at transfer_at, nft.*").
		Joins("LEFT JOIN luk_nft nft ON nft.id = transfer.nft_id", "nft")
	var lukNftTransfers []response.LukNftTransferReponse
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("transfer.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.To != "" {
		db.Where("transfer.to = ?", info.To)
	}
	if info.From != "" {
		db.Where("transfer.from = ?", info.From)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("transfer.created_at desc").Find(&lukNftTransfers).Error
	return lukNftTransfers, total, err
}
