package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strings"
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
	extensionArr := strings.SplitAfter(filename, ".")
	index := len(extensionArr)
	var extension = ""
	if index > 1 {
		extension = extensionArr[index-1]
	}
	filePre := uuid.NewV4().String()
	newFilename := filePre + "-" + filename + "." + extension
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
