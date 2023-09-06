package flag

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/service/user_service"
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

	// if two pwd same
	if password != rePassword {
		global.Logger.Error("The passwords do not match, please re-enter.")
		return
	}

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	} else if permissions == "user" {
		role = ctype.PermissionUser
	}

	err := user_service.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")

	if err != nil {
		global.Logger.Error(err)
		return
	}

	global.Logger.Infof("User %s successfully created!", userName)
}
