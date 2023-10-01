package image_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImageApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
}
