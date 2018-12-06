package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	data := loadData(utils.ReadFileOfStrings("input.txt"))
	fmt.Println("The largest non-infinite area is ", calculateLargestArea(data))
	fmt.Println("The area with max concentration is ", calculateAreaWithMaxConcentration(data, 10000))
}

func loadData(data []string) []*Point {
	re := regexp.MustCompile("(\\d+), (\\d+)")
	points := make([]*Point, 0)
	for _, line := range data {
		submatch := re.FindStringSubmatch(line)
		xStr, _ := strconv.ParseInt(submatch[1], 10, 32)
		yStr, _ := strconv.ParseInt(submatch[2], 10, 32)

		points = append(points, &Point{
			x: int(xStr),
			y: int(yStr),
		})
	}

	return points
}

func calculateLargestArea(points []*Point) int {
	maxX := getMaxX(points)
	maxY := getMaxY(points)

	areas := make([]int, len(points))
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			owner := getOwner(i, j, points)
			if owner != -1 {
				if i == 0 || j == 0 || i == maxX || j == maxY {
					areas[owner] = -1
				} else if areas[owner] != -1 {
					areas[owner]++
				}
			}
		}
	}

	return max(areas)
}

func calculateAreaWithMaxConcentration(points []*Point, maxDistance int) int {
	maxX := getMaxX(points)
	maxY := getMaxY(points)

	areaSize := 0
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if getTotalDistance(i, j, points) < maxDistance {
				areaSize++
			}
		}
	}

	return areaSize
}
func getTotalDistance(i, j int, points []*Point) int {
	totalDistance := 0
	for _, point := range points {
		totalDistance += int(abs(point.x-i) + abs(point.y-j))
	}
	return totalDistance
}

func getOwner(i, j int, points []*Point) int {
	distances := make([]int, len(points))
	for index, point := range points {
		distances[index] = int(abs(point.x-i) + abs(point.y-j))
	}

	minDistance := min(distances)

	closestPoint := -1
	for index, dist := range distances {
		if dist == minDistance {
			if closestPoint != -1 {
				return -1
			}
			closestPoint = index
		}
	}

	return closestPoint
}

func min(arr []int) int {
	minVal := math.MaxInt32
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func max(arr []int) int {
	maxVal := -1
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getMaxX(points []*Point) int {
	max := 0
	for i := 0; i < len(points); i++ {
		if points[i].x > max {
			max = points[i].x
		}
	}
	return max
}

func getMaxY(points []*Point) int {
	max := 0
	for i := 0; i < len(points); i++ {
		if points[i].y > max {
			max = points[i].y
		}
	}
	return max
}
