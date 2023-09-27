package image_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service/image_service"
	"blog_server/utils"
	"blog_server/utils/jwts"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)
	if claim.Role == 3 {
		res.FailWithMessage("Sign in to upload image", c)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	//filePath := path.Join(basePath, fileName)

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

	// create /uploads/file/nick_name
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		res.FailWithMessage("Path doesn't exist", c)
		return
	}
	if !isInDirEntry(dirList, claim.NickName) {
		err := os.Mkdir(path.Join(basePath, claim.NickName), 077)
		if err != nil {
			global.Logger.Error(err)
		}
	}
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	filePath := path.Join(basePath, claim.NickName, fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData("/"+filePath, c)

}

func isInDirEntry(dirList []os.DirEntry, name string) bool {
	for _, entry := range dirList {
		if entry.Name() == name && entry.IsDir() {
			return true
		}
	}
	return false
}
