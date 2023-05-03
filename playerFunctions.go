package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

type userObject struct {
	userNum        int
	username       string
	remainingShips int
}

var users []userObject

var totalRemainingShips int

func userInitialisation(userNum int) error {
	userNumStr := string(rune(userNum))
	inputReader = bufio.NewReader(os.Stdin)
	newUserObject := userObject{}
	fmt.Printf("Input your username\n>")
	username, err := inputReader.ReadString('\n')
	if err != nil {
		return err
	}
	newUserObject.username = username
	fmt.Printf("Welcome to the game, %v. You are player %v", username, userNumStr)
	users = append(users, newUserObject)
	return nil
}

func playerTurn(userNum int) error {
	// displayBoardStates(userNum)
	coordOk := false
	for !coordOk {
		result := false
		fmt.Printf("Please enter the co-ordinate you'd like to hit.\n>")
		coords, err := inputReader.ReadString('\n')
		if err != nil {
			return err
		}
		result, err = playerShoot(userNum, coords)
		if err != nil {
			fmt.Println(err.Error())
			coordOk = false
		} else if !result {
			fmt.Printf("%v was a miss!\n", coords)
			// registerShot(userNum,coords, false)
			coordOk = true
		} else if result {
			fmt.Printf("%v was a hit! Ship sunk\n", coords)
			// registerShot(userNum, coords, true)
			coordOk = true
		}
	}
	return nil
}
