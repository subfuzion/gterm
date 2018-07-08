package driver

import (
	"github.com/nsf/termbox-go"
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
	return termbox.Init()
}

func (d driver) End() {
	termbox.Close()
}

func (d driver) Refresh() {
	if err := termbox.Flush(); err != nil {
		d.log(err)
	}
}

func (d driver) Clear(fg, bg style.CellStyle) {
	if err := termbox.Clear(termbox.Attribute(fg), termbox.Attribute(bg)); err != nil {
		d.log(err)
	}
}

