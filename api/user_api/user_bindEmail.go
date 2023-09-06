package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugins/email"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"blog_server/utils/verification"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"Email illegal"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	// first time
	var req BindEmailRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	session := sessions.Default(c)
	if req.Code == nil {
		// generate verification code and save in session
		code := verification.GetRandomCode()
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("session error", c)
			return
		}

		// send verification code to user
		err = email.NewVerification().Send(req.Email, "Verification code: "+code)
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("session error", c)
			return
		}
		res.OkWithMessage("Please check verification code in your email!", c)
		return
	}

	// Second time
	code := session.Get("valid_code")
	// check valid_code
	if code != *req.Code {
		res.FailWithMessage("Verification code incorrect", c)
		return
	}

	// Save email and pwd to database
	var user models.UserModel
	err = global.DB.Take(&user, claim.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exist", c)
		return
	}
	hashPwd := pwd.HashPwd(req.Password)

	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    req.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to bind email", c)
		return
	}

	res.OkWithMessage("Bind Email successfully", c)
}
