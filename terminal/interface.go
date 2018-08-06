package terminal

import (
	"github.com/subfuzion/gterm/geometry"
	"github.com/subfuzion/gterm/style"
)

// Interface is the main terminal interface for relaying keyboard
// and mouse events and managing the screen.
type Interface interface {
	// Done restores the tty modes to the status before initializing the
	// terminal and closes open terminal device files.
	Done()

	// Refresh updates the current terminal with the output of this screen.
	Refresh()

	// Size returns the dimensions of the current terminal.
	Size() (width, height int)

	// Get current color and other style attributes.
	GetStyle() (fg, bg style.CellStyle)

	// Set color and other style attributes for subsequent rendering operations.
	SetStyle(fg, bg style.CellStyle)

	// Clear erases the terminal screen.
	Clear()

	// Fill sets the terminal screen contents with the specified rune.
	Fill(rune)

	// Fill sets the terminal contents with the specified rune.
	// Returns an error if the rect exceeds the terminal bounds.
	FillRect(geometry.Rectangle, rune) error

	// ClearRect erases the terminal screen contents specified by rect
	// (uses FillRect with the current background color and the run for
	// a space character).
	// Returns an error if the rect exceeds the terminal bounds.
	ClearRect(geometry.Rectangle) error

	// SetCell writes the rune to the screen location using the current
	// style.
	// Returns an error if the location specified is out of bounds.
	SetCell(x, y int, ch rune) error

	// DrawLineHorizontal draws a horizontal line using the first and last
	// runes for the segment endponts, and the middle rune repeated in between.
	// The length can be negative.
	// Returns an error if the line is drawn out of bounds.
	DrawLineHorizontal(p geometry.Point, length int, first, middle, last rune) error

	// DrawLineVertical draws a vertical line using the first and last
	// runes for the segment endponts, and the middle rune repeated in between.
	// The length can be negative.
	// Returns an error if the line is drawn out of bounds.
	DrawLineVertical(p geometry.Point, length int, first, middle, last rune) error
}
