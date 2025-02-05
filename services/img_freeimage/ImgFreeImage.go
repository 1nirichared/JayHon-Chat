package img_freeimage

import (
	"JayHonChat/result"
	"JayHonChat/services"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
)

type ImgFreeImageService struct {
	services.ImgUploadInterface
}

func (serve *ImgFreeImageService) Upload(fielname string) string {
	return Upload(fielname)
}

func Upload(uploadFile string) string {
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	bodyWriter.WriteField("type", "file")
	bodyWriter.WriteField("action", "upload")
	//从bodywriter生成filewriter，并将文件内容写入file writer，多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("source", path.Base(uploadFile))
	if err != nil {
		log.Println(err)
		result.Failture(result.APIcode.FileWirteError,
			result.APIcode.GetMessage(result.APIcode.FileWirteError), nil, &err)
		return ""
	}
	file, err2 := os.Open(uploadFile)
	if err2 != nil {
		log.Println(err2)
		result.Failture(result.APIcode.OpenFileError, result.
			APIcode.GetMessage(result.APIcode.OpenFileError), nil, &err)
		return ""
	}
	defer file.Close()
	_, err3 := io.Copy(fileWriter, file)
	if err3 != nil {
		log.Println(err3)
		result.Failture(result.APIcode.CopyFileError,
			result.APIcode.GetMessage(result.APIcode.CopyFileError), nil, &err)
		return ""
	}
	//关闭bodywriter停止写入数据
	bodyWriter.Close()
	contenType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	request := fasthttp.AcquireRequest()   //创建请求对象
	response := fasthttp.AcquireResponse() //创建响应对象
	defer func() {
		fasthttp.ReleaseResponse(response)
		fasthttp.ReleaseRequest(request)
	}()
	request.Header.SetContentType(contenType)
	//直接将构建好的数据放入post的body中
	request.SetBody(bodyBufer.Bytes())
	request.Header.SetMethod("POST")
	request.SetRequestURI("https://freeimage.host/json")
	err4 := fasthttp.Do(request, response)
	if err4 != nil {
		log.Println(err4)
		result.Failture(result.APIcode.RequsetError,
			result.APIcode.GetMessage(result.APIcode.RequsetError), nil, &err)
		return ""
	}
	var res map[string]interface{}
	e := json.Unmarshal(response.Body(), &res)
	if e != nil {
		log.Println(e, string(response.Body()))
		return ""
	}
	if _, ok := res["image"]; ok {
		if _, set := res["image"].(map[string]interface{})["display_url"]; set {
			return res["image"].(map[string]interface{})["display_url"].(string)
		}
	} else {
		log.Println(res)
	}
	return ""
}
