package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed web/*
var webAssets embed.FS

func main() {
	r := gin.Default()

	// Serve embedded static files
	r.StaticFS("/static", http.FS(webAssets))

	r.LoadHTMLFiles("../web/index.html")

	r.GET("/", func(c *gin.Context) {
		c.FileFromFS("web/index.html", http.FS(webAssets))
	})

	r.POST("/print", func(c *gin.Context) {
		var req struct {
			ZPL string `form:"zpl" json:"zpl" binding:"required"`
		}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ZPL input"})
			return
		}
		err := SendZPLToPrinter(req.ZPL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "sent"})
	})

	r.Run(":8080")
}
