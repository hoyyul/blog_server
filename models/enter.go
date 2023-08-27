package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
