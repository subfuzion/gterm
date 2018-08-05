package terminal

import (
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

	// Set cell color and other style attributes for subsequent screen operations.
	SetStyle(fg, bg style.CellStyle)

	// Fill sets the screen contents with the specified rune.
	Fill(rune)

	// Clear erases the screen (by calling Fill with the rune for a space and
	// using the current background color)
	Clear()
}

