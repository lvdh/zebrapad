package printer

type Printer interface {
	Print(data []byte) error
}
