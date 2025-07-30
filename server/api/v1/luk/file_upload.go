package luk

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	lukRes "github.com/flipped-aurora/gin-vue-admin/server/model/luk/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type FileUploadApi struct {
}

var FileUploadService = service.ServiceGroupApp.LukServiceGroup.FileUploadService

// @Tags Upload
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件"
// @Success 200 {object} response.Response{data=whoRes.FileUploadResponse,msg=string} "上传图片,返回包括图片详情"
// @Router /fileUpload/upload [post]
func (FileUploadApi *FileUploadApi) Upload(c *gin.Context) {
	file, errLoad := c.FormFile("file")
	if errLoad != nil {
		global.GVA_LOG.Error("获取上传文件错误：" + errLoad.Error())
		response.FailWithMessage("获取上传文件错误："+errLoad.Error(), c)
	}

	f, uploadErr := FileUploadService.UploadFile(file)
	if uploadErr != nil {
		response.FailWithMessage("获取上传文件错误："+uploadErr.Error(), c)
	}
	response.OkWithDetailed(lukRes.FileUploadResponse{File: f}, "上传成功", c)
}

// @Tags Delete
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body who.FileUpload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUpload/delete [post]
func (FileUploadApi *FileUploadApi) Delete(c *gin.Context) {
	var file luk.FileUpload
	_ = c.ShouldBindJSON(&file)
	if err := FileUploadService.DeleteFile(file.Key); err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
