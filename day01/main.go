package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(input string) []string {
	inputFile, err := filepath.Abs(input);
	check(err)
	dat, err := ioutil.ReadFile(inputFile)
	check(err)
	return strings.Split(string(dat), "\r\n")
}

func main() {
	data :=readFile("input.txt")
	var freq []int
	sum := 0
	for _, st := range data{
		num, _ := strconv.Atoi(st)
		freq = append(freq, num)
		sum += num
	}
	fmt.Println(sum)

	hashset := make(map[int]struct{})

	currentFreq := 0
	hashset[currentFreq] = struct{}{}
	i:=0
	for {
		currentFreq += freq[i%len(freq)]
		if _, hasValue := hashset[currentFreq]; hasValue{
			break;
		}

		hashset[currentFreq] = struct{}{}
		i++
	}

	fmt.Println(currentFreq)
	fmt.Print(i)
}