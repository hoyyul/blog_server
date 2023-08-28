package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
