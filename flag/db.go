package flag

import (
	"blog_server/global"
	"blog_server/models"
)

func MigrateTables() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "UserCollections", &models.UserArticleCollectionModels{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertiseModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.UserArticleCollectionModels{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FeedbackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Logger.Error("[ error ] Table schema migration failed.")
		return
	}
	global.Logger.Info("[ success ] Table schema migration successful.")
}
