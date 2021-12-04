package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

func main() {
	// get the numbers to be drawn and the boards from the input file
	drawnNumbers, boards := parseInput()

	// simulate playing by going through numbers
	var currentDrawnNumbers []int
	for _, drawnNumber := range drawnNumbers {
		currentDrawnNumbers = append(currentDrawnNumbers, drawnNumber)
		// checking if any board is winning
		for _, board := range boards {
			if isWinningBoard(currentDrawnNumbers, board) {
				// solution sum of the unselected numbers multiplied by number that was just called
				solution := getUnselectedBoardSum(currentDrawnNumbers, board) * drawnNumber

				// print the solution!
				fmt.Println("Winner, winner, chicken dinner!")
				fmt.Println(solution)

				// end processing, we want the first winning board
				return
			}
		}
	}
}

func parseInput() ([]int, [][][]int) {
	var drawnNumbers []int
	var boards [][][]int
	var board [][]int

	lines := utils.ReadLines("input.txt")
	boardIndex := 0
	boardRowCount := 0
	for index, line := range lines {
		if index == 0 {
			drawnNumbersStrings := strings.Split(line, ",")
			drawnNumbers = utils.StringArrayToIntArray(drawnNumbersStrings)
			continue
		}
		if line == "" {
			continue
		}
		boardNumbersStrings := strings.Fields(line)
		boardNumbers := utils.StringArrayToIntArray(boardNumbersStrings)
		board = append(board, boardNumbers)
		boardRowCount++

		// we've made a full board - reset indexes, add board to boards, and clear board
		if boardRowCount == 5 {
			boardIndex++
			boardRowCount = 0
			boards = append(boards, board)
			board = nil
		}
	}

	return drawnNumbers, boards
}

func isWinningBoard(drawnNumbers []int, board [][]int) bool {
	// check for full row match
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		rowMatchCount := 0
		for colIndex := 0; colIndex < 5; colIndex++ {
			if utils.Contains(drawnNumbers, board[rowIndex][colIndex]) {
				rowMatchCount++
			}
		}
		if rowMatchCount == 5 {
			return true
		}
	}

	// check for full col match
	for colIndex := 0; colIndex < 5; colIndex++ {
		colMatchCount := 0
		for rowIndex := 0; rowIndex < 5; rowIndex++ {
			if utils.Contains(drawnNumbers, board[rowIndex][colIndex]) {
				colMatchCount++
			}
		}
		if colMatchCount == 5 {
			return true
		}
	}

	return false
}

func getUnselectedBoardSum(drawnNumbers []int, board [][]int) int {
	sum := 0
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		for colIndex := 0; colIndex < 5; colIndex++ {
			if !utils.Contains(drawnNumbers, board[rowIndex][colIndex]) {
				sum += board[rowIndex][colIndex]
			}
		}
	}
	return sum
}
