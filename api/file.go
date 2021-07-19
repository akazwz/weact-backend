package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file does not upload",
		})
		return
	}
	filename := file.Filename

	filePre := uuid.NewV4().String()
	newFilename := filePre + "-" + filename
	err = c.SaveUploadedFile(file, "public/file/"+newFilename)
	url := "/static-file/" + newFilename
	if err != nil {
		log.Fatal("file save err", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     url,
	})
}
