package main

type Orientation int

const (
	Vertical     Orientation = 1
	Horizontal   Orientation = 2
	LeftTurn     Orientation = 3 // при движении по вертикали
	RightTurn    Orientation = 4 // при движении по вертикали
	Intersection Orientation = 7
)

type Direction int

const (
	Up    Direction = 0
	Left  Direction = 1
	Down  Direction = 2
	Right Direction = 3
)

type Choice int

const (
	TurnLeft   Choice = 0
	GoStraight Choice = 1
	TurnRight  Choice = 2
)

type Cart struct {
	Position  *Point
	Direction Direction
	Choice    Choice
	Dead      bool
}

func (cart *Cart) move(grid *Grid) {
	cart.Position.move(cart.Direction)
	if (*grid)[*cart.Position] == LeftTurn {
		if cart.Direction == Up || cart.Direction == Down {
			cart.rotateLeft()
		} else {
			cart.rotateRight()
		}
	} else if (*grid)[*cart.Position] == RightTurn {
		if cart.Direction == Up || cart.Direction == Down {
			cart.rotateRight()
		} else {
			cart.rotateLeft()
		}
	} else if (*grid)[*cart.Position] == Intersection {
		cart.makeChoice()
	}
}
func (cart *Cart) rotateLeft() {
	cart.Direction = Direction((cart.Direction + 1)%4)
}
func (cart *Cart) rotateRight() {
	cart.Direction = Direction((cart.Direction - 1 + 4)%4)
}
func (cart *Cart) makeChoice() {
	cart.Choice = Choice((cart.Choice + 1)%3)
	if cart.Choice == TurnLeft{
		cart.rotateLeft()
	} else if cart.Choice == TurnRight{
		cart.rotateRight()
	}
}
func (cart *Cart) GetDirection() string {
	switch cart.Direction {
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	}
	return "X"
}

func (cart *Cart) GetChoice() string {
	switch cart.Choice {
	case TurnLeft:
		return "\\"
	case TurnRight:
		return "/"
	case GoStraight:
		return "|"
	}
	return "X"
}

