package models

type TagModel struct {
	MODEL
	Title         string         `json:"title"`
	ArticleModels []ArticleModel `gorm:"many2many:article_tag_models" json:"-"`
}
