package message_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

// find chat history between user a and user b
func (MessageApi) MessageReadRecordView(c *gin.Context) {
	var req MessageRecordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)

	// find user a message
	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claim.UserID, claim.UserID)

	// find user b message
	for _, model := range _messageList {
		if model.RevUserID == req.UserID || model.SendUserID == req.UserID {
			messageList = append(messageList, model)
		}
	}
	res.OkWithData(messageList, c)
}
