package terminal

import (
	"fmt"

	"github.com/subfuzion/gterm"
	"github.com/subfuzion/gterm/driver"
	"github.com/subfuzion/gterm/geometry"
	"github.com/subfuzion/gterm/logging"
	"github.com/subfuzion/gterm/style"
)

// New initializes a new Terminal instance.
func New() (Interface, error) {
	t := &term{
		log:    gterm.Logger,
		driver: driver.NewTermbox(gterm.Logger),
	}

	t.log("Init")
	if err := t.driver.Init(); err != nil {
		return nil, &driverError{
			"term initialization failed",
			err.Error(),
		}
	}
	return t, nil
}

type driverError struct {
	err       string
	driverErr string
}

func (e *driverError) Error() string {
	return fmt.Sprintf("error: %s (driver error: '%s')", e.err, e.driverErr)
}

// term implements the Terminal interface and controls the underlying
// term device through the Driver interface.
type term struct {
	driver driver.Driver
	log    logging.LogPrinter
	fg     style.CellStyle
	bg     style.CellStyle
}

func (t term) Done() {
	t.log("Done")
	t.driver.Done()
}

func (t term) Refresh() {
	t.log("Refresh")
	t.driver.Refresh()
}

func (t term) Size() (width, height int) {
	t.log("Size")
	return t.driver.Size()
}

func (t *term) SetStyle(fg, bg style.CellStyle) {
	t.log("SetStyle(%+v, %+v)", fg, bg)
	t.fg = fg
	t.bg = bg
}

func (t term) GetStyle() (fg, bg style.CellStyle) {
	return t.fg, t.bg
}

func (t term) Clear() {
	t.log("Clear(%+v, %+v)", t.fg, t.bg)
	t.driver.Clear(t.fg, t.bg)
}

func (t term) Fill(ch rune) {
	t.log("Fill('%c')", ch)
	t.driver.Fill(t.fg, t.bg, ch)
}

func (t term) FillRect(rect geometry.Rectangle, ch rune) error {
	t.log("FillRect(%s, '%c')", rect, ch)

	if !t.InScreenBoundsRectangle(rect) {
		return ErrOutOfBoundsRequest(t.log, t.ScreenBounds(), rect)
	}

	t.driver.FillRect(rect, t.fg, t.bg, ch)
	return nil
}

func (t term) ClearRect(rect geometry.Rectangle) error {
	t.log("ClearRect(%s, '%c')", rect)
	return t.FillRect(rect, ' ')
}

func (t term) SetCell(x, y int, ch rune) error {
	t.log("SetCell(%d, %d, '%c')", x, y, ch)

	if !t.InScreenBoundsCoords(x, y) {
		return ErrOutOfBoundsRequest(t.log, t.ScreenBounds(), geometry.MakePoint(x, y))
	}

	t.driver.SetCell(x, y, ch, t.fg, t.bg)
	return nil
}

func (t term) DrawLineHorizontal(p geometry.Point, length int) error {

	return nil
}

func (t term) DrawLineVertical(p geometry.Point, length int) error {

	return nil
}

func (t term) ScreenBounds() geometry.Rectangle {
	w, h := t.Size()
	return geometry.MakeRectangle(0, 0, w, h)
}

func (t term) InScreenBoundsCoords(x, y int) bool {
	return t.ScreenBounds().ContainsCoords(x, y)
}

func (t term) InScreenBoundsPoint(p geometry.Point) bool {
	return t.ScreenBounds().ContainsPoint(p)
}

func (t term) InScreenBoundsLine(l geometry.Line) bool {
	return t.ScreenBounds().ContainsLine(l)
}

func (t term) InScreenBoundsRectangle(rect geometry.Rectangle) bool {
	return t.ScreenBounds().ContainsRectangle(rect)
}
