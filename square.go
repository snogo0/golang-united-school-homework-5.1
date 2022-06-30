package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq Square) End() Point {
	return Point{sq.start.x + int(sq.a), sq.start.y + int(sq.a)}
}

func (sq Square) Area() uint {
	return sq.a * sq.a
}

func (sq Square) Perimeter() uint {
	return sq.a * 4
}
