package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
)

func main() {
	data := utils.ReadFileOfStrings("input.txt")

	fmt.Println(fmt.Sprintf("The checksum is %v", getChecksum(data)))

	fmt.Println(fmt.Sprintf("The common letters of right boxes ids are '%v'", getRightBoxesId(data)));
}

type distanceResult struct {
	id1 []rune
	id2 []rune
	distance int8
}

func getRightBoxesId(data []string) string {
	runesData := make([][]rune, len(data))
	for i, id := range data {
		runesData[i] = []rune(id)
	}

	c := make(chan distanceResult)
	for i := 0; i < len(runesData); i++ {
		for j := i + 1; j < len(runesData); j++ {
			go calculateDistance(runesData[i], runesData[j], c)
		}
	}

	for distance := range c{
		if distance.distance == 1{
			return string(getCommonLetters(distance.id1, distance.id2))
		}
	}

	panic("No right boxed found")
}

func getCommonLetters(id1 []rune, id2 []rune) []rune {
	var commonLetters = make([]rune, len(id1) - 1)
	j := 0
	for i := 0; i < len(id1); i++ {
		if id1[i] == id2[i] {
			commonLetters[j] = id1[i]
			j++
		}
	}
	return commonLetters
}

func calculateDistance(id1 []rune, id2 []rune, c chan distanceResult) {
	var distance int8 = 0
	for i := 0; i < len(id1); i++ {
		if id1[i] != id2[i] {
			distance++
		}
	}

	c <- distanceResult{
		id1:id1,
		id2:id2,
		distance:distance,
	}
}

func getChecksum(data []string) int {
	doubles := 0
	triples := 0
	for _, id := range data {
		runes := []rune(id)

		runeCounts := make(map[rune]byte)

		for _, letter := range runes {
			if _, hasValue := runeCounts[letter]; hasValue {
				runeCounts[letter]++
			} else {
				runeCounts[letter] = 1
			}
		}

		hasDouble := false
		hasTriple := false
		for _, runeCount := range runeCounts {
			if runeCount == 2 {
				hasDouble = true
			} else if runeCount == 3 {
				hasTriple = true
			}
		}

		if hasDouble {
			doubles++
		}
		if hasTriple {
			triples++
		}
	}
	checksum := doubles * triples
	return checksum
}
