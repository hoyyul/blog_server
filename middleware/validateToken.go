package middleware

import (
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		// check if empty
		if token == "" {
			res.FailWithMessage("Request is missing a token", c)
			c.Abort()
			return
		}

		// check if token valid
		claim, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("No information found for the given token", c)
			c.Abort()
			return
		}

		// check if in redis
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token expired", c)
			c.Abort()
			return
		}

		// set key
		c.Set("claim", claim)
	}
}

func CheckAdminToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		// check if empty
		if token == "" {
			res.FailWithMessage("Request is missing a token", c)
			c.Abort()
			return
		}

		// check if token valid
		claim, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("No information found for the given token", c)
			c.Abort()
			return
		}

		// check role
		if claim.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("Permission doesn't macth", c)
			c.Abort()
			return
		}

		// check if in redis
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token expired", c)
			c.Abort()
			return
		}

		// set key
		c.Set("claim", claim)
	}
}
