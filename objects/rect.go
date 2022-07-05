package objects

//shortcut for getting Rect lines
var lines [4][2]int = [4][2]int{{0,1},{1,2},{2,3},{3,0}}

type Rect struct {
	origin Point
	size   Point
	player byte
}

func (this *Rect) Point(i int) *Point {
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
	return &Point{this.origin.x + this.size.x * xShift, this.origin.y + this.size.y * yShift}
}

//check if line is inside of a Rect
func (a *Rect) lineCollision(i int, b *Rect) bool {
	//calculate difference between start and end points of both lines
	d1 := b.Point(lines[i][0]).Diff(*(a.Point(lines[i][0])))
	d2 := b.Point(lines[i][1]).Diff(*(a.Point(lines[i][1])))
	/*
	 ----(0)----
	 |         |
	(3)       (1)
	 |         |
	 ----(2)----
	*/
	switch i {
		case 0:
			return (d1.x == 0 && d2.x == 0 && d1.y > 0)
		case 1:
			return (d1.y == 0 && d2.y == 0 && d1.x < 0)
		case 2:
			return (d1.x == 0 && d2.x == 0 && d1.y < 0)
		case 3:
			return (d1.y == 0 && d2.y == 0 && d1.x > 0)
		default:
			return true
	}
}

func NewRect(x int, y int, width int, height int, player byte) Rect {
	return Rect{Point{x, y}, Point{width, height}, player}
}

func (this *Rect) Start() *Point {
	return &this.origin
}

func (this *Rect) Size() *Point {
	return &this.size
}

func (this *Rect) Player() byte {
	return this.player
}

func (this *Rect) Area() int {
	return this.size.x * this.size.y
}

func (a *Rect) CollidesWith(b *Rect) bool {
	for i := 0; i < 4; i++ {
		//if any Point or any line is inside of a Rect
		if ((b.Point(i).GreaterThan(*a.Point(0)) && b.Point(i).LessThan(*a.Point(2))) || a.lineCollision(i, b)) {
			return true
		}
	}
	return false
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