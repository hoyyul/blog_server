package es_service

import (
	"blog_server/models"
)

type Option struct {
	models.PageInfo          // pagnation info
	Fields          []string // search article by title, abstract, content
	Tag             string   // search article by tag
}

type SortField struct {
	Field     string
	Ascending bool
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
