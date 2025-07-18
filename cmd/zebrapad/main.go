package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/lvdh/zebrapad/internal/printer"
	"github.com/lvdh/zebrapad/internal/server"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	var tcpEnv = getEnvOrDefault("ZEBRAPAD_TCP", ":9100")
	var tcpAddr = flag.String("tcp", tcpEnv, "Printer TCP address ([host]:port)")
	var httpEnv = getEnvOrDefault("ZEBRAPAD_HTTP", ":8080")
	var httpAddr = flag.String("http", httpEnv, "Server address ([host]:port)")
	flag.Parse()

	p := printer.NewTCPPrinter(*tcpAddr)
	s := server.NewHTTPServer(p)
	http.ListenAndServe(*httpAddr, s)
}
