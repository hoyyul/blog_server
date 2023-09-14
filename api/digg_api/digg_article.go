package digg_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/redis_service"

	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	redis_service.Digg(cr.ID)
	res.OkWithMessage("Dig article successfully", c)
}
