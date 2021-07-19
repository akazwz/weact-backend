package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"weact-backend/api"
)

func main() {
	fmt.Println("Hello,Gin")
	r := gin.Default()
	r.POST("/file", api.FileUpload)
	err := r.Run()
	if err != nil {
		log.Fatal("run err", err)
	}
}
