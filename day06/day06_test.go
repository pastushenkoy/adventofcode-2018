package main

import (
	"strings"
	"testing"
)

const testData = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

func TestTask1(t *testing.T) {
	points := loadData(strings.Split(testData, "\n"))
	if calculateLargestArea(points) != 17 {
		t.Fail()
	}
}

func TestTask2(t *testing.T) {
	points := loadData(strings.Split(testData, "\n"))
	if calculateAreaWithMaxConcentration(points, 32) != 16 {
		t.Fail()
	}
}
