package synchro_service

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/redis_service"

	"github.com/sirupsen/logrus"
)

func SyncCommentData() {
	commentDiggInfo := redis_service.NewCommentDigg().GetInfo()

	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = global.DB.Model(&comment).Update("digg_count", count).Error
		if err != nil {
			logrus.Error(err)
			continue
		}
		logrus.Infof("%s updated successfully, new like count is %d", comment.Content, comment.DiggCount)

	}
	redis_service.NewCommentDigg().Clear()
}
