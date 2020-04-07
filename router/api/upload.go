package api

import (
	"net/http"
	"um_sys/pkg/app"
	"um_sys/pkg/e"
	"um_sys/pkg/upload"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	appG := app.Gin{Ctx: c}

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAM, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	savePath := upload.GetImagePath()
	fullPath := upload.GetImageFullPath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.PrepareCreateImage(fullPath)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_PREPARE_IMAGE_FAIL, nil)
		return
	}

	err = c.SaveUploadedFile(image, src)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
