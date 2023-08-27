package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
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

	// if path not exists, make the dir
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Logger.Error(err)
		}
	}

	//process images
	var responseList []FileUploadResponse
	for _, image := range imageList { // image is a file header type
		imagePath := path.Join(basePath, image.Filename)
		responseList = append(responseList, processImage(image, imagePath, c))
	}

	res.OkWithData(responseList, c)
}

func processImage(image *multipart.FileHeader, imagePath string, c *gin.Context) FileUploadResponse {
	//check if the file type is legal
	if !checkFileType(image.Filename) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        "Illegal file type",
		}
	}

	//check if the file size exceeds the maximum size
	if !checkFileSize(image.Size) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        fmt.Sprintf("File size exceeds the maximum size %d", global.Config.Upload.Size),
		}
	}

	//check if the image already exists in the database
	imageFile, err := image.Open()
	if err != nil {
		global.Logger.Error(err)
	}
	byteData, err := io.ReadAll(imageFile)
	if err != nil {
		global.Logger.Error(err)
	}
	imageHash := utils.Md5(byteData)

	if checkIfImageExists(imageHash) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        "Image already existed",
		}
	}

	//save images to path and database
	err = c.SaveUploadedFile(image, imagePath)
	if err != nil {
		global.Logger.Error(err)
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        err.Error(),
		}
	}

	global.DB.Create(&models.BannerModel{
		Path: imagePath,
		Hash: imageHash,
		Name: image.Filename,
	})

	return FileUploadResponse{
		FileName:   image.Filename,
		IsUploaded: true,
		Msg:        "File uploaded successfully!",
	}
}

func checkFileType(filename string) bool {
	nameList := strings.Split(filename, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	return utils.InList(suffix, ImageWhiteList)
}

func checkFileSize(size int64) bool {
	return float64(size)/float64(1024*1024) <= float64(global.Config.Upload.Size)
}

func checkIfImageExists(hash string) bool {
	var banner models.BannerModel
	err := global.DB.Take(&banner, "hash = ?", hash).Error
	return err == nil
}
