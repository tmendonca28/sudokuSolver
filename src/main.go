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
	fmt.Println("The board looks as follows:")
	displayBoard(sudokuBoard)
}

func solve(board [9][9]int) {

}

func isValid(board[9][9]int, num int, position int) {

}

func displayBoard(board[9][9]int) {
	for i, _ := range board {
		if i%3 == 0 {
			if i == 0 {
				fmt.Println(" ┎─────────┰─────────┰─────────┒")
			} else {
				fmt.Println(" ┠─────────╂─────────╂─────────┨")
			}
		}
		for j, _ := range board {
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

func findEmptyCell(board[9][9]int) (int, int){
	for i, _ := range board {
		for j, _ := range board {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}