package main

import (
	"fmt"
)

type userObject struct {
	userNum        int
	username       string
	remainingShips int
	score          int
}

var users [2]userObject

var totalRemainingShips int

func userInitialisation(userNum int) error {
	newUserObject := userObject{}
	newUserObject.remainingShips = 5
	newUserObject.score = 0
	fmt.Printf("Input your username\n> ")
	username, err := inputReader.ReadString('\n')
	if err != nil {
		return err
	}
	newUserObject.username = username[:len(username)-1]
	newUserObject.userNum = userNum
	fmt.Printf("Welcome to the game, %v. You are player %d\n\n", newUserObject.username, userNum)
	users[userNum-1] = newUserObject
	err = shipPlacement(userNum)
	if err != nil {
		return err
	}
	newUserObject.remainingShips = 5
	return nil
}

func playerTurn(userNum int) error {
	fmt.Print(users)
	displayBoardStates(userNum)
	fmt.Print(gameBoard)
	coordOk := false
	for !coordOk {
		result := false
		fmt.Printf("Please enter the coordinate you'd like to hit.\n>")
		coords, err := inputReader.ReadString('\n')
		if err != nil {
			return err
		}
		result, err = playerShoot(userNum, coords[:len(coords)-1])
		if err != nil {
			fmt.Println(err)
			coordOk = false
		} else if !result {
			fmt.Printf("%v was a miss!\n", coords[:len(coords)-1])
			registerShot(userNum, false)
			coordOk = true
		} else if result {
			fmt.Printf("%v was a hit! Ship sunk\n", coords[:len(coords)-1])
			registerShot(userNum, true)
			coordOk = true
		}
	}
	return nil
}

func shipPlacement(userNum int) error {
	for i := 1; i <= 5; i++ {
		coordsOk := false
		fmt.Println(userNum)
		for !coordsOk {
			fmt.Print("Please enter the position for a ship.\n> ")
			userCoords, err := inputReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if len(userCoords) <= 1 {
				coordsOk = false
				fmt.Printf("\nPlease try again, and actually enter co-ords this time.\n\n")
				continue
			}
			err = playerOccupation(userNum, userCoords[:len(userCoords)-1])
			if err != nil {
				fmt.Printf("\nPlease try again, %v.\n\n", err)
				coordsOk = false
				continue
			}
			coordsOk = true
		}
	}
	return nil
}

func registerShot(userNum int, hitResult bool) {
	if hitResult {
		totalRemainingShips--
		users[userNum-1].score++
		if userNum == 1 {
			users[1].remainingShips--
		} else if userNum == 2 {
			users[0].remainingShips--
		}
	}
}
