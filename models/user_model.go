package models

import "blog_server/models/ctype"

// UserModel
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:128" json:"-"` // hash value of pwd, not real pwd
	Avatar     string           `gorm:"size:256" json:"avatar_id"`
	Email      string           `gorm:"size:128" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr"`
	Token      string           `gorm:"size:64" json:"token"`
	IP         string           `gorm:"size:20" json:"ip"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
}
