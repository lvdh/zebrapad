package main

import (
	"net/http"

	"github.com/lvdh/zebrapad/internal/printer"
	"github.com/lvdh/zebrapad/internal/server"
)

func main() {
	var printerTcpAddr = ":9100"
	var serverHttpAddr = ":8080"

	p := printer.NewTCPPrinter(printerTcpAddr)
	s := server.NewHTTPServer(p)
	http.ListenAndServe(serverHttpAddr, s)
}
