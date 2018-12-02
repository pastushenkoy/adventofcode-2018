package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"strconv"
)

func main() {
	freq := loadFrequencies()

	fmt.Println(fmt.Sprintf("The result frequency is %v", getResultFrequency(freq)))

	fmt.Println(fmt.Sprintf("The first duplicated frequency is %v", getFirstDuplicatedFrequency(freq)))
}

func getFirstDuplicatedFrequency(freq []int) int {
	hashSet := make(map[int]struct{})
	currentFreq := 0
	hashSet[currentFreq] = struct{}{}
	i := 0
	for {
		currentFreq += freq[i%len(freq)]
		if _, hasValue := hashSet[currentFreq]; hasValue {
			break;
		}

		hashSet[currentFreq] = struct{}{}
		i++
	}
	return currentFreq
}

func getResultFrequency(freq []int) int {
	sum := 0
	for _, num := range freq {
		sum += num
	}
	return sum
}

func loadFrequencies() []int {
	data := utils.ReadFileOfStrings("input.txt")
	var freq []int
	for _, st := range data {
		num, _ := strconv.Atoi(st)
		freq = append(freq, num)
	}
	return freq
}