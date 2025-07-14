package printer

import (
	"fmt"
	"net"
)

type Printer interface {
	Print(data []byte) error
}

type TCPPrinter struct {
	Network string
	Address string // "[host]:port"
}

func NewTCPPrinter(address string) *TCPPrinter {
	p := &TCPPrinter{
		Network: "tcp",
		Address: address,
	}
	return p
}

func (p TCPPrinter) Print(data []byte) error {
	conn, err := net.Dial(p.Network, p.Address)
	if err != nil {
		return fmt.Errorf("failed to connect to printer: %w", err)
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to print: %w", err)
	}
	return nil
}
