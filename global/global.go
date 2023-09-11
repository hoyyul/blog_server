package global

import (
	"blog_server/config"

	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Logger      *logrus.Logger
	MysqlLogger logger.Interface
	Redis       *redis.Client
	ESClient    *elastic.Client
)
