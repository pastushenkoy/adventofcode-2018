package main

import (
	"fmt"
	"sort"
)

type Field struct {
	Grid       Grid
	Carts      []*Cart
	Collisions []Collision
}

type Collision struct {
	Position Point
	Tick int
}

type Grid map[Point]Orientation

func (field *Field) nextStep(i int) {
	sort.Slice(field.Carts, func(i, j int) bool {
		if field.Carts[i].Position.Y == field.Carts[j].Position.Y {
			return field.Carts[i].Position.X < field.Carts[j].Position.X
		}

		return field.Carts[i].Position.Y < field.Carts[j].Position.Y
	})

	for _, cart := range field.Carts {
		if cart.Dead{
			continue
		}

		cart.move(&field.Grid)
		field.checkCollisions(cart, i)
	}
}

func (field *Field) checkCollisions(cart *Cart, i int) {
	for _, oppositeCart := range field.Carts{
		if oppositeCart != cart && oppositeCart.Position.X == cart.Position.X && oppositeCart.Position.Y == cart.Position.Y {
			if oppositeCart.Dead{
				continue
			}

			field.Collisions = append(field.Collisions, Collision{
				Position : Point{
					X: cart.Position.X,
					Y: cart.Position.Y,
				},
				Tick: i,
			})
			cart.Dead = true
			oppositeCart.Dead = true
		}
	}
}

func (field *Field) iterateToTheEnd() Point {
	//field.printCartStates(-1)
	for i := 0; ; i++ {
		field.nextStep(i)

		cart, oneCartLeft := field.onlyOneCartLeft();

		if oneCartLeft{
			if cart == nil {
				return Point{}
			}
			return *cart.Position
		}
		//field.printCartStates(i)
	}
}

func (field *Field) printCartStates(i int) {
	fmt.Printf("%3d ", i)
	for _, cart := range field.Carts {
		fmt.Printf("(%3d,%3d %v:%v) ", cart.Position.X+1, cart.Position.Y+1, cart.GetDirection(), cart.GetChoice())
	}
	fmt.Println()
}

func (field *Field) getCollision() (Point, bool) {
	for i := 0; i < len(field.Carts); i++ {
		for j := i + 1; j < len(field.Carts); j++{
			if pointsEqual(field.Carts[i].Position, field.Carts[j].Position){
				return *field.Carts[i].Position, true
			}
		}
	}
	return Point{}, false
}

func (field *Field) onlyOneCartLeft() (*Cart, bool) {
	var alive *Cart = nil
	for _, cart := range field.Carts{
		if !cart.Dead{
			if alive != nil{
				return nil, false
			}
			alive = cart
		}
	}

	return alive, true
}


func LoadField(lines []string) Field {
	field := Field{
		Grid:  make(map[Point]Orientation, 0),
		Carts: make([]*Cart, 0),
	}
	for j, line := range lines {
		for i, r := range line {
			switch r {
			case '|':
				field.Grid[Point{i, j}] = Vertical
			case '-':
				field.Grid[Point{i, j}] = Horizontal
			case '\\':
				field.Grid[Point{i, j}] = LeftTurn
			case '/':
				field.Grid[Point{i, j}] = RightTurn
			case '+':
				field.Grid[Point{i, j}] = Intersection
			case '^':
				field.Carts = append(field.Carts, &Cart{&Point{i, j}, Up, TurnRight, false})
				field.Grid[Point{i, j}] = Vertical
			case 'v':
				field.Carts = append(field.Carts, &Cart{&Point{i, j}, Down, TurnRight, false})
				field.Grid[Point{i, j}] = Vertical
			case '>':
				field.Carts = append(field.Carts, &Cart{&Point{i, j}, Right, TurnRight, false})
				field.Grid[Point{i, j}] = Horizontal
			case '<':
				field.Carts = append(field.Carts, &Cart{&Point{i, j}, Left, TurnRight, false})
				field.Grid[Point{i, j}] = Horizontal
			}
		}
	}
	return field
}
