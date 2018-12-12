package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := getLargest3x3Square(6878)
	fmt.Printf("The largest square starts in %d, %d", x, y)

	x, y, size := getLargestSquareOfAnySize(6878)
	fmt.Printf("The largest square starts in %d, %d of size %d", x, y, size)
}

func getCellPower(x int, y int, serial int) int {
	return getHundreds(((x+10)*y+serial)*(x+10)) - 5
}

func getHundreds(number int) int {
	power := int(math.Log10(float64(number)))

	if power < 2 {
		return 0
	}

	rem := number
	for i := power; i > 2; i-- {
		rem = int(rem % int(math.Pow10(i)))
	}

	return int(rem / 100)
}

func getLargest3x3Square(serial int) (x, y int) {
	var field [300][300]int

	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			field[i][j] = getCellPower(i, j, serial)
		}
	}

	maxPower := 0
	largestX := 0
	largestY := 0
	for i := 0; i < 300-2; i++ {
		for j := 0; j < 300-2; j++ {
			power := 0
			for iSq := i; iSq < i+3; iSq++ {
				for jSq := j; jSq < j+3; jSq++ {
					power += field[iSq][jSq]
				}
			}

			if power > maxPower {
				maxPower = power
				largestX = i
				largestY = j
			}
		}
	}

	return largestX, largestY
}

func getLargestSquareOfAnySize(serial int) (x, y, lsize int) {
	const squareSize = 300

	var field [squareSize][squareSize]int

	for i := 0; i < squareSize; i++ {
		for j := 0; j < squareSize; j++ {
			field[i][j] = getCellPower(i, j, serial)
		}
	}

	maxPower := -math.MaxInt32
	largestX := 0
	largestY := 0
	largestSize := 0
	for i := 0; i < squareSize; i++ {
		fmt.Println(i)
		for j := 0; j < squareSize; j++ {
			maxSize := min(squareSize-i, squareSize-j)
			for size := 1; size <= maxSize; size++ {
				power := getPartialPower(i, j, size, &field)

				if power > maxPower {
					maxPower = power
					largestX = i
					largestY = j
					largestSize = size
				}
			}
		}
	}

	return largestX, largestY, largestSize
}

func getPartialPower(iCorner, jCorner, size int, field *[300][300]int) int {
	power := 0
	for i := iCorner; i < iCorner+size; i++ {
		for j := jCorner; j < jCorner+size; j++ {
			power += field[i][j]
		}
	}
	return power
}

func min(i int, i2 int) int {
	if i <= i2 {
		return i
	}
	return i2
}
