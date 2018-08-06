package main

import (
	"fmt"

		"github.com/subfuzion/gterm/terminal"
	"github.com/subfuzion/gterm/style"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	// set cell
	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.SetCell(0, 0, '0')
	t.SetCell(1, 1, '1')
	t.SetCell(2, 2, '2')
	t.SetCell(3, 3, '3')
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()

	// restore tty
	t.Done()
}
