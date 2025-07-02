package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed web/*
var webAssets embed.FS

func main() {
	r := gin.Default()

	// chroot the embedded FS to the web/ directory
	webFS, err := fs.Sub(webAssets, "web")
	if err != nil {
		panic(err)
	}

	// serve static files from /static/
	r.StaticFS("/static", http.FS(webFS))

	// serve index.html at root
	r.GET("/", func(c *gin.Context) {
		f, err := webFS.Open("index.html")
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

	r.Run(":8080")
}
