package main

import (
	"fmt"

	"github.com/subfuzion/gterm/terminal"
	"github.com/subfuzion/gterm/style"
	"github.com/subfuzion/gterm/geometry"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	// draw line
	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.DrawLineHorizontal(geometry.MakePoint(0, 0), 5, '*', '-', '*')
	t.DrawLineVertical(geometry.MakePoint(6, 0), 5, '+', '|', '+')
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()

	// restore tty
	t.Done()
}
