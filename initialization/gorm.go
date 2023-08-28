package initialization

import (
	"blog_server/global"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Logger.Warnln("No configuration host.")
		return nil
	}

	dsn := global.Config.Mysql.DSN()

	var mysqlLogger logger.Interface // logger from gorm lib

	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	global.MysqlLogger = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logger.Fatalf(fmt.Sprintf("[%s] Mysql connection failed.", dsn))
	}
	sqlDB, _ := db.DB() // db is *DB type, db.DB() is *sql.DB type
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
