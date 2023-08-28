package models

import (
	"blog_server/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path             string                 `json:"path"`
	Hash             string                 `json:"hash"`
	Name             string                 `json:"name"`
	StorageLocatioin ctype.StorageLocatioin `gorm:"default:1" json:"storage_location"` // local or cloud
}

// hook function to remove image from local path
func (b *BannerModel) BeforeDelete(db *gorm.DB) (err error) {
	if b.StorageLocatioin == ctype.Local {
		err := os.Remove(b.Path)
		if err != nil {
			return err
		}
	}
	return nil
}
