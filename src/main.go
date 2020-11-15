package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// defining a board
	defer timeTaken(time.Now())
	var sudokuBoard = [9][9]int{
		{0, 0, 0, 0, 0, 0, 6, 9, 0},
		{0, 0, 0, 1, 4, 0, 8, 2, 0},
		{5, 0, 0, 9, 0, 0, 0, 0, 4},
		{6, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 3, 0, 2, 0, 4, 0, 0},
		{0, 0, 0, 7, 0, 0, 0, 0, 9},
		{3, 0, 0, 0, 0, 1, 0, 0, 7},
		{0, 8, 4, 0, 3, 2, 0, 0, 0},
		{0, 9, 2, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println("The board looks as follows:")
	displayBoard(sudokuBoard)
	fmt.Println("Solving ......")
	solve(sudokuBoard)
}

func solve(board [9][9]int) bool {
	row, col := findEmptyCell(board)
	if row == -1{
		displayBoard(board)
		return true
	}
	for num := 1; num < 10; num++ {
		if isValid(board, num, row, col) == true{
			board[row][col] = num
			if solve(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

func findEmptyCell(board[9][9]int) (int, int){
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(board[9][9]int, num int, row int, col int)  bool{
	// Check row
	for i := range board {
		if (board[row][i] == num) && (col != i) {
			return false
		}
	}
	// Check column
	for j := range board {
		if (board[j][col] == num) && (row != j) {
			return false
		}
	}

	// Check box
	boxX := col/3
	boxY := row/3

	for i := boxY *3; i < boxY*3 + 3; i++ {
		for j := boxX *3; j < boxX*3 + 3; j++ {
			if (board[i][j] == num) && (i != row && j != col) {
				return false
			}
		}
	}
	return true
}

func displayBoard(board[9][9]int) {
	for i := range board {
		if i%3 == 0 {
			if i == 0 {
				fmt.Println(" ┎─────────┰─────────┰─────────┒")
			} else {
				fmt.Println(" ┠─────────╂─────────╂─────────┨")
			}
		}
		for j := range board {
			if j%3 == 0 {
				fmt.Print(" ┃  ")
			}
			if j == 8 {
				fmt.Println(board[i][j], " |")
			} else {
				fmt.Print(board[i][j], " ")
			}
		}
		}
	fmt.Println(" ┖─────────┸─────────┸─────────┚")
}

func timeTaken(start time.Time) {
	duration := time.Since(start)
	log.Printf("Took %s", duration)
}