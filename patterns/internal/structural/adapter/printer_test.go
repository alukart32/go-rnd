package adapter

import "testing"

func TestPrinterAdapter(t *testing.T) {
	msg := "Hello from printer"

	oldPrinter := new(MyLegacyPrinter)
	returnedMsg1 := oldPrinter.Print(msg)

	printer := &PrinterAdapter{
		OldPrinter: oldPrinter,
		Msg:        msg,
	}

	returnedMsg2 := printer.PrintSorted()

	if returnedMsg1 != returnedMsg2 {
		t.Errorf("Message didn't match: %s\n", returnedMsg1)
	}
}
