package server

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvdh/zebrapad/internal/printer"
)

//go:embed web/*
var webAssets embed.FS

func NewHTTPServer(p printer.Printer) *gin.Engine {
	r := gin.Default()

	// chroot the web/* assets
	webAssetsFS, err := fs.Sub(webAssets, "web")
	if err != nil {
		panic(err)
	}

	// GET /assets
	//   serves web/*
	r.StaticFS("/assets", http.FS(webAssetsFS))

	// GET /
	//   serves index.html
	r.GET("/", func(c *gin.Context) {
		f, err := webAssetsFS.Open("index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found") // HTTP 500
			return
		}
		defer f.Close()
		c.DataFromReader(http.StatusOK, -1, "text/html", f, nil) // HTTP 200
	})

	// POST /print
	//   sends raw ZPL to the label printer, example request:
	//     curl 'http://localhost:8080/print' \
	//       -X POST \
	//       -H 'Content-Type: application/json' \
	//       --data-raw '{"zpl":"^XA^FO50,50^ADN,36,20^FDTest Label^FS^XZ"}'
	r.POST("/print", func(c *gin.Context) {
		var req struct {
			Data string `form:"zpl" json:"zpl" binding:"required"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing data input"}) // HTTP 400
			return
		}
		err = p.Print([]byte(req.Data))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // HTTP 500
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "sent"}) // HTTP 200
	})

	r.POST("/render", func(c *gin.Context) {
		var req struct {
			Data string `form:"zpl" json:"zpl" binding:"required"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing data input"}) // HTTP 400
			return
		}
		// TODO: implement reder call
		c.JSON(http.StatusOK, gin.H{"status": "sent"}) // HTTP 200
	})

	return r
	// r.Run(":8080")
}
