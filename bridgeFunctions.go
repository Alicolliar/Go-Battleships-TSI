package main

import (
	"errors"
	"fmt"
)

func stringInSlice(searchFor string, searchMedia []string) bool {
	for _, indexPosition := range searchMedia {
		if indexPosition == searchFor {
			return true
		}
	}
	return false
}

func displayBoardStates(userNum int) {
	xCounter := 0
	var shipContaining, hitContaining []string
	for _, coordPoint := range gameBoard {
		if xCounter%xLength == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%v ", coordPoint.coordString)
		if userNum == 1 {
			if coordPoint.player1Ship {
				shipContaining = append(shipContaining, coordPoint.coordString)
			} else if coordPoint.player1HitShot {
				hitContaining = append(hitContaining, coordPoint.coordString)
			}
		} else if userNum == 2 {
			if coordPoint.player2Ship {
				shipContaining = append(shipContaining, coordPoint.coordString)
			} else if coordPoint.player2HitShot {
				hitContaining = append(hitContaining, coordPoint.coordString)
			}
		}
		xCounter++
	}
	fmt.Print("\n\n")
	fmt.Printf("Coord\tShip?\tHit?\n")
	for _, coordPoint := range gameBoard {
		fmt.Printf("%v\t", coordPoint.coordString)
		if stringInSlice(coordPoint.coordString, shipContaining) {
			fmt.Print("x\t")
		} else {
			fmt.Print(" \t")
		}
		if stringInSlice(coordPoint.coordString, hitContaining) {
			fmt.Print("x")
		}
		fmt.Print("\n")
	}
}

func playerShoot(playerNum int, positionString string) (bool, error) {
	coordI, err := decodeCoordString(positionString)
	if err != nil {
		return false, err
	}
	occupied, err := checkOccupation(playerNum, coordI)
	if err != nil {
		return false, err
	}
	if occupied {
		if playerNum == 1 {
			gameBoard[coordI].player1HitShot = true
			gameBoard[coordI].player2Ship = false
			return true, nil
		} else if playerNum == 2 {
			gameBoard[coordI].player2HitShot = true
			gameBoard[coordI].player1Ship = false
			return true, nil
		} else {
			return false, errors.New("player not in range")
		}
	} else {
		return false, nil
	}
}
