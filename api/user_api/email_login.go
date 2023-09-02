package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"

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

	// validate username
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", req.UserName, req.UserName).Error
	if err != nil {
		res.FailWithMessage("Username or password incorrect", c)
		return
	}

	// validate password
	isCheck := pwd.CheckPwd(userModel.Password, req.Password)
	if !isCheck {
		res.FailWithMessage("Username or password incorrect", c)
		return
	}

	// create a token
	token, err := jwts.GenerateToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to generate a token", c)
		return
	}
	res.OkWithData(token, c)

}
