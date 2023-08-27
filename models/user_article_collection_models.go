package models

import "time"

type UserArticleCollectionModels struct {
	UserID    uint         `gorm:"primaryKey"`
	User      UserModel    `gorm:"foreignKey:UserID"`
	ArticleID uint         `gorm:"primaryKey"`
	Article   ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time
}
