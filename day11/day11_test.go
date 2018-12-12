package main

import (
	"testing"
)

func Test_getHundreds(t *testing.T) {
	tables := []struct {
		number int
		result int
	}{
		{23, 0},
		{345, 3},
		{3456, 4},
		{65452234, 2},
	}

	for i := 0; i < len(tables); i++ {
		row := tables[i]
		actual := getHundreds(row.number)
		if actual != row.result {
			t.Errorf("Get hundreds for %d returned %d", row.number, row.result)
		}
	}
}

func Test_getCellPower(t *testing.T) {
	tables := []struct {
		x      int
		y      int
		serial int
		result int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for i := 0; i < len(tables); i++ {
		row := tables[i]
		actual := getCellPower(row.x, row.y, row.serial)
		if actual != row.result {
			t.Errorf("Result for cell %d,%d and serial '%d' expected %d, but got %d", row.x, row.y, row.serial, row.result, actual)
		}
	}
}

func Test_getLargest3x3Square(t *testing.T) {
	tables := []struct {
		serial  int
		resultX int
		resultY int
	}{
		{18, 33, 45},
		{42, 21, 61},
	}

	for i := 0; i < len(tables); i++ {
		row := tables[i]
		actualX, actualY := getLargest3x3Square(row.serial)
		if actualX != row.resultX || actualY != row.resultY {
			t.Errorf("Expected result for serial %d was %d, %d but got %d, %d", row.serial, row.resultX, row.resultY, actualX, actualY)
		}
	}
}

func Test_getLargestSquareOfAnySize(t *testing.T) {
	tables := []struct {
		serial     int
		resultX    int
		resultY    int
		resultSize int
	}{
		{18, 90,269,16},
		{42, 232,251,12},
	}

	for i := 0; i < len(tables); i++ {
		row := tables[i]
		actualX, actualY, actualSize := getLargestSquareOfAnySize(row.serial)
		if actualX != row.resultX || actualY != row.resultY {
			t.Errorf("Expected result for serial %d was %d, %d of size %d but got %d, %d of size %d",
				row.serial, row.resultX, row.resultY, row.resultSize, actualX, actualY, actualSize)
		}
	}
}
