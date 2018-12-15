package main

import (
	"strings"
	"testing"
)

const testData = `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

func Test_loadData(t *testing.T){
	lines := strings.Split(testData, "\n")

	field := LoadField(lines)
	if len(field.Grid) != 48{
		t.Errorf("Expected 48 tracks but got %d", len(field.Grid))
	}
	if len(field.Carts) != 2{
		t.Errorf("Expected 2 carts but got %d", len(field.Carts))
	}
}

func Test_nextStep(t *testing.T){
	lines := strings.Split(testData, "\n")
	field := LoadField(lines)

	field.nextStep(0)

	expectedCart1Pos := Point{3, 0}
	actualCart1Pos := field.Carts[0].Position
	if !pointsEqual(actualCart1Pos, &expectedCart1Pos) {
		t.Errorf("Cart 1 is expected to be in %v but was in %v", expectedCart1Pos, actualCart1Pos)
	}
	expectedCart2Pos := Point{9, 4}
	actualCart2Pos := field.Carts[1].Position
	if !pointsEqual(actualCart2Pos, &expectedCart2Pos) {
		t.Errorf("Cart 2 is expected to be in %v but was in %v", expectedCart2Pos, actualCart2Pos)
	}
}

func Test_getFirstCollisionPoint(t *testing.T){
	lines := strings.Split(testData, "\n")
	field := LoadField(lines)

	expected := Point{7, 3}
	field.iterateToTheEnd()
	actual := field.Collisions[0].Position

	if !pointsEqual(&actual, &expected){
		t.Errorf("Expected first collision point to be %v but was %v", expected, actual)
	}
}
