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
	End()

	// only the following two methods update the screen contents
	// (by moving the back buffer to the terminal screen buffer
	Refresh()
	Clear(fg, bg style.CellStyle)

	Size() (width, height int)

	Fill(fg, bg style.CellStyle, ch rune)
	FillRect(rect geometry.Rectangle, fg, bg style.CellStyle, ch rune)
}
