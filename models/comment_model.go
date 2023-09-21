package models

type CommentModel struct {
	MODEL           `json:",select(c)"`
	SubComments     []CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"`
	ParentComment   *CommentModel  `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID *uint          `json:"parent_comment_id,select(c)"` // can be nil
	Content         string         `gorm:"size:256" json:"content,select(c)"`
	DiggCount       int            `gorm:"size:8;default:0;" json:"digg_count,select(c)"`
	CommentCount    int            `gorm:"size:8;default:0;" json:"comment_count,select(c)"`
	ArticleID       string         `gorm:"size:32" json:"article_id,select(c)"`
	User            UserModel      `gorm:"foreignKey:UserID" json:"user,select(c)"`
	UserID          uint           `json:"user_id,select(c)"`
}
