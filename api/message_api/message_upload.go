package message_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageUploadView(c *gin.Context) {
	var req MessageRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	var sendUser, recvUser models.UserModel

	err = global.DB.Take(&sendUser, req.SendUserID).Error
	if err != nil {
		res.FailWithMessage("Sender id doesn't exist", c)
		return
	}
	err = global.DB.Take(&recvUser, req.RevUserID).Error
	if err != nil {
		res.FailWithMessage("Receiver id doesn't exist", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       req.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        req.RevUserID,
		RevUserNickName:  recvUser.NickName,
		RevUserAvatar:    recvUser.Avatar,
		IsRead:           false,
		Content:          req.Content,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to send message", c)
		return
	}
	res.OkWithMessage("Send message successfullt", c)
}
