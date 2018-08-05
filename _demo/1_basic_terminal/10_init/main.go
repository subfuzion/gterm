package main

import (
	"fmt"

	"github.com/subfuzion/gterm/terminal"
)

func main() {
	t, err := terminal.New()
	if err != nil {
		panic(err)
	}

	// wait for keypress to exit
	fmt.Scanln()

	// restore tty
	t.Done()
}
