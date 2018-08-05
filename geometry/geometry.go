package geometry

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

// x and y correspond to column and row
type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point{x: %d, y: %d}", p.x, p.y)
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func (p *Point) Move(x, y int) {
	p.x = x
	p.y = y
}

func (p *Point) Offset(x, y int) {
	p.x += x
	p.y += y
}

func (p Point) IsLessThan(other Point) bool {
	return p.IsAbove(other) &&
		p.IsLeft(other)
}

func (p Point) IsLessThanOrEqual(other Point) bool {
	return p.IsAboveOrEqual(other) &&
		p.IsLeftOrEqual(other)
}

func (p Point) IsGreaterThan(other Point) bool {
	return p.IsBelow(other) &&
		p.IsRight(other)
}

func (p Point) IsGreaterThanOrEqual(other Point) bool {
	return p.IsBelowOrEqual(other) &&
		p.IsRightOrEqual(other)
}

func (p Point) IsLeft(other Point) bool {
	return p.x < other.x
}

func (p Point) IsLeftOrEqual(other Point) bool {
	return p.x <= other.x
}

func (p Point) IsRight(other Point) bool {
	return p.x > other.x
}

func (p Point) IsRightOrEqual(other Point) bool {
	return p.x >= other.x
}
func (p Point) IsAbove(other Point) bool {
	return p.y < other.y
}

func (p Point) IsAboveOrEqual(other Point) bool {
	return p.y <= other.y
}

func (p Point) IsBelow(other Point) bool {
	return p.y > other.y
}

func (p Point) IsBelowOrEqual(other Point) bool {
	return p.y >= other.y
}

func (p Point) InBounds(r Rectangle) bool {
	return r.ContainsPoint(p)
}

func (p Point) Distance(other Point) (rows, cols int) {
	return other.x - p.x, other.y - p.y
}

func LeftMostPoint(p1, p2 Point) Point {
	if p1.x < p2.x {
		return p1
	}
	return p2
}

func RightMostPoint(p1, p2 Point) Point {
	if p1.x > p2.x {
		return p1
	}
	return p2
}

func TopMostPoint(p1, p2 Point) Point {
	if p1.y < p2.y {
		return p1
	}
	return p2
}

func BottomMostPoint(p1, p2 Point) Point {
	if p1.y > p2.y {
		return p1
	}
	return p2
}

// MinPoint returns a new Point that represents the
// minimum x and y values of either point
func MinPoint(p1, p2 Point) Point {
	return Point{
		x: min(p1.x, p2.x),
		y: min(p1.y, p2.y),
	}
}

// MaxPoint returns a new Point that represents the
// maximum x and y values of either point
func MaxPoint(p1, p2 Point) Point {
	return Point{
		x: max(p1.x, p2.x),
		y: max(p1.y, p2.y),
	}
}

// Line is defined by two endpoints (technically a closed line segment)
type Line struct {
	a Point
	b Point
}

func (l Line) String() string {
	return fmt.Sprintf("Line{a: %s, b: %s}", l.a, l.b)
}

func (l Line) A() Point {
	return l.a
}

func (l Line) B() Point {
	return l.b
}

func (l Line) Width() (cols int) {
	return abs(l.b.x - l.a.x)
}

func (l Line) Height() (rows int) {
	return abs(l.b.y - l.a.y)
}

func (l Line) Extent() (cols, rows int) {
	return l.Width(), l.Height()
}

func (l Line) InBounds(r Rectangle) bool {
	return r.ContainsLine(l)
}

func (l Line) Vertical() bool {
	return l.a.x == l.b.x
}

func (l Line) Horizontal() bool {
	return l.a.y == l.b.y
}

// Rectangle defines the 0-based coordinates of the top-left (inclusive)
// and bottom-right (exclusive) corners of a rectangle.
type Rectangle struct {
	leftTop     Point
	rightBottom Point
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{leftTop: %s, rightBottom: %s [width: %d, height: %d]}", r.leftTop, r.rightBottom, r.Width(), r.Height())
}

func NormalRectangle(p1, p2 Point) Rectangle {
	return Rectangle{
		leftTop:     MinPoint(p1, p2),
		rightBottom: MaxPoint(p1, p2),
	}
}

func (r Rectangle) LeftTop() Point {
	return r.leftTop
}

func (r Rectangle) RightBottom() Point {
	return r.rightBottom
}

func (r Rectangle) Left() int {
	return r.leftTop.x
}

func (r Rectangle) Top() int {
	return r.leftTop.y
}

func (r Rectangle) Right() int {
	return r.rightBottom.x
}

func (r Rectangle) Bottom() int {
	return r.rightBottom.y
}

func (r Rectangle) Width() int {
	return abs(r.rightBottom.x - r.leftTop.x)
}

func (r Rectangle) Height() int {
	return abs(r.rightBottom.y - r.leftTop.y)
}

func (r Rectangle) Size() (width, height int) {
	return r.Width(), r.Height()
}

func (r *Rectangle) Resize(width, height int) {
	r.rightBottom = Point{r.leftTop.x + abs(width), r.leftTop.y + abs(height)}
}

func (r *Rectangle) Move(x, y int) {
	r.leftTop.x = x
	r.leftTop.y = y
}

func (r *Rectangle) Offset(x, y int) {
	r.leftTop.x += x
	r.leftTop.y += y
}

func (r Rectangle) InBounds(other Rectangle) bool {
	return other.ContainsRectangle(r)
}

func (r Rectangle) ContainsRectangle(other Rectangle) bool {
	return r.leftTop.IsLessThanOrEqual(other.leftTop) &&
		r.rightBottom.IsGreaterThanOrEqual(other.rightBottom)
}

func (r Rectangle) ContainsPoint(p Point) bool {
	return r.leftTop.IsLessThanOrEqual(p) &&
		r.rightBottom.IsGreaterThanOrEqual(p)
}

func (r Rectangle) ContainsLine(l Line) bool {
	return r.leftTop.IsLeftOrEqual(LeftMostPoint(l.a, l.b)) &&
		r.leftTop.IsAboveOrEqual(TopMostPoint(l.a, l.b)) &&
		r.rightBottom.IsRightOrEqual(RightMostPoint(l.a, l.b)) &&
		r.rightBottom.IsBelowOrEqual(BottomMostPoint(l.a, l.b))
}

func (r Rectangle) ClipLine(l Line) Line {
	return Line{
		a: MaxPoint(r.leftTop, MinPoint(l.a, l.b)),
		b: MinPoint(r.rightBottom, MaxPoint(l.a, l.b)),
	}
}

func MakeRectangle(x, y, width, height int) Rectangle {
	return NormalRectangle(Point{x, y}, Point{x + width, y + height})
}

func MakeRectangleOffset(r Rectangle, x, y int) Rectangle {
	return MakeRectangle(r.leftTop.x+x, r.leftTop.y+y, r.Width(), r.Height())
}
