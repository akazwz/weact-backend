package api

import "github.com/gin-gonic/gin"

type HotSearch struct {
}

func CreateHotSearch(c *gin.Context) {
	hs := HotSearch{}
	c.ShouldBindJSON(hs)
}

func WriteHotSearchToInflux() {

}
