package api

import (
	"github.com/gin-gonic/gin"
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
	filesize := file.Size
	err = c.SaveUploadedFile(file, "public/"+filename)
	if err != nil {
		log.Fatal("file save err", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"filename": filename,
		"filesize": filesize,
	})
}
