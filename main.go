package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "zhuandong")
	})
	r.Run()
}
