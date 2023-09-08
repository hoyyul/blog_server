package message_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/utils/jwts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageReadUserView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var messageList []models.MessageModel
	global.DB.Find(&messageList, "send_user_id = ? or rev_user_id = ?", claim.UserID, claim.UserID)
	fmt.Println(messageList)
}
