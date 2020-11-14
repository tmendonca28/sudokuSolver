package main
import "fmt"

func main() {
	// defining a board
	var sudokuBoard = [9][9]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	fmt.Println("I will solve all sudoku problems")
}

func solve(board [9][9]int) {

}

func isValid(board[9][9]int, num int, position int) {

}

func displayBoard(board[9][9]int) {

}

func findEmptyCell(board[9][9]int) {

}