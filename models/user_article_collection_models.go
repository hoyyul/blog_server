package models

import "time"

type UserArticleCollectionModels struct {
	UserID    uint      `gorm:"primaryKey"`
	User      UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32"`
	CreatedAt time.Time
}
