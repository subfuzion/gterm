package terminal

import (
	"github.com/subfuzion/gterm/geometry"
	"github.com/subfuzion/gterm/style"
)

// Interface is the main terminal interface for relaying keyboard
// and mouse events and managing the screen.
type Interface interface {
	// End restores the tty modes to the status before initializing the
	// terminal and closes open terminal device files.
	End()

	// Refresh updates the current terminal with the output of this screen.
	Refresh()

	// Size returns the dimensions of the current terminal.
	Size() (width, height int)

	// Get current color and other style attributes.
	GetStyle() (fg, bg style.CellStyle)

	// Set color and other style attributes for subsequent rendering operations.
	SetStyle(fg, bg style.CellStyle)

	// Fill sets the terminal screen contents with the specified rune.
	Fill(rune)

	// Fill sets the terminal contents with the specified rune.
	// Returns an error if the rect exceeds the terminal bounds.
	FillRect(geometry.Rectangle, rune) error

	// Clear erases the terminal screen.
	Clear()

	// ClearRect erases the terminal screen contents specified by rect
	// (uses FillRect with the current background color and the run for
	// a space character).
	// Returns an error if the rect exceeds the terminal bounds.
	ClearRect(geometry.Rectangle) error
}
