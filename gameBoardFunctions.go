package main

import (
	"errors"
	"strconv"
	"strings"
)

type CoordinatePoint struct {
	coordString    string
	player2Ship    bool
	player2HitShot bool
	player1Ship    bool
	player1HitShot bool
}

var alphabet = []string{"A", "B", "C", "D", "E"}

var xLength int

var gameBoard [25]CoordinatePoint

func decodeCoordString(coordString string) (int, error) {
	coordList := strings.Split(coordString, "")
	if len(coordList) > 2 || len(coordList) <= 1 {
		return 0, errors.New("coord string wrong length")
	}
	firstSection := (int(coordString[0]) - 65)
	secondSection := (int(coordString[1]) - 49)
	if firstSection < 0 || firstSection > 4 {
		return 0, errors.New(`invalid X-Coord`)
	} else if secondSection < 0 || secondSection > 4 {
		return 0, errors.New(`invalid Y-Coord`)
	}
	return (((firstSection) * 5) + secondSection), nil
}

func createGameBoard() [25]CoordinatePoint {
	xLength = len(alphabet)
	var gameBoardCreation [25]CoordinatePoint
	var indexPoint int
	totalRemainingShips = 10
	for _, letter := range alphabet {
		for i := 1; i <= xLength; i++ {
			combinedCoord := letter + strconv.Itoa(i)
			newBoardPoint := CoordinatePoint{coordString: combinedCoord, player2Ship: false, player1Ship: false, player2HitShot: false, player1HitShot: false}
			gameBoardCreation[indexPoint] = newBoardPoint
			indexPoint++
		}
	}
	return gameBoardCreation
}

func checkOccupation(playerNum int, coordIndex int) (bool, error) {
	currentPoint := gameBoard[coordIndex]
	if playerNum == 1 {
		return currentPoint.player2Ship, nil
	} else if playerNum == 2 {
		return currentPoint.player1Ship, nil
	} else if playerNum == 0 {
		return currentPoint.player1Ship || currentPoint.player2Ship, nil
	}
	return false, errors.New("player not in range")
}

func playerOccupation(playerNum int, positionString string) error {
	coordIndex, err := decodeCoordString(positionString)
	if err != nil {
		return err
	}
	occupCheck := false
	occupCheck, err = checkOccupation(0, coordIndex)
	if err != nil {
		return err
	}
	if occupCheck {
		return errors.New("position already occupied")
	}
	if playerNum == 1 {
		gameBoard[coordIndex].player1Ship = true
		return nil
	} else if playerNum == 2 {
		gameBoard[coordIndex].player2Ship = true
		return nil
	} else {
		return errors.New("player not in range")
	}
}
