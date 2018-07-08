package driver

import (
	"github.com/subfuzion/gterm/style"
)

// Driver specifies the private interface required of system-specific
// terminal libraries. This layer of indirection allows the Terminal
// interface to evolve independently from the relatively small, stable
// Driver interface.
type Driver interface {
	Init() error
	End()
	Refresh()
	Clear(fg, bg style.CellStyle)
}

