package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
)

type FileUploadResponse struct {
	File luk.FileUpload `json:"file"`
}
