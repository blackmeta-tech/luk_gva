package luk

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UploaderRouter struct {
}

//InitFileUploaderApi 初始化
func (u *UploaderRouter) InitFileUploaderApi(Router *gin.RouterGroup) {
	Uploader := Router.Group("fileUpload").Use(middleware.OperationRecord())
	var UploaderApi = v1.ApiGroupApp.LukApiGroup.FileUploadApi
	{
		Uploader.POST("upload", UploaderApi.Upload)
		Uploader.POST("delete", UploaderApi.Delete)
	}
	Router.Group("apih5/fileUpload").POST("upload", UploaderApi.Upload)
}
