package images_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/utils"
	"fmt"
	"io/fs"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ImageWhiteList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName   string `json:"file_name"`
	IsUploaded bool   `json:"is_uploaded"`
	Msg        string `json:"msg"`
}

func processImage(image *multipart.FileHeader) FileUploadResponse {
	nameList := strings.Split(image.Filename, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])

	if !utils.InList(suffix, ImageWhiteList) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        "Illegal file type",
		}
	}

	size := float64(image.Size) / float64(1024*1024)
	if size > float64(global.Config.Upload.Size) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        fmt.Sprintf("File size exceeds the maximum size %d", global.Config.Upload.Size),
		}
	}

	return FileUploadResponse{
		FileName:   image.Filename,
		IsUploaded: true,
		Msg:        "File uploaded successfully!",
	}
}

func (ImagesApi) ImagesUploadView(c *gin.Context) {
	imageForm, err := c.MultipartForm()

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	imageList, ok := imageForm.File["images"]
	if !ok {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}

	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Logger.Error(err)
		}
	}

	var responseList []FileUploadResponse
	for _, image := range imageList { // image is a file header type
		responseList = append(responseList, processImage(image))
	}

	res.OkWithData(responseList, c)
}
