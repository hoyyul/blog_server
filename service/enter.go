package service

import (
	"blog_server/service/image_service"
	"blog_server/service/user_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
	UserService  user_service.UserService
}

var ServiceApp = new(ServiceGroup)
