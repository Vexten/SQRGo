package objects

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func (this *Point) X() int {
	return this.x
}

func (this *Point) Y() int {
	return this.y
}

func (a Point) Sum(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (a Point) Diff(b Point) Point {
	return Point{a.x - b.x, a.y - b.y}
}

func (a *Point) Add(b Point) {
	a.x += b.x
	a.y += b.y
}

func (a *Point) Sub(b Point) {
	a.x -= b.x
	a.y -= b.y
}

func (a Point) GreaterThan(b Point) bool {
	return (a.x > b.x && a.y > b.y)
}

func (a Point) LessThan(b Point) bool {
	return (a.x < b.x && a.y < b.y)
}