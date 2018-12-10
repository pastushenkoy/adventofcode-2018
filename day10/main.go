package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type Star struct {
	pX int
	pY int
	vX int
	vY int
}

func main() {
	lines := utils.ReadFileOfStrings("input.txt")

	var stars = getStarsFromStrings(lines)

	//distance := getDistance(stars)
	//var entropy int64 = math.MaxInt64

	const start = 10002

	jumpTo(stars, start)

	for i := start; i <= 10002 ; i++ {
		for _, star := range stars {
			star.pX += star.vX
			star.pY += star.vY
		}

		fmt.Println(i)
		printAllStars(stars)

		//time.Sleep(1000)

		//newEntropy := getEnthrophy(stars)
		//if newEntropy > entropy{
		//	fmt.Println(i)
		//	//listAllStars(stars)
		//	printAllStars(stars)
		//	//time.Sleep(500)
		//	break
		//} else {
		//	entropy = newEntropy
		//}


	}

}

func jumpTo(stars []*Star, start int) {
	for _, star := range stars{
		star.pX += star.vX * start
		star.pY += star.vY * start
	}
}

func getEnthrophy(stars []*Star) int64 {
	var result int64
	for i := 0; i < len(stars); i++{
		for j := 0; j< len(stars); j++{
			result += int64(getDistance(stars[i], stars[j]))
		}
	}
	return result
}

func listAllStars(stars []*Star) {
	for _, star := range stars {
		fmt.Printf("%d, %d\n", star.pX, star.pY)
	}
}

func printAllStars(stars []*Star) {
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := -math.MaxInt64
	maxY := -math.MaxInt64
	for _, star := range stars {
		if minX > star.pX {
			minX = star.pX
		}
		if minY > star.pY {
			minY = star.pY
		}
		if maxX < star.pX {
			maxX = star.pX
		}
		if maxY < star.pY {
			maxY = star.pY
		}
	}

	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := minY; i <= maxY; i++{
		for j := minX; j < maxX; j++{
			if isEmpty(stars, j, i){
				fmt.Print(".")
			} else {
				fmt.Print("G")
			}
		}
		fmt.Println()
	}
}

func isEmpty(stars []*Star, i int, j int) bool {
	for _, star := range stars{
		if star.pX == i && star.pY == j{
			return false
		}
	}
	return true
}

func getDistance(i *Star, j *Star) int {
	return abs(i.pX-j.pX) + abs(i.pY-j.pY)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func getStarsFromStrings(lines []string) []*Star {
	re := regexp.MustCompile("position=< *(-?\\d+), *(-?\\d+)> velocity=< *(-?\\d+), *(-?\\d+)>")
	var stars []*Star
	for _, line := range lines {
		submatches := re.FindStringSubmatch(line)
		pX, _ := strconv.ParseInt(submatches[1], 10, 32)
		pY, _ := strconv.ParseInt(submatches[2], 10, 32)
		vX, _ := strconv.ParseInt(submatches[3], 10, 32)
		vY, _ := strconv.ParseInt(submatches[4], 10, 32)

		stars = append(stars, &Star{
			pX: int(pX),
			pY: int(pY),
			vX: int(vX),
			vY: int(vY),
		})
	}

	return stars
}
