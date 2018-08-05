package main

import (
	"fmt"

	"github.com/subfuzion/gterm/geometry"
		"github.com/subfuzion/gterm/terminal"
	"github.com/subfuzion/gterm/style"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	// fill rect
	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.FillRect(geometry.MakeRectangle(2, 1, 10, 5), '*')
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()

	// restore tty
	t.Done()
}
