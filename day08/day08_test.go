package main

import (
	"testing"
)

const testData = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func TestTask1(t *testing.T) {
	rootNode := getRootNode(testData)

	expected := 138
	actual := getMetadataSum(rootNode)
	if actual != expected {
		t.Error("Expected ", expected, ", got ", actual)
	}

}

func TestTask2(t *testing.T) {
	rootNode := getRootNode(testData)

	expected := 66
	actual := rootNode.Value
	if actual != expected {
		t.Error("Expected ", expected, ", got ", actual)
	}

}
