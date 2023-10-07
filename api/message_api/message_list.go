package message_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type MessageGroup map[uint]*Message

// find chat history between one user and other users
func (MessageApi) MessageListView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var messageList []models.MessageModel
	var messageGroup = MessageGroup{} //init a map
	var messages []Message

	// get message a certain user involved
	global.DB.Order("created_at asc").Find(&messageList, "send_user_id = ? or rev_user_id = ?", claim.UserID, claim.UserID)
	for _, model := range messageList {
		// get the message
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}

		idPair := model.SendUserID + model.RevUserID

		// if such a combination exists, count plus one
		if val, ok := messageGroup[idPair]; ok {
			message.MessageCount = val.MessageCount + 1
		}

		// update to latest message
		messageGroup[idPair] = &message
	}

	// save message to a list
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)
}
