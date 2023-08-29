package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/plugins/aws"
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

	// process images
	var responseList []FileUploadResponse
	for _, image := range imageList { // image is a file header type
		imagePath := path.Join(basePath, image.Filename)
		responseList = append(responseList, processImage(image, imagePath, c))
	}

	res.OkWithData(responseList, c)
}

func processImage(image *multipart.FileHeader, imagePath string, c *gin.Context) FileUploadResponse {
	// get image hash value used in database
	imageFile, err := image.Open() // return *os.file, err
	if err != nil {
		global.Logger.Error(err)
	}
	byteData, err := io.ReadAll(imageFile)
	if err != nil {
		global.Logger.Error(err)
	}
	defer imageFile.Close()

	// reset pointer
	imageFile.Seek(0, io.SeekStart)
	imageHash := utils.Md5(byteData)

	// check if the file type is legal
	if !checkIfFileLegal(image.Filename) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        "Illegal file type",
		}
	}

	// check if the file size exceeds the maximum size
	if !checkIfFileOversize(image.Size) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        fmt.Sprintf("File size exceeds the maximum size %d", global.Config.Upload.Size),
		}
	}

	// check if the image already exists in the database
	if checkIfImageExists(imageHash) {
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        "Image already existed",
		}
	}

	// upload according to location(local/aws)
	var storageLocation ctype.StorageLocation
	if checkIfStorageCloud() {
		storageLocation = ctype.AWS
	} else {
		storageLocation = ctype.Local
	}

	return handleImageStorage(image, imageFile, imagePath, imageHash, storageLocation, c)
}

func handleImageStorage(image *multipart.FileHeader, imageFile multipart.File, imagePath, imageHash string, storageLocation ctype.StorageLocation, c *gin.Context) FileUploadResponse {
	var err error
	var storagePath string

	if storageLocation == ctype.AWS {
		// upload to aws
		objKey := path.Join("images", image.Filename)
		err = aws.UploadFile(objKey, imageFile)
		if err == nil {
			storagePath = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", global.Config.AWS.Bucket, global.Config.AWS.Region, objKey)
		}
	} else {
		// save to local
		err = c.SaveUploadedFile(image, imagePath)
		if err == nil {
			storagePath = imagePath
		}
	}

	if err != nil {
		global.Logger.Error(err)
		return FileUploadResponse{
			FileName:   image.Filename,
			IsUploaded: false,
			Msg:        err.Error(),
		}
	}

	// save to database
	global.DB.Create(&models.BannerModel{
		Path:            storagePath,
		Hash:            imageHash,
		Name:            image.Filename,
		StorageLocation: storageLocation,
	})

	return FileUploadResponse{
		FileName:   storagePath,
		IsUploaded: true,
		Msg:        "File uploaded successfully",
	}
}

func checkIfFileLegal(filename string) bool {
	nameList := strings.Split(filename, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	return utils.InList(suffix, ImageWhiteList)
}

func checkIfFileOversize(size int64) bool {
	return float64(size)/float64(1024*1024) <= float64(global.Config.Upload.Size)
}

func checkIfImageExists(hash string) bool {
	var banner models.BannerModel
	err := global.DB.Take(&banner, "hash = ?", hash).Error
	return err == nil
}

func checkIfStorageCloud() bool {
	return global.Config.AWS.Enable
}
