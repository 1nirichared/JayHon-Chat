package controller

import (
	"JayHonChat/result"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func ImgKrUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failture(result.APIcode.PostFormError, result.APIcode.GetMessage(result.APIcode.PostFormError), c, &err)
		return
	}
	filepath := viper.GetString(`app.upload_file_path`)
	//指定文件或目录元信息
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}
	filename := filepath + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		result.Failture(result.APIcode.UploadFileError,
			result.APIcode.GetMessage(result.APIcode.UploadFileError), c, &err)
		return
	}

}
