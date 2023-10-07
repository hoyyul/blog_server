package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/plugins/log_stash"
	"blog_server/utils"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"fmt"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"Enter a username"`
	Password string `json:"password" binding:"required" msg:"Enter a password"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var req EmailLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	logger := log_stash.NewLogByGin(c)

	// validate username
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", req.UserName, req.UserName).Error
	if err != nil {
		logger.Warn(fmt.Sprintf("%s doesn't exist", req.UserName))
		res.FailWithMessage("Username or password incorrect", c)
		return
	}

	// validate password
	isCheck := pwd.CheckPwd(userModel.Password, req.Password)
	if !isCheck {
		logger.Warn(fmt.Sprintf("Username or password incorrect %s %s", req.UserName, req.Password))
		res.FailWithMessage("Username or password incorrect", c)
		return
	}

	// create a token
	token, err := jwts.GenerateToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Avatar:   userModel.Avatar,
	})

	if err != nil {
		global.Logger.Error(err)
		logger.Error(fmt.Sprintf("Failed to generate a token %s", err.Error()))
		res.FailWithMessage("Failed to generate a token", c)
		return
	}

	// log
	logger = log_stash.New(c.ClientIP(), token)
	ip, addr := utils.GetAddrByGin(c)
	logger.Info("Login successfully!")

	// save login record to db
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})

	res.OkWithData(token, c)

}
