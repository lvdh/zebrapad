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

type LabelaryPrinter struct {
	url     string  `default:"http://api.labelary.com/v1/printers"`
	density string  `default:"8dpmm"`
	height  float32 // in inches *cries in bald eagle*
	width   float32 // in inches *cries in bald eagle*
}

func NewLabelaryPrinter() *LabelaryPrinter {
	p := &LabelaryPrinter{
		width:  2.25,
		height: 0.75,
	}
	return p
}

func (p LabelaryPrinter) Print(data []byte) error {
	return nil
}
