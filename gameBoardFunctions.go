package main

import (
	"errors"
	"strconv"
)

type CoordinatePoint struct {
	coordString    string
	player2Ship    bool
	player2HitShot bool
	player1Ship    bool
	player1HitShot bool
}

var gameBoard [25]CoordinatePoint

func decodeCoordString(coordString string) (int, error) {
	if len(coordString) > 2 {
		return 0, errors.New("string too long")
	}
	firstSection := (int(coordString[0]) - 64)
	secondSection := (int(coordString[1]) - 48)
	if firstSection < 0 || firstSection > 4 {
		return 0, errors.New(`incorrect X-Coord`)
	} else if secondSection < 0 || secondSection > 4 {
		return 0, errors.New(`incorrect Y-Coord`)
	}
	return ((((firstSection - 1) * 5) + secondSection) - 1), nil
}

func createGameBoard() [25]CoordinatePoint {
	var gameBoardCreation [25]CoordinatePoint
	alphabet := [5]string{"A", "B", "C", "D", "E"}
	var indexPoint int
	for _, letter := range alphabet {
		for i := 1; i <= len(alphabet); i++ {
			combinedCoord := letter + strconv.Itoa(i)
			newBoardPoint := CoordinatePoint{coordString: combinedCoord, player2Ship: false, player1Ship: false, player2HitShot: false, player1HitShot: false}
			gameBoardCreation[indexPoint] = newBoardPoint
			indexPoint++
		}
	}
	return gameBoardCreation
}

func checkOccupation(playerNum int, positionString string) (error, bool) {
	coordIndex, err := decodeCoordString(positionString)
	if err != nil {
		return err, false
	}
	currentPoint := gameBoard[coordIndex]
	if playerNum == 1 {
		return nil, currentPoint.player1Ship
	} else if playerNum == 2 {
		return nil, currentPoint.player1Ship
	} else if playerNum == 0 {
		return nil, currentPoint.player1Ship || currentPoint.player2Ship
	}
	return errors.New("player not in range"), false
}

func playerOccupation(playerNum int, positionString string) error {
	coordIndex, err := decodeCoordString(positionString)
	if err != nil {
		return err
	}
	if playerNum == 1 {
		gameBoard[coordIndex].player1Ship = true
		return nil
	} else if playerNum == 2 {
		gameBoard[coordIndex].player2Ship = true
		return nil
	}
	return errors.New("player not in range")
}
