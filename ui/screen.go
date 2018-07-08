package ui

import "github.com/subfuzion/gterm/style"

// Cell contains the content and style attributes to be applied to
// a specific location (row, column) on the screen.
type Cell struct {
	Rune rune
	Fg style.CellStyle
	Bg style.CellStyle
}

