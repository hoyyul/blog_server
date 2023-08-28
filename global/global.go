package global

import (
	"blog_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Logger      *logrus.Logger
	MysqlLogger logger.Interface
)
