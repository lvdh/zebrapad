package main

import (
	"fmt"
	"net"
)

var PrinterAddress = "192.168.1.100:9100" // TODO: make configurable

func SendZPLToPrinter(zpl string) error {
	conn, err := net.Dial("tcp", PrinterAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to printer: %w", err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(zpl))
	if err != nil {
		return fmt.Errorf("failed to send ZPL: %w", err)
	}
	return nil
}
