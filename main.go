package main

import (
	"fmt"
)

/*
Global constants and variables definations:

DeadCell        : Represents value '0' in the board
LiveCell        : Represents value '1' in the board
ConvertedToDead : For inplace change of values, any cell which is converted to dead from live will be represented with value '2'
ConvertedToLive : For inplace change of values, any cell which is converted to live from dead will be represented with value '3'
m               : Number of rows in the board
n               : Number of cols in the board
directions      : Slice of eight directions
*/
const (
	DeadCell        = iota // 0
	LiveCell               // 1
	ConvertedToDead        // 2
	ConvertedToLive        // 3
)

var (
	directions = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	board = [][]int{}
	m = 25
	n = 25
)

/*
Description: This function initilises the board, takes input (dummy 25x25 board input already given), and calculates the next state, the number times user has asked for
*/
func gameOfLife(countOfNextState int) {

	initialiseBoard()

	inputBoard()

	fmt.Println("\nInitial life board:")
	displayBoard()

	for i := 0; i < countOfNextState; i += 1 {

		getNextState()

		fmt.Printf("\nState %v:\n", i+1)
		displayBoard()
	}

}

/*
Description: This function initilises the board with 'm' rows and 'n' columns
*/
func initialiseBoard() {

	board = make([][]int, m)

	for row := 0; row < m; row += 1 {
		board[row] = make([]int, n)
	}
}

/*
Description: This function can be modified for taking input from the user, but for now it has initialised the board with 25x25 board as given in the assignment
*/
func inputBoard() {

	// default 25 X 25 board using Glider pattern
	board = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

/*
Description: This function calculates the next state in-place as per the given rules:

At each step in time, the following transitions occur:
1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
*/
func getNextState() {

	for row := 0; row < m; row += 1 {

		for col := 0; col < n; col += 1 {

			liveNeighbourCount, _ := getLiveDeadNeighbourCount(row, col)

			// when current cell is live
			if board[row][col] == LiveCell {

				if liveNeighbourCount < 2 || liveNeighbourCount > 3 {
					board[row][col] = ConvertedToDead
				}

				// when current cell is dead
			} else if board[row][col] == DeadCell {

				if liveNeighbourCount == 3 {
					board[row][col] = ConvertedToLive
				}
			}
		}
	}

	revertBoard()
}

/*
Description: This function caluculates the live and dead neighbour count for the given cell
*/
func getLiveDeadNeighbourCount(row, col int) (int, int) {

	liveNeighbourCount, deadNeighbourCount := 0, 0

	for _, direction := range directions {

		validNeighbourExists, neighbourXCoordinate, neighbourYCoordinate := getNextValidNeighbour(row, col, direction)

		if !validNeighbourExists {
			continue
		}

		if board[neighbourXCoordinate][neighbourYCoordinate] == LiveCell || board[neighbourXCoordinate][neighbourYCoordinate] == ConvertedToDead {
			liveNeighbourCount += 1
		} else {
			deadNeighbourCount += 1
		}
	}

	return liveNeighbourCount, deadNeighbourCount
}

/*
Description: This function validates the neighbour is valid cell or not i.e neighbour index should fall within the boundaries of the board
*/
func getNextValidNeighbour(currPosX, currPosY int, direction []int) (bool, int, int) {

	nextPosX := currPosX + direction[0]
	nextPosY := currPosY + direction[1]

	if nextPosX < 0 || nextPosX >= m || nextPosY < 0 || nextPosY >= n {

		return false, -1, -1
	}

	return true, nextPosX, nextPosY
}

/*
Description: This function reverts the board from values which include intermediate state to values with only final state values i.e. Dead '0' or Live '1'
*/
func revertBoard() {

	for row := 0; row < m; row += 1 {

		for col := 0; col < n; col += 1 {

			if board[row][col] == ConvertedToDead {
				board[row][col] = 0
			} else if board[row][col] == ConvertedToLive {
				board[row][col] = 1
			}
		}
	}
}

/*
Description: This function prints the current state of board
*/
func displayBoard() {

	for _, row := range board {

		fmt.Printf("%v\n", row)
	}
}

func main() {

	var nextStateCount int
	fmt.Println("\n\nPlease enter the number of times you want to generate the next generation: ")
	fmt.Scanln(&nextStateCount)

	gameOfLife(nextStateCount)
}
