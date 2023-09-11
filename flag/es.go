package flag

import "blog_server/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
