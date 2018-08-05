package main

import (
	"fmt"

	"github.com/subfuzion/gterm/style"
	"github.com/subfuzion/gterm/terminal"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.Fill('*')
	t.Refresh()

	// wait for keypress to exit
	fmt.Scanln()
	t.End()
}
