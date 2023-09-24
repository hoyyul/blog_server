package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/plugins/qq"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("No code", c)
		return
	}

	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID

	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	if err != nil {
		// if user doesn't exist; register it
		hashPwd := pwd.HashPwd(pwd.GetRandomPwd(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,
			Password:   hashPwd, // 16 digit random pwd
			Avatar:     qqInfo.Avatar,
			Addr:       "127.0.0.1",
			Token:      openID,
			IP:         c.ClientIP(),
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("Failed to register", c)
			return
		}

	}

	// login
	token, err := jwts.GenerateToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to generate token", c)
		return
	}

	// save login record to db
	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        c.ClientIP(),
		NickName:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      "Internal Network",
		LoginType: ctype.SignQQ,
	})

	res.OkWithData(token, c)
}
