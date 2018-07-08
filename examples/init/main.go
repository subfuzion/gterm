package main

import (
	"fmt"
	"github.com/subfuzion/gterm/term"
)

func main() {
	t, err := term.New()
	if err != nil {
		panic(err)
	}
	defer t.End()

	// pause until keypress
	fmt.Scanln()
}
