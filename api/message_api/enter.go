package message_api

import "time"

type MessageApi struct {
}

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"`
	RevUserID  uint   `json:"rev_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"Enter a user id to search"`
}

type Message struct {
	SendUserID       uint      `json:"send_user_id"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"`
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`
	CreatedAt        time.Time `json:"created_at"`
	MessageCount     int       `json:"message_count"`
}
