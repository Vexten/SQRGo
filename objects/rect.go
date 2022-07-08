package objects

//line point for collision check
var lines [4]Point = [4]Point{*NewPoint(1,0),*NewPoint(0,1),*NewPoint(-1,0),*NewPoint(0,-1)}

type Rect struct {
	origin Point
	size   Point
	player byte
}

func (rect *Rect) Point(i int) *Point {
	/*
	(0)------(1)
	 |		  |
	(3)------(2)
	*/
	xShift := int(0)
	if ((i % 3) > 0) {
		xShift = 1
	}
	yShift := i / 2
	return &Point{rect.origin.x + rect.size.x * xShift, rect.origin.y + rect.size.y * yShift}
}

func (point *Point) inside(rect *Rect) bool {
	return point.GreaterThan(*rect.Point(0)) && point.LessThan(*rect.Point(2))
} 

//check if line is inside of a Rect
func (a *Rect) lineCollision(i int, b *Rect) bool {
	p1 := a.Point(i).Sum(lines[i])
	p2 := b.Point(i).Sum(lines[i])
	aLineColl := p1.inside(b)
	bLineColl := p2.inside(a)
	return (aLineColl || bLineColl)
}

func NewRect(x int, y int, width int, height int, player byte) Rect {
	return Rect{Point{x, y}, Point{width, height}, player}
}

func (rect *Rect) Start() *Point {
	return &rect.origin
}

func (rect *Rect) Size() *Point {
	return &rect.size
}

func (rect *Rect) Player() byte {
	return rect.player
}

func (rect *Rect) Area() int {
	return rect.size.x * rect.size.y
}

func (a *Rect) CollidesWith(b *Rect) bool {
	eqPoints := 0
	for i := 0; i < 4; i++ {
		bPointInsideA := b.Point(i).inside(a)
		aPointInsideB := a.Point(i).inside(b)
		//if any Point or any line is inside of a Rect
		if (a.Point(i).Equals(*b.Point(i))) {
			eqPoints++
		}
		if (bPointInsideA || aPointInsideB || a.lineCollision(i, b)) {
			return true
		}
	}
	return eqPoints > 2
}

/*
TODO:
rewrite to be human-readable
*/
//checks if rects are touching
func (a *Rect) Near(b* Rect) bool {
	if (b.origin.x + b.size.x == a.origin.x || a.origin.x + a.size.x == b.origin.x) {
		return (b.origin.y + b.size.y >= a.origin.y && b.origin.y <= a.origin.y + a.size.y);
	} else if (b.origin.y + b.size.y == a.origin.y || a.origin.y + a.size.y == b.origin.y) {
		return (b.origin.x + b.size.x >= a.origin.x && b.origin.x <= a.origin.x + a.size.x);
	}
	return false;
}