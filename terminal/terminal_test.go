package terminal_test

import (
	"os"
	"testing"
	"time"

	"github.com/subfuzion/gterm/style"
	"github.com/subfuzion/gterm/terminal"
)

var term terminal.Interface
var err error

func TestMain(m *testing.M) {
	term, err = terminal.New()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	term.Done()
	os.Exit(code)
}

func Test(t *testing.T) {
	t.Run("terminal", func(t *testing.T) {
		t.Run("screen should flash cyan", func(t *testing.T) {
			term.SetStyle(style.ColorBlack, style.ColorCyan)
			term.Clear()
			term.Refresh()
			// HACK: pause long enough to see something for visual check
			// TODO: remove sleep and actually query the screen to match expected
			time.Sleep(time.Millisecond * 100)
		})
	})
}
