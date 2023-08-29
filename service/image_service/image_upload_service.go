package image_service

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/plugins/aws"
	"blog_server/utils"
	"fmt"
	"io"
	"mime/multipart"
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

func (ImageService) ProcessImage(image *multipart.FileHeader, c *gin.Context) (res FileUploadResponse) {
	imageName := image.Filename
	basePath := global.Config.Upload.Path
	imagePath := path.Join(basePath, imageName)
	res.FileName = imageName

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

	imageFile.Seek(0, io.SeekStart) // reset pointer
	imageHash := utils.Md5(byteData)

	// check if the file type is legal
	if !checkIfFileLegal(imageName) {
		res.Msg = "Illegal file type"
		return
	}

	// check if the file size exceeds the maximum size
	if !checkIfFileOversize(image.Size) {
		res.Msg = fmt.Sprintf("File size exceeds the maximum size %d", global.Config.Upload.Size)
		return
	}

	// check if the image already exists in the database
	if checkIfImageExists(imageHash) {
		res.Msg = "Image already existed"
		return
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
