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
	fmt.Print(users)
	return err
}

func main() {
	fmt.Printf("Hello!\nWelcome to Battleships!\n\n")
	err := bootstrapGame()
	if err != nil {
		panic(err)
	}
	turnCounter := 0
	for (users[0].remainingShips > 0 && users[1].remainingShips > 0) && turnCounter <= 20 {
		fmt.Printf("Turn no.: %d\n\n", turnCounter)
		turnCounter++
		for _, player := range users {
			playerNum := player.userNum
			if player.username == "Computer" {
				aiPlayerTurn(playerNum)
			} else {
				playerTurn(playerNum)
			}
		}
	}
	if users[0].score > users[1].score {
		fmt.Printf("Well done %v, you have won!", users[0].username)
	} else if users[0].score == users[1].score {
		fmt.Print("Something has gone horribly, horribly wrong, as it seems you both drew, which shouldn't really be possibly")
	} else {
		fmt.Print("Ah, better luck next time, the computer won that one.")
	}
}
