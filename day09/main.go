package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main(){
	fmt.Println("The highest score is ", getScore(400, 71864, false))
	fmt.Println("The highest score is ", getScore(400, 71864*100, true))
}


func getScore(playerCount int, lastBall int, verbose bool) int {
	playersScore := make([]int, playerCount)

	board := make([]int, 2)
	board[0] = 0
	board[1] = 1

	percent := int(lastBall / 100)

	percentage := 0
	currentMarble := 1
	for i := 2; i <= lastBall; i++{
		currentPlayer := (i - 1) % playerCount

		if i % 23 == 0 {
			removingMarblePosition := (currentMarble - 7 + len(board)) % (len(board))
			//fmt.Printf("On ball %d player %d gets %d + %d = %d points\n", i, currentPlayer, i, board[removingMarblePosition], i + board[removingMarblePosition])

			playersScore[currentPlayer] += i + board[removingMarblePosition]
			board = append(board[:removingMarblePosition], board[removingMarblePosition+1:]...)
			currentMarble = removingMarblePosition
		} else {
			newMarblePosition := (currentMarble + 2) % (len(board))
			board = append(board[:newMarblePosition], append([]int{i}, board[newMarblePosition:]...)...)
			currentMarble = newMarblePosition
		}

		if verbose && i % percent == 0{
			fmt.Println(time.Now().Format("15:04:05"), strconv.Itoa(percentage), "%")
			percentage++
		}

		//printBoard(board, currentMarble, currentPlayer)
	}

	sort.Slice(playersScore, func(i, j int) bool {
		return playersScore[i] > playersScore[j]
	})

	return playersScore[0]
}
func printBoard(board []int, currentMarble int, currentPlayer int) {
	fmt.Printf("[%1d] ", currentPlayer + 1)
	for i, value := range board{
		if i == currentMarble {
			fmt.Printf("(%2d)", value)
		}else {
			fmt.Printf(" %2d ", value)
		}
	}
	fmt.Println()
}
