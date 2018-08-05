package driver

import (
	tb "github.com/nsf/termbox-go"
	"github.com/subfuzion/gterm/logging"
	"github.com/subfuzion/gterm/style"
)

func NewTermbox(logger logging.LogPrinter) Driver {
	return &driver{
		log: logger,
	}
}

type driver struct {
	log logging.LogPrinter
}

func (d driver) Init() error {
	return tb.Init()
}

func (d driver) End() {
	tb.Close()
}

func (d driver) Refresh() {
	if err := tb.Flush(); err != nil {
		d.log(err)
	}
}

func (d driver) Clear(fg, bg style.CellStyle) {
	if err := tb.Clear(tb.Attribute(fg), tb.Attribute(bg)); err != nil {
		d.log(err)
	}
}

