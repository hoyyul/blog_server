package flag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils"
	"fmt"
)

func CreateUser(permissions string) {
	// profile info
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)

	fmt.Printf("Enter a username: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter a nickname: ")
	fmt.Scan(&nickName)
	fmt.Printf("Enter an email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter a password: ")
	fmt.Scan(&password)
	fmt.Printf("Enter the password again: ")
	fmt.Scan(&rePassword)

	// if user already exists
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Logger.Error("Username already exists")
		return
	}

	// if two pwd same
	if password != rePassword {
		global.Logger.Error("The passwords do not match, please re-enter.")
		return
	}

	// get a hash value
	hashPwd := utils.HashPwd(password)
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	} else if permissions == "user" {
		role = ctype.PermissionUser
	}

	// set a default avatar
	avatar := "/uploads/avatar/default.png"

	// save to database
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "Internal Network Address",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Infof("User %s successfully created!", userName)
}
