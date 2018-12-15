package main

type Point struct {
	X, Y int
}

func (point *Point) move(direction Direction) {
	switch direction {
	case Up:
		point.Y--
	case Down:
		point.Y++
	case Left:
		point.X--
	case Right:
		point.X++
	}
}

func pointsEqual(p1 *Point, p2 *Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

