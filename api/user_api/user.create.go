package user_api

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/user_service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"Enter a nickname"`
	UserName string     `json:"user_name" binding:"required" msg:"Enter a username"`
	Password string     `json:"password" binding:"required" msg:"Enter your password"`
	Role     ctype.Role `json:"role" binding:"required" msg:"Enter the user rule"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	err := user_service.UserService{}.CreateUser(req.UserName, req.NickName, req.Password, req.Role, "", c.ClientIP())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("User %s successfully created!", req.UserName), c)
}
