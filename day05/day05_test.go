package main

import "testing"

func TestTask01(t *testing.T) {
	if getSurvivedPolymersCount("dabAcCaCBAcCcaDA") != 10 {
		t.Fail()
	}
}

func TestTask02(t *testing.T) {
	if getShortestPolymerLength("dabAcCaCBAcCcaDA") != 4 {
		t.Fail()
	}
}
