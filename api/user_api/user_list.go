package user_api

import (
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/common_service"
	"blog_server/utils/jwts"
	"blog_server/utils/mask"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)
	var page UserListRequest

	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// get paginated list
	var users []UserResponse
	list, count, _ := common_service.FetchPaginatedData[models.UserModel](models.UserModel{}, common_service.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})

	// mask information
	for _, user := range list {
		if ctype.Role(claim.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		user.Tel = mask.MaskTel(user.Tel)
		user.Email = mask.MaskEmail(user.Email)
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)
}
