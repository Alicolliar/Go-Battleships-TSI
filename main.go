package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

func bootstrapGame() error {
	inputReader = bufio.NewReader(os.Stdin)
	gameBoard = createGameBoard()
	err := userInitialisation(1)
	if err != nil {
		return err
	}
	aiInitialisation(2)
	return err
}

func main() {
	fmt.Printf("Hello!\nWelcome to Battleships!\n\n")
	err := bootstrapGame()
	if err != nil {
		panic(err)
	}
	turnCounter := 0
	for totalRemainingShips > 0 && turnCounter < 20 {
		fmt.Printf("Turn no.: %d\n\n", turnCounter)
		turnCounter++
		for _, player := range users {
			playerNum := player.userNum
			playerTurn(playerNum)

		}
	}
}
