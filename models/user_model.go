package models

import "blog_server/models/ctype"

// UserModel
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name,select(c|info)"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:128" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(c)"`
	Email      string           `gorm:"size:128" json:"email,select(info)"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr,select(c|info)"`
	Token      string           `gorm:"size:64" json:"token"`
	IP         string           `gorm:"size:20" json:"ip,select(c)"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role,select(info)"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status,select(info)`
	Integral   int              `gorm:"default:0" json:"integral,select(info)"`
	Sign       string           `gorm:"size:128" json:"sign,select(info)"`
	Link       string           `gorm:"size:128" json:"link,select(info)"`
}
