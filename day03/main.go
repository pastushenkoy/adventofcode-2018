package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"regexp"
	"strconv"
)

type Claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func main() {
	data := utils.ReadFileOfStrings("input.txt")

	claims := make([]Claim, 0)
	for _, row := range data {
		claims = append(claims, ClaimFromString(row))
	}

	fmt.Println("Number of intersections is ", getNumberOfOverlaps(claims))
	fmt.Println("The id of non-overlaping claim is ", getNonOverlappingClaim(claims).id)
}

func getNonOverlappingClaim(claims []Claim) Claim {
	var nonOverlappingClaim Claim
	for i := 0; i < len(claims); i++ {
		hasOverlaps := false
		for j := 0; j < len(claims); j++ {
			if i == j {
				continue
			}
			if Intersect(claims[i], claims[j]) {
				hasOverlaps = true
				break
			}
		}
		if !hasOverlaps {
			nonOverlappingClaim = claims[i]
		}
	}
	return nonOverlappingClaim
}

func getNumberOfOverlaps(claims []Claim) int {
	maxHorizontal := 0
	maxVertical := 0
	for _, claim := range claims {
		if claim.right() > maxHorizontal {
			maxHorizontal = claim.right()
		}
		if claim.bottom() > maxVertical {
			maxVertical = claim.bottom()
		}
	}
	filled := false
	overlaps := 0
	for i := 0; i <= maxHorizontal; i++ {
		for j := 0; j <= maxVertical; j++ {
			filled = false
			for _, claim := range claims {
				if claim.top <= i && claim.bottom() >= i && claim.left <= j && claim.right() >= j {
					if filled {
						overlaps++
						break
					} else {
						filled = true
					}
				}
			}
		}
	}
	return overlaps
}

func Intersect(claim1 Claim, claim2 Claim) bool {
	return claim1.left <= claim2.right() &&
		claim2.left <= claim1.right() &&
		claim1.top <= claim2.bottom() &&
		claim2.top <= claim1.bottom()
}

func ClaimFromString(row string) Claim {

	re := regexp.MustCompile("\\#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	parameters := re.FindStringSubmatch(row)
	id, _ := strconv.ParseInt(parameters[1], 10, 32)
	left, _ := strconv.ParseInt(parameters[2], 10, 32)
	top, _ := strconv.ParseInt(parameters[3], 10, 32)
	width, _ := strconv.ParseInt(parameters[4], 10, 32)
	height, _ := strconv.ParseInt(parameters[5], 10, 32)

	return Claim{
		id:     int(id),
		left:   int(left),
		top:    int(top),
		width:  int(width),
		height: int(height),
	}
}

func (c *Claim) right() int {
	return c.left + c.width - 1
}

func (c *Claim) bottom() int {
	return c.top + c.height - 1
}
