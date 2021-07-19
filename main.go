package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"weact-backend/api"
)

func main() {
	fmt.Println("Hello,Gin")
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))
	r.POST("/file", api.FileUpload)
	r.Static("/file", "./public/file")
	err := r.Run()
	if err != nil {
		log.Fatal("run err", err)
	}
}
