package main

import (
	"strings"
	"testing"
)

const testData = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func TestTask1(t *testing.T){
	data := getConditions(strings.Split(testData, "\n"))
	expected := "CABDFE"
	actual := getStepSequence(data, false)
	if actual != expected {
		t.Error("Expected ", expected, ", got ", actual)
	}
}

func TestTask2(t *testing.T){
	data := getConditions(strings.Split(testData, "\n"))
	expected := 15
	actual := getWorkTime(data, 2, 0)
	if actual != expected {
		t.Error("Expected ", expected, ", got ", actual)
	}
}
