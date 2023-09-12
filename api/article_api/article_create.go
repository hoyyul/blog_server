package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"math/rand"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"Enter a title"`
	Abstract string      `json:"abstract"`
	Content  string      `json:"content" binding:"required" msg:"Enter the content"`
	Category string      `json:"category"`
	Source   string      `json:"source"`
	Link     string      `json:"link"`
	BannerID uint        `json:"banner_id"`
	Tags     ctype.Array `json:"tags"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var req ArticleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)
	userID := claim.UserID
	userNickName := claim.NickName

	// remove script from html
	unsafe := blackfriday.MarkdownCommon([]byte(req.Content))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		req.Content = markdown
	}

	if req.Abstract == "" {
		// get chinese character
		abs := []rune(req.Content)
		if len(abs) > 100 {
			req.Abstract = string(abs[:100])
		} else {
			req.Abstract = string(abs)
		}
	}

	// random select a banner if none
	if req.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("No banner data", c)
			return
		}
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		req.BannerID = bannerIDList[r.Intn(len(bannerIDList))]
	}

	// find banner
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", req.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner doesn't exist", c)
		return
	}

	// Find avatar
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("User doesn't exist", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        req.Title,
		Keyword:      req.Title,
		Abstract:     req.Abstract,
		Content:      req.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     req.Category,
		Source:       req.Source,
		Link:         req.Link,
		BannerID:     req.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         req.Tags,
	}

	// create article
	if article.ISExistData() {
		res.FailWithMessage("Article already exists", c)
		return
	}
	err = article.Create()
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("Publish article successfully", c)

}
