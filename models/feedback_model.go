package models

type FeedbackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"apply_content"`
	IsApply      bool   `json:"is_apply"`
}
