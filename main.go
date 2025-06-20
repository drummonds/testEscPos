package main

import (
	"fmt"
	"time"

	"github.com/mect/go-escpos"
)

var p *escpos.Printer

func label(message string, num int) {

	p.Init()       // start
	p.Smooth(true) // use smootth printing
	p.Size(1, 1)   // set font size
	p.Align(escpos.AlignCenter)
	p.PrintLn("Hello Humphrey")

	p.Size(2, 2)
	p.Font(escpos.FontB) // change font
	p.Align(escpos.AlignLeft)
	for len(message) > 0 {
		p.PrintLn(message[:min(30, len(message))])
		if len(message) > 30 {
			message = message[30:]
		} else {
			break
		}

	}
	// p.Font(escpos.FontA)

	// p.Align(escpos.AlignRight) // change alignment
	// p.PrintLn("An all Go\neasy to use\nEpson POS Printer library")
	// p.Align(escpos.AlignLeft)

	// p.Size(2, 2)
	// p.PrintLn("* No magic numbers")
	// p.PrintLn("* ISO8859-15 ŠÙþþØrt")
	// p.Underline(true)
	// p.PrintLn("* Extended layout")
	// p.Underline(false)
	// p.PrintLn("* All in Go!")

	p.Align(escpos.AlignCenter)
	p.Barcode(fmt.Sprintf("%d", num), escpos.BarcodeTypeCODE39) // print barcode
	p.Align(escpos.AlignLeft)

	p.Cut() // cut
	p.End() // stop
	time.Sleep(time.Second * 1)
}

func main() {
	var err error
	p, err = escpos.NewUSBPrinterByPath("") // empry string will do a self discovery
	if err != nil {
		fmt.Println(err)
		return
	}
	// label("Do email", 5)
	label("Talk to exa a 10", 5)
	label("Dry cleaning", 5)
	label("Water garden", 5)
	label("WEb IF for labels", 5)
	label("sell keybaords", 5)
	label("sell toothbrush", 5)
	label("Move chest in bthroom", 5)
	// label("", 5)
	// label("Take FLORA for wallk", 6)
}
