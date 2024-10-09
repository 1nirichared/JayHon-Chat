package Img_upload_connector

import (
	"JayHonChat/services"
	"JayHonChat/services/img_freeimage"
)

var serveMap = map[string]services.ImgUploadInterface{
	"fi": &img_freeimage.ImgFreeImageService{},
}

func ImgCreate() services.ImgUploadInterface {
	return serveMap["fi"]
}
