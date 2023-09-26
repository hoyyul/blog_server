package image_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service/image_service"
	"blog_server/utils"
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)

	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, image_service.ImageWhiteList) {
		res.FailWithMessage("Illegal file", c)
		return
	}

	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("Image size exceeds the set size, current size is: %.2fMB, set size is: %dMB ", size, global.Config.Upload.Size)
		res.FailWithMessage(msg, c)
		return
	}
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData("/"+filePath, c)

}
