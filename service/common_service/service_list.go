package common_service

import (
	"blog_server/global"
	"blog_server/models"
	"fmt"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string
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
	for index, column := range op.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", op.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", op.Key))
	}

	count = DB.Where(model).Find(&list).RowsAffected
	query := DB.Where(model) // reset

	// caculate page
	offset := (op.PageInfo.Page - 1) * op.PageInfo.Limit

	if offset < 0 {
		offset = 0
	}

	// store search result to list
	if op.Limit == 0 { // present all data； before gorm 1.25. gorm will ignore limit = 0
		err = query.Order(op.Sort).Find(&list).Error
	} else { // present paginated data
		err = query.Limit(op.Limit).Offset(offset).Order(op.Sort).Find(&list).Error
	}
	return list, count, err
}
