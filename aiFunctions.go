package main

import (
	"fmt"
	"math"
	"math/rand"
)

func aiInitialisation(aiPlayerNum int) {
	fmt.Println("Initialising AI Player")
	newUserObject := userObject{}
	newUserObject.username = "Computer"
	newUserObject.userNum = aiPlayerNum
	newUserObject.remainingShips = 5
	users[1] = newUserObject
	aiPlaceships(aiPlayerNum)
}

func genRandomCoords() string {
	coordsOk := false
	var randCoordString string
	for !coordsOk {
		xCoord := int(math.Round(rand.Float64() * 4))
		yCoord := int(math.Round(rand.Float64()*4)) + 1
		randCoordString = fmt.Sprintf("%s%d", string(alphabet[xCoord]), yCoord)
		_, err := decodeCoordString(randCoordString)
		if err != nil {
			coordsOk = false
		}
		coordsOk = true
	}
	return randCoordString
}

func aiPlaceships(aiPlayerNum int) {
	for i := 0; i <= 5; i++ {
		coordsOk := false
		for !coordsOk {
			coordString := genRandomCoords()
			err := playerOccupation(aiPlayerNum, coordString)
			if err != nil {
				coordsOk = false
				continue
			}
			coordsOk = true
		}
	}
}

func aiPlayerTurn(aiPlayerNum int) {
	coords := genRandomCoords()
	result, _ := playerShoot(aiPlayerNum, coords)
	fmt.Print(result)
	if !result {
		fmt.Printf("%v was a miss!\n", coords)
		registerShot(aiPlayerNum, false)
	} else if result {
		fmt.Printf("%v was a hit! Ship sunk\n", coords)
		registerShot(aiPlayerNum, true)
	}
}
