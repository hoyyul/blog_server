package common

import (
	"blog_server/global"
	"blog_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func FetchPaginatedData[T any](op Option) (list T, count int64, err error) {
	//print sql statement if debug mode
	DB := global.DB
	if op.Debug {
		DB = global.DB.Session(&gorm.Session{ //return a new *gorm.DB instance
			Logger: global.MysqlLogger,
		})
	}

	// caculate page
	count = DB.Select("id").Find(&list).RowsAffected
	offset := (op.PageInfo.Page - 1) * op.PageInfo.Limit

	if offset < 0 {
		offset = 0
	}

	// store search result to list
	err = DB.Limit(op.PageInfo.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
