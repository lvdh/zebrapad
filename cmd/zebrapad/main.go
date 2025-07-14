package main

import (
	"net/http"

	"github.com/lvdh/zebrapad/internal/printer"
	"github.com/lvdh/zebrapad/internal/server"
)

func main() {
	p := *printer.NewTCPPrinter(":9100")
	s := server.NewHTTPServer(p)
	http.ListenAndServe(":8080", s)
}
