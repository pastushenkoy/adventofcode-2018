package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileOfStrings(input string) []string {
	inputFile, err := filepath.Abs(input);
	check(err)
	dat, err := ioutil.ReadFile(inputFile)
	check(err)
	return strings.Split(string(dat), "\r\n")
}
