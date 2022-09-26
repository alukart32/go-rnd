package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintSorted() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintSorted() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = p.OldPrinter.Print(p.Msg)
	} else {
		newMsg = p.Msg
	}
	return
}
