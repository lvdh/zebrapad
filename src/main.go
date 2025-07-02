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

	// serve static files from /static/
	r.StaticFS("/static", http.FS(webAssets))

	// serve index.html at root
	r.GET("/", func(c *gin.Context) {
		f, err := webAssets.Open("web/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found")
			return
		}
		defer f.Close()
		c.DataFromReader(http.StatusOK, -1, "text/html", f, nil)
	})

	r.POST("/print", func(c *gin.Context) {
		var req struct {
			ZPL string `form:"zpl" json:"zpl" binding:"required"`
		}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing ZPL input"})
			return
		}
		err := SendZPLToPrinter(req.ZPL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "sent"})
	})

	r.GET("/debug", func(c *gin.Context) {
		files, err := webAssets.ReadDir("web")
		if err != nil {
			c.String(500, "error: %v", err)
			return
		}
		var list string
		for _, f := range files {
			list += f.Name() + "\n"
		}
		c.String(200, list)
	})

	r.Run(":8080")
}
