package service

import "blog_server/service/image_service"

type ServiceGroup struct {
	ImageService image_service.ImageService
}

var ServiceApp = new(ServiceGroup)
