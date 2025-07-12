package main

import (
	"github.com/lvdh/zebrapad/internal/printer"
	"github.com/lvdh/zebrapad/internal/server"
)

func main() {
	p := printer.NewTCPPrinter(":9100")
	s := server.NewHTTPServer(p)
	s.Run(":8080")
}
