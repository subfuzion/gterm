package driver

import (
	tb "github.com/nsf/termbox-go"
	"github.com/subfuzion/gterm/geometry"
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

func (d driver) Done() {
	tb.Close()
}

func (d driver) Refresh() {
	if err := tb.Flush(); err != nil {
		d.log(err)
	}
}

func (d driver) Size() (width, height int) {
	return tb.Size()
}

func (d driver) Clear(fg, bg style.CellStyle) {
	if err := tb.Clear(tb.Attribute(fg), tb.Attribute(bg)); err != nil {
		d.log(err)
	}
}

func (d driver) Fill(fg, bg style.CellStyle, ch rune) {
	buf := tb.CellBuffer()
	for i := range buf {
		cell := &buf[i]
		cell.Ch = ch
		cell.Fg = tb.Attribute(fg)
		cell.Bg = tb.Attribute(bg)
	}
}

func (d driver) FillRect(rect geometry.Rectangle, fg, bg style.CellStyle, ch rune) {
	buf := tb.CellBuffer()
	width, _ := tb.Size()

	colStart := rect.Left()
	colStop := rect.Right()
	rowStart := rect.Top()
	rowStop := rect.Bottom()

	for row := rowStart; row < rowStop; row++ {
		for col := colStart; col < colStop; col++ {
			i := (row * width) + col
			cell := &buf[i]
			cell.Ch = ch
			cell.Fg = tb.Attribute(fg)
			cell.Bg = tb.Attribute(bg)
		}
	}
}

func (d driver) SetCell(x, y int, ch rune, fg, bg style.CellStyle) {
	tb.SetCell(x, y, ch, tb.Attribute(fg), tb.Attribute(bg))
}
