package models

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`
	Abstract      string         `json:"abstract"`
	Content       string         `json:"content"`
	LookCount     int            `json:"look_count"`
	CommentCount  int            `json:"comment_count"`
	DiggCount     int            `json:"digg_count"`
	CollectsCount int            `json:"collects_count"`
	Tags          []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"`
	Comments      []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`
	User          UserModel      `gorm:"foreignKey:UserID" json:"-"` //preload Loading user infomation
	UserID        uint           `json:"user_id"`
	Category      string         `gorm:"size:20" json:"category"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`
	BannerID      uint           `json:"cover_id"`
	NickName      string         `gorm:"size:42" json:"nick_name"`
	CoverPath     string         `json:"cover_path"`
	//Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`
}
