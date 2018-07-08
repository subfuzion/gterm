package gterm_test

import (
	"os"
	"testing"
	"github.com/subfuzion/gterm"
	"github.com/subfuzion/gterm/term"
)

var terminal gterm.Terminal
var err error

func TestMain(m *testing.M) {
	terminal, err = term.New()
	if err != nil {
		panic(err)
	}
	defer terminal.End()
	os.Exit(m.Run())
}

func TestTerminal(t *testing.T) {
	t.Run("gterm", func(t *testing.T) {
		t.Run("Clear", func(t *testing.T) {
			terminal.Clear()
		})

		t.Run("Refresh", func(t *testing.T) {
			terminal.Refresh()
		})
	})
}