package main

import(
	"fmt"
	"strings"
	"strconv"
)

func EightQueensSolver() string {
	// Initialize an empty array to store the solutions
	var answer []string
	// Initialize the chessboard with all cells set to false (no queens placed)
	board := make([][]bool, 8)
	// Create an 8x8 chessboard
	for i:= range board{
		board[i] = make([]bool, 8)
	}
	// Solve the Eight Queens problem recursively
	solveQueens(&answer, board, 0)
	// Sort the solutions alphabetically using bubble sort
	bubbleSort(answer)
	// Concatenate the solutions into a single string separated by newline characters
	return strings.Join(answer, "\n")
}

// Implementation of bubble sort algorithm to sort an array of strings
func bubbleSort(s []string) {
    n := len(s)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if s[j] > s[j+1] {
                s[j], s[j+1] = s[j+1], s[j]
            }
        }
    }
}

func solveQueens(answer *[]string, board [][]bool, col int) {
	// Recursive function to solve the Eight Queens problem
	if col == 8 {
		// If all columns have been filled, add the current solution to the answer array
		*answer = append(*answer, boardToString(board))
		return

	}
	// Try placing a queen in each row of the current column
	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			// If it's safe to place a queen in this row of the current column, place the queen and move to the next column
			board[row][col] = true
			solveQueens(answer, board, col+1)
			// Backtrack by removing the queen from this cell
			board[row][col] = false
		}
	}
}

// Check if it's safe to place a queen in the given cell
func isSafe(board [][]bool, row, col int) bool {
	// Check if there's no queen in the same row
	for i := 0; i < col; i++ {
		if board[row][i] {
			return false
		}
	}
	// Check if there's no queen in the upper-left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	// Check if there's no queen in the lower-left diagonal
	for i, j := row, col; i < 8 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}
	// If all checks pass, it's safe to place a queen in this cell
	return true
}

func boardToString(board [][]bool) string {
	// Convert the chessboard representation to a string
	var sb strings.Builder
	for _, row := range board {
		queenIndex := -1
		for i, cell := range row {
			if cell {
				queenIndex = i
				break
			}
		}
		if queenIndex == -1 {
		// If no queen is found in a row, return an empty string (invalid solution)
			return ""

		}
		// Append the column index of the queen to the string
		sb.WriteString(strconv.Itoa(queenIndex + 1))
	}
	return sb.String()

}


func main(){
	fmt.Println(EightQueensSolver())
}
