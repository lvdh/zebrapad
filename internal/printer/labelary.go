package printer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

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

func (p LabelaryPrinter) Print(data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", p.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("labelary API error: %s", string(body))
	}
	return resp, nil
}
