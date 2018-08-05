package main

import (
	"fmt"

	"github.com/subfuzion/gterm/geometry"
	"github.com/subfuzion/gterm/style"
	"github.com/subfuzion/gterm/terminal"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	// fill rect
	t.SetStyle(style.ColorDefault, style.ColorCyan)
	t.ClearRect(geometry.MakeRectangle(2, 1, 10, 5))
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()

	// restore tty
	t.Done()
}
