package redis_service

type RedisService struct {
}

const prefix = "logout_"

const (
	articleVisitPrefix        = "article_visit"
	articleCommentCountPrefix = "article_comment_count"
	articleDiggPrefix         = "article_digg"
	commentDiggPrefix         = "comment_digg"
)

func NewArticleDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewArticleVisit() CountDB {
	return CountDB{
		Index: articleVisitPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
