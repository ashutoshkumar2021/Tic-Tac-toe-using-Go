package main

import (
	"fmt"
)

//initializing the variable board and currentPlayer
var board [3][3]string
var currentPlayer string

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

//Here we printing the board using nested loop
func printBoard() {
	fmt.Println("Tic-Tac-Toe")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s", board[i][j])
			if j < 2 {
				fmt.Print(" |")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("-----------")
		}
	}
	fmt.Println()
}

//As we have to switch player in the game so using this function, we can switch the player
func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

//so this function checks whether the board is full or not
func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

//this function basically checks whether the player one will win or player two
func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}

	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}

	return false
}

func main() {
	initializeBoard()
	//start with Player X
	currentPlayer = "X"
	var row, col int

	for {
		printBoard()
		fmt.Printf("Player %s's turn. Enter row (1-3) and column (1-3): ", currentPlayer)
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != " " {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		board[row-1][col-1] = currentPlayer

		if checkWin() {
			printBoard()
			fmt.Printf("Player %s wins! Congratulations!\n", currentPlayer)
			break
		} else if isBoardFull() {
			printBoard()
			fmt.Println("It's a draw! The game is over.")
			break
		}

		switchPlayer()
	}
}
