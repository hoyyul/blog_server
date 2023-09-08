package message_api

type MessageApi struct {
}

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"`
	RevUserID  uint   `json:"rev_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}
