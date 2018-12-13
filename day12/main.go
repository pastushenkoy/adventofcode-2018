package main

import (
	"bytes"
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"regexp"
)

type SurviveRule [5]bool

func main() {
	lines := utils.ReadFileOfStrings("input.txt")
	initState, rules := readData(lines)

	fmt.Printf("State hash after 20 generations is %d\n", getStateHashAfterXGenerations(initState, rules, 20))

	fmt.Printf("State hash after 50000000000 generations is %d\n", getStateHashAfterXGenerations(initState, rules, 50000000000))

}

func getStateHashAfterXGenerations(initState []bool, rules []SurviveRule, genCount int) int {
	currentState := initState
	zeroPosition := 0

	previousHash := 0
	previousHashDifference := 0
	timesDifferenceEquals := 0

	patternFound := false

	i := 1
	for ; i <= genCount; i++ {
		currentState = nextGeneration(&currentState, &zeroPosition, &rules)
		hash := getStateHash(currentState, zeroPosition)

		hashDifference := hash - previousHash
		previousHash = hash
		if hashDifference == previousHashDifference{
			timesDifferenceEquals++
		} else {
			previousHashDifference = hashDifference
			timesDifferenceEquals = 0
		}

		if timesDifferenceEquals > 2000{
			patternFound = true
			break
		}
	}

	hash := getStateHash(currentState, zeroPosition)
	if patternFound {
		fmt.Printf("step: %d, hash: %d, pattern: %d\n", i, hash, previousHashDifference)
		return hash + previousHashDifference*(genCount-i)
	} else {
		return hash
	}
}

func nextGeneration(initialState *[]bool, zeroPosition *int, rules *[]SurviveRule) []bool {
	if needToResize(initialState) {
		newStartIndex := 0
		*initialState, newStartIndex = resizeState(initialState)
		*zeroPosition = newStartIndex + *zeroPosition
	}

	nextState := make([]bool, len(*initialState))

	low, high := findEdges(*initialState)

	for i := low - 2; i <= high+2; i++ {
		for _, rule := range *rules {
			if matches(initialState, i, &rule) {
				nextState[i] = true
				break
			}
		}
	}

	return nextState
}

func needToResize(initialState *[]bool) bool {
	lastIndex := len(*initialState) - 1
	for i := 0; i < 4; i++{
		if (*initialState)[i] || (*initialState)[lastIndex - i]{
			return true
		}
	}
	return false
}

func convertToString(input []bool) string {
	var buffer bytes.Buffer

	for _, v := range input {
		if v {
			buffer.WriteRune('#')
		} else {
			buffer.WriteRune('.')
		}
	}

	return buffer.String()
}

func matches(initialState *[]bool, current int, pattern *SurviveRule) bool {
	for i := current - 2; i <= current+2; i++ {
		pV := pattern[i-current+2]

		//fmt.Printf("i: %d, l: %d\n", i, len(*initialState))



		isV := (*initialState)[i]
		if isV != pV {
			return false
		}
	}
	return true
}

func findEdges(initialState []bool) (low, high int) {
	for i := 0; i < len(initialState); i++ {
		if initialState[i] {
			low = i
			break
		}
	}

	for i := len(initialState) - 1; i >= 0; i-- {
		if initialState[i] {
			high = i
			break
		}
	}

	return
}

func resizeState(initialState *[]bool) ([]bool, int) {
	initLength := len(*initialState)
	newState := make([]bool, initLength*2)

	startIndex := int(initLength / 2)

	copy(newState[startIndex:], *initialState)
	return newState, startIndex
}

func readData(lines []string) ([]bool, []SurviveRule) {
	initRe := regexp.MustCompile("initial state: ([\\#\\.]+)")
	submatch := initRe.FindStringSubmatch(lines[0])

	initState := parsePattern(submatch[1])

	var rules []SurviveRule
	ruleRe := regexp.MustCompile("([\\.\\#]{5}) => ([\\.\\#])")
	for i := 2; i < len(lines); i++ {
		submatchRule := ruleRe.FindStringSubmatch(lines[i])
		if parseSymbol(rune(submatchRule[2][0])) {
			rules = append(rules, createRule(submatchRule[1], submatchRule[2]))
		}
	}

	return initState, rules
}

func createRule(pattern, result string) SurviveRule {
	var rule SurviveRule
	copy(rule[:], parsePattern(pattern))
	return rule
}

func parsePattern(pattern string) []bool {
	var result []bool
	for _, r := range pattern {
		result = append(result, parseSymbol(r))
	}
	return result
}

func parseSymbol(r rune) bool {
	switch r {
	case '.':
		return false
	case '#':
		return true
	default:
		panic("Unknown symbol")
	}
}

func getStateHash(input []bool, zeroPosition int) int {
	res := 0
	for i, e := range input {
		if e {
			res += i - zeroPosition
		}
	}
	return res
}

