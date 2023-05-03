package main

import "fmt"

func displayBoardStates(userNum int) {
	xCounter := 1
	var shipContaining, hitContaining []string
	for _, coordPoint := range gameBoard {
		if xCounter%xLength == 0 {
			fmt.Print('\n')
		}
		fmt.Printf("%v", coordPoint.coordString)
		if userNum == 1 {
			if coordPoint.player1Ship {
				shipContaining = append(shipContaining, coordPoint.coordString)
			} else if coordPoint.player1HitShot {
				hitContaining = append(hitContaining, coordPoint.coordString)
			}
		}
		xCounter++
	}
}
