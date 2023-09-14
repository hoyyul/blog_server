package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`
	Abstract string   `json:"abstract"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Source   string   `json:"source"`
	Link     string   `json:"link"`
	BannerID uint     `json:"banner_id"`
	Tags     []string `json:"tags"`
	ID       string   `json:"id"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var req ArticleUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithValidation(err, &req, c)
		return
	}

	// get banner url by id
	var bannerUrl string
	if req.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", req.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			res.FailWithMessage("Banner doesn't exist", c)
			return
		}
	}

	// save update info to map
	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     req.Title,
		Keyword:   req.Title,
		Abstract:  req.Abstract,
		Content:   req.Content,
		Category:  req.Category,
		Source:    req.Source,
		Link:      req.Link,
		BannerID:  req.BannerID,
		BannerUrl: bannerUrl,
		Tags:      req.Tags,
	}

	//
	err = article.ISExistByID(req.ID)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Article doesn't exist", c)
		return
	}

	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	// update elastic search
	err = es_service.ArticleUpdate(req.ID, maps)
	if err != nil {
		res.FailWithMessage("Failed to update article", c)
		return
	}

	res.OkWithMessage("Update article successfully", c)
}
