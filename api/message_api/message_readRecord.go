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
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidation(err, &cr, c)
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
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	res.OkWithData(messageList, c)
}
