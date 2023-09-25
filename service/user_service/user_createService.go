package user_service

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils"
	"blog_server/utils/pwd"
)

const Avatar = "/uploads/avatar/default.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// if user already exists
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return err
	}

	// get a hash value
	hashPwd := pwd.HashPwd(password)

	addr := utils.GetAddr(ip)
	// save to database
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         "127.0.0.1",
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error

	if err != nil {
		return err
	}
	return nil
}
