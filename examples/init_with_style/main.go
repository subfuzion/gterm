package main

import (
	"fmt"
	"github.com/subfuzion/gterm/style"
	"github.com/subfuzion/gterm/term"
)

func main() {
	t, err := term.New()
	if err != nil {
		panic(err)
	}
	defer t.End()

	t.SetStyle(style.ColorBlack, style.ColorCyan)
	t.Clear()
	t.Refresh()

	// pause until keypress
	fmt.Scanln()
}
