package printer

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
