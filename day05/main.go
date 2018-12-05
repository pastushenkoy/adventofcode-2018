package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"strings"
)

func main() {
	data := utils.ReadStringFromFile("input.txt")

	fmt.Println("The length of the rest string is ", getSurvivedPolymersCount(data))

	fmt.Println("The length of the rest string after excluging most recent char is ", getShortestPolymerLength(data))
}

func getSurvivedPolymersCount(data string) int {
	runes := []rune(data)

	for allSurvived := false; !allSurvived; {
		allSurvived = true
		for i := 0; i < len(runes)-1; {
			currentRune := runes[i]
			if currentRune == '_' {
				i++
				continue
			}
			nextIndex, exists := findNextRune(runes, i)
			if !exists {
				break
			}
			nextRune := runes[nextIndex]

			if destroyed(currentRune, nextRune) {
				runes[i] = '_'
				runes[nextIndex] = '_'
				i = nextIndex + 1
				allSurvived = false
			} else {
				i++
			}

		}
	}

	runeCount := 0
	for _, r := range runes {
		if r != '_' {
			runeCount++
		}
	}

	return runeCount
}

func findNextRune(runes []rune, currentRuneIndex int) (nextIndex int, exists bool) {
	for i := currentRuneIndex + 1; i < len(runes); i++ {
		if runes[i] != '_' {
			return i, true
		}
	}
	return 0, false
}

func destroyed(r1 rune, r2 rune) bool {
	return r1 != '_' && r1 != r2 && strings.ToLower(string(r1)) == strings.ToLower(string(r2))
}

func getShortestPolymerLength(data string) int {
	shortestPolymerLength := len(data)
	for i := 97; i <= 122; i++ {
		lower := string(i)
		upper := strings.ToUpper(lower)

		tempData := strings.Replace(data, lower, "_", -1)
		tempData = strings.Replace(tempData, strings.ToUpper(upper), "_", -1)
		length := getSurvivedPolymersCount(tempData)
		if length < shortestPolymerLength{
			shortestPolymerLength = length
		}

	}

	return shortestPolymerLength
}
