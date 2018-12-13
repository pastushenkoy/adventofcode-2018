package main

import (
	"strings"
	"testing"
)

const testRules string = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
..... => .`

func Test_getStateHashAfterXGenerations(t *testing.T){
	lines := strings.Split(testRules, "\n")
	initState, rules := readData(lines)

	expected := 325
	actual := getStateHashAfterXGenerations(initState, rules, 20)

	if expected != actual{
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func Test_readData(t *testing.T){
	lines := strings.Split(testRules, "\n")
	initState, rules := readData(lines)
	if len(initState) != 25{
		t.Errorf("Length of initial state should be 25 but was %d", len(initState))
	}
	if len(rules) != 14 {
		t.Errorf("The count of rules should be 14 but was %d", len(rules))
	}
}


func Test_nextGeneration(t *testing.T) {
	lines := strings.Split(testRules, "\n")
	_, rules := readData(lines)

	tables := []struct {
		current string
		next string
	}{
		{"#..#.#..##......###...###", "#...#....#.....#..#..#..#"},
		{"#...#....#.....#..#..#..#", "##..##...##....#..#..#..##"},
	}

	for i := 0; i < len(tables); i++ {
		zeroPosition := 0

		row := tables[i]
		expected := parsePattern(row.next)
		pattern := parsePattern(row.current)
		surviveRules := rules
		actual := nextGeneration(&pattern, &zeroPosition, &surviveRules)

		if getStateHash(expected, 0) != getStateHash(actual, zeroPosition){
			t.Errorf(`
Expected '%v'
but got  '%v''`, convertToString(expected), convertToString(actual))
		}

	}


}


