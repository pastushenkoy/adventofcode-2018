package main

import "testing"

func TestTask1(t *testing.T) {
	tables := []struct {
		playerCount int
		lastBall    int
		highScore   int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for i := 0; i < len(tables); i++ {
		row := tables[i]
		actual := getScore(tables[i].playerCount, row.lastBall)
		if actual != row.highScore {
			t.Errorf("Players count: %d, last ball: %d, Expected %d, got %d", row.playerCount, row.lastBall, row.highScore, actual)
		}
	}
	/*Here are a few more examples:

	10 players; last marble is worth 1618 points: high score is 8317
	13 players; last marble is worth 7999 points: high score is 146373
	17 players; last marble is worth 1104 points: high score is 2764
	21 players; last marble is worth 6111 points: high score is 54718
	30 players; last marble is worth 5807 points: high score is 37305*/
}
