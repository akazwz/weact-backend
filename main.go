package main

import (
	"fmt"
	"time"
	"weact-backend/influxdb"
)

func main() {
	fmt.Println("Hello,Gin")
	/*r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))
	r.POST("/file", api.FileUpload)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Static("/file", "./public/file")
	err := r.Run(":8000")
	if err != nil {
		log.Fatal("run err", err)
	}*/
	measurement := "stat"
	tags := make(map[string]string)
	tags["unit"] = "temperature"
	fields := make(map[string]interface{})
	fields["avg"] = 50
	fields["max"] = 100

	influxdb.WritePointData(measurement, tags, fields, time.Now())
}
