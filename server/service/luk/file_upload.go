package luk

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"strings"
)

type FileUploadService struct {
}

func (e *FileUploadService) UploadFile(header *multipart.FileHeader) (file luk.FileUpload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		err = uploadErr
		return
	}
	s := strings.Split(header.Filename, ".")
	f := luk.FileUpload{
		Url:  "/api/" + filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	file = f
	return
}

func (e *FileUploadService) DeleteFile(key string) (err error) {
	oss := upload.NewOss()
	if err = oss.DeleteFile(key); err != nil {
		return errors.New("文件删除失败")
	}
	return err
}
