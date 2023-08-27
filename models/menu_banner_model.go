package models

type MenuBannerModel struct {
	MenuID   uint        `json:"menu_id"`
	Menu     MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID uint        `json:"banner_id"`
	Banner   BannerModel `gorm:"foreignKey:BannerID"`
	Sort     int         `gorm:"size:10" json:"sort"`
}
