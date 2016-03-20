package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	ROUTER := gin.Default()
	ROUTER.LoadHTMLFiles("layout.html", "index.html")
	ROUTER.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	ROUTER.Run(":6000")
}
