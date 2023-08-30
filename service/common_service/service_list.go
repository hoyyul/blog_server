package common_service

import (
	"blog_server/global"
	"blog_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// get paginated data
func FetchPaginatedData[T any](model T, op Option) (list []T, count int64, err error) {
	// print sql statement if debug mode
	DB := global.DB
	if op.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogger}) // return a new *gorm.DB instance
	}

	// Sort according to time to create
	if op.Sort == "" {
		op.Sort = "created_at desc"
	}

	DB = DB.Where(model)
	count = DB.Where(model).Find(&list).RowsAffected
	query := DB.Where(model) // reset

	// caculate page
	offset := (op.PageInfo.Page - 1) * op.PageInfo.Limit

	if offset < 0 {
		offset = 0
	}

	// store search result to list
	if op.Limit == 0 { // present all data
		err = query.Offset(offset).Order(op.Sort).Find(&list).Error
	} else { // present paginated data
		err = query.Limit(op.Limit).Offset(offset).Order(op.Sort).Find(&list).Error
	}
	return list, count, err
}
