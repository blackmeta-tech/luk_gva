package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukRes "github.com/flipped-aurora/gin-vue-admin/server/model/luk/response"
	"gorm.io/gorm"
	"time"
)

type LukPoolBasicService struct {
}

// 更新or创建数据
func (LukPoolBasicService) Update(tx *gorm.DB, data luk.LukPoolBasic) (err error) {
	basic := luk.LukPoolBasic{}
	basic.ID = 1
	data.UpdatedAt = time.Now()
	err = tx.Debug().Where(basic).Assign(data).FirstOrCreate(&basic).Error
	return
}

// 获取地址数据
func (LukPoolBasicService) GetData() (data luk.LukPoolBasic, err error) {
	err = global.GVA_DB.Where("id = 1").First(&data).Error
	return
}

func (LukPoolBasicService) GetDataDapp() (data lukRes.PoolBasicInfo, err error) {
	err = global.GVA_DB.Table((luk.LukPoolBasic{}).TableName()).Where("id = 1").First(&data).Error
	return
}
