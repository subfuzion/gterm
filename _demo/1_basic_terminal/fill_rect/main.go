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

	// update screen to show cyan color
	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.Clear()

	// fill rect
	t.FillRect(geometry.MakeRectangle(2, 1, 10, 5), '*')
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()
	t.End()
}
