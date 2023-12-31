package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)"`
	CreatedAt time.Time `json:"created_at,select($any)"`
	UpdatedAt time.Time `json:"-"`
}

// request for Pagination
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

// request for Remove
type RemoveRequest struct {
	IdList []int `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}
