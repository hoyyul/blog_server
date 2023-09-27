package image_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"
	"blog_server/service/image_service"
	"blog_server/utils/jwts"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

func (ImageApi) ImageCreateListView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)
	if claim.Role == 3 {
		res.FailWithMessage("Sign in to upload image", c)
		return
	}

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
	var responseList []image_service.FileUploadResponse
	for _, image := range imageList { // image is a file header type
		responseList = append(responseList, service.ServiceApp.ImageService.ProcessImage(image, c))
	}

	res.OkWithData(responseList, c)
}
