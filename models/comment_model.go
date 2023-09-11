package models

type CommentModel struct {
	MODEL
	SubComments     []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentComment   *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID *uint           `json:"parent_comment_id"` // can be nil
	Content         string          `gorm:"size:256" json:"content"`
	DiggCount       int             `gorm:"size:8;default:0;" json:"digg_count"`
	CommentCount    int             `gorm:"size:8;default:0;" json:"comment_count"`
	ArticleID       string          `gorm:"size:32" json:"article_id"`
	User            UserModel       `gorm:"foreignKey:UserID" json:"user"`
	UserID          uint            `json:"user_id"`
}
