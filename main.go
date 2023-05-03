package main

import (
	"fmt"
)

func bootstrapGame() {
	gameBoard = createGameBoard()
}

func main() {
	fmt.Printf("Hello!\n Welcome to the game!\n\n")
	bootstrapGame()
	fmt.Print(gameBoard)
}
