package news_api

import (
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"blog_server/utils"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewsResponse struct {
	Code int                      `json:"code"`
	Data []redis_service.NewsData `json:"data"`
	Msg  string                   `json:"msg"`
}

const newAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

func (NewsApi) NewsReadListView(c *gin.Context) {
	var req params
	var headers header
	c.ShouldBindJSON(&req)
	err := c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	if req.Size == 0 {
		req.Size = 1
	}

	// redis key
	key := fmt.Sprintf("%s-%d", req.ID, req.Size)

	// check if exist in redis
	newsData, _ := redis_service.GetNews(key)
	if len(newsData) != 0 {
		res.OkWithData(newsData, c)
		return
	}

	// http request for news API
	httpResponse, err := utils.Post(newAPI, req, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewsResponse
	byteData, _ := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}

	// response
	res.OkWithData(response.Data, c)

	// cache in redis
	redis_service.SetNews(key, response.Data)
}
