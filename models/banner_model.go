package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}
