package images_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

type FileUploadResponse struct {
	FileName   string `json:"file_name"`
	IsUploaded bool   `json:"is_uploaded"`
	Msg        string `json:"msg"`
}

func (ImagesApi) ImagesUploadView(c *gin.Context) {
	imageForm, err := c.MultipartForm()

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	imageList, ok := imageForm.File["images"]
	if !ok {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var responseList []FileUploadResponse
	for _, image := range imageList { // image is a file header type
		//filePath := path.Join("uploads", image.Filename)
		size := float64(image.Size) / float64(1024*1024)

		if size > float64(global.Config.Upload.Size) {
			responseList = append(responseList, FileUploadResponse{
				FileName:   image.Filename,
				IsUploaded: false,
				Msg:        fmt.Sprintf("File size exceeds the maximum size %d", global.Config.Upload.Size),
			})
		} else {
			responseList = append(responseList, FileUploadResponse{
				FileName:   image.Filename,
				IsUploaded: true,
				Msg:        "File uploaded successfully!",
			})
		}
	}

	res.OkWithData(responseList, c)
}
