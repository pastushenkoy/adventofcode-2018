package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"strconv"
	"strings"
)

type Node struct {
	Nodes    []Node
	Metadata []int
	Value    int
}

func main() {
	data := utils.ReadStringFromFile("input.txt")
	rootNode := getRootNode(data)

	fmt.Printf("The metadata sum is %v\n", getMetadataSum(rootNode))
	fmt.Printf("The value of the root node is %v\n", rootNode.Value)
}

func getMetadataSum(rootNode Node) int {
	metadata := 0
	countMetadata(rootNode, &metadata)
	return metadata
}

func countMetadata(node Node, count *int) {
	for _, md := range node.Metadata {
		*count += md
	}

	for _, childNode := range node.Nodes {
		countMetadata(childNode, count)
	}
}

func readNode(ints []int, currentPosition *int) Node {
	nodesCount := readNext(ints, currentPosition, 1)[0]
	metadataCount := readNext(ints, currentPosition, 1)[0]

	currentNode := Node{
		Nodes: make([]Node, 0),
	}

	for i := 0; i < nodesCount; i++ {
		currentNode.Nodes = append(currentNode.Nodes, readNode(ints, currentPosition))
	}

	currentNode.Metadata = readNext(ints, currentPosition, metadataCount)

	for _, md := range currentNode.Metadata {
		if len(currentNode.Nodes) == 0 {
			currentNode.Value += md
		} else {
			if md <= len(currentNode.Nodes) {
				currentNode.Value += currentNode.Nodes[md-1].Value
			}
		}
	}

	return currentNode
}

func getRootNode(data string) Node {
	currentPosition := 0
	return readNode(getIntArrayFromString(data), &currentPosition)
}

func readNext(ints []int, currentPosition *int, count int) []int {
	value := ints[*currentPosition : *currentPosition+count]
	*currentPosition += count
	return value
}

func getIntArrayFromString(data string) []int {
	splitted := strings.Split(data, " ")

	result := make([]int, len(splitted))
	for i, rune := range splitted {
		num, _ := strconv.Atoi(rune)
		result[i] = num
	}

	return result
}
