package gterm

import (
	"github.com/subfuzion/gterm/style"
)

// Terminal relays keyboard and mouse events and manages the screen.
type Terminal interface {
	// End restores the tty modes to the status before initializing the
	// terminal and closes open terminal device files.
	End()

	// Refresh updates the current terminal with the output of this screen.
	Refresh()

	// Set cell color and other style attributes for subsequent screen operations.
	SetStyle(fg, bg style.CellStyle)

	// Fill sets the screen contents with the specified rune.
	Fill(rune)

	// Clear erases the screen (by calling Fill with the rune for a space and
	// using the current background color)
	Clear()
}

