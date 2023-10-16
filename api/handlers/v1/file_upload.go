package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// File upload
// @Security ApiKeyAuth
// @Router /v1/file-upload [post]
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Security    BearerAuth
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Success 200 {object} string
func (h *handlerV1) UploadFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting ", err)
		return
	}
	if filepath.Ext(file.File.Filename) != ".png" && filepath.Ext(file.File.Filename) != ".jpg" && filepath.Ext(file.File.Filename) != ".jpeg" {
		fmt.Println(filepath.Ext(file.File.Filename) != ".png" || filepath.Ext(file.File.Filename) != ".jpg")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couln't find matching file format",
		})
		h.log.Error("Error while getting uploading img file-upload", err)
		return
	}
	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	if _, err := os.Stat(dst + "/media"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media", os.ModePerm)
	}

	filePath := "/media/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting customer by email post", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}
