package advertise_api

type AdvertiseApi struct {
}

type AdvertiseRequest struct {
	Title  string `json:"title" binding:"required" msg:"Enter a title"`
	Href   string `json:"href" binding:"required,url" msg:"Enter a valid url"`
	Images string `json:"images" binding:"required,url" msg:"Enter a valid image path"`
	IsShow bool   `json:"is_show" binding:"required" msg:"Select show or not"`
}
