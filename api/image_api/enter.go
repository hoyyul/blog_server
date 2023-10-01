package image_api

type ImageApi struct {
}

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"Enter image id"`
	Name string `json:"name" binding:"required" msg:"Enter new image name"`
}
