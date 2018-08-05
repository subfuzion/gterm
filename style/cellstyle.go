package style

// CellStyle is used to specify the color and other
// attributes of an indivdiual screen cell.
type CellStyle uint16

// Color styles
const (
	ColorDefault CellStyle = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// Character attributes
// Note: these attributes will only take effect when applied
// to the foreground style of a cell.
const (
	Bold CellStyle = 1 << (iota + 9)
	Underline
	Reverse
)

