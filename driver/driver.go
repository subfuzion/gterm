package driver

import (
	"github.com/subfuzion/gterm/geometry"
	"github.com/subfuzion/gterm/style"
)

// Driver specifies the private interface required of system-specific
// terminal libraries. This layer of indirection allows the Terminal
// interface to evolve independently from the relatively small, stable
// Driver interface.
type Driver interface {
	Init() error
	Done()

	// only the following two methods update the screen contents
	// (by moving the back buffer to the terminal screen buffer
	Refresh()
	Clear(fg, bg style.CellStyle)

	Size() (width, height int)

	Fill(fg, bg style.CellStyle, ch rune)
	FillRect(rect geometry.Rectangle, fg, bg style.CellStyle, ch rune)

	SetCell(x, y int, ch rune, fg, bg style.CellStyle)

	// DrawLineHorizontal draws a horizontal line using the first and last
	// runes for the segment endponts, and the middle rune repeated in between.
	// The length can be negative.
	// Returns an error if the line is drawn out of bounds.
	//	DrawLineHorizontal(p geometry.Point, length int, first, middle, last rune)

	// DrawLineVertical draws a vertical line using the first and last
	// runes for the segment endponts, and the middle rune repeated in between.
	// The length can be negative.
	// Returns an error if the line is drawn out of bounds.
	//	DrawLineVertical(p geometry.Point, length int, first, middle, last rune)
}
