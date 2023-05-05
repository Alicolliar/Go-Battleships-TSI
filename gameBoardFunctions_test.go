package main

import (
	"errors"
	"fmt"
	"testing"
)

type testData struct {
	// Used for TestDecodeCoordString currently
	inputText      string
	expectedOutput int
	expectedError  error
}

type checkOccupData struct {
	// Used for TestCheckOccupation
	userInt         int
	testInputs      []int
	expectedOutputs []bool
}

var testBoard [25]CoordinatePoint

func bootStrapTest() {
	testBoard = createGameBoard()
	testBoard[13].player1Ship = true
	testBoard[15].player2Ship = true
	testBoard[20].player1HitShot = true
	testBoard[17].player2HitShot = true
	gameBoard = testBoard
	totalRemainingShips = 10
	users[0] = userObject{userNum: 1, username: "ATestUser", remainingShips: 5}
	users[1] = userObject{userNum: 2, username: "Computer", remainingShips: 5}
}

func tearDownTest() {
	testBoard = [25]CoordinatePoint{}
	totalRemainingShips = 0
	users = [2]userObject{}
}

func TestDecodeCoordString(t *testing.T) {
	// Tests for decodeCoordString
	testData := []testData{{"B2", 6, nil}, {"Z4", 0, errors.New(`incorrect X-Coord`)}, {"B19", 0, errors.New(`string too long`)}, {"E7", 0, errors.New(`incorrect Y-Coord`)}}
	for _, input := range testData {
		testOutput, err := decodeCoordString(input.inputText)
		fmt.Printf("decodeCoordString(\"%v\") = %d\n", input.inputText, testOutput)
		if testOutput != input.expectedOutput {
			t.Errorf("Want %d\n", input.expectedOutput)
		} else if input.expectedError != nil && err == nil {
			t.Errorf("Got no error message,  wanted %d\n", input.expectedError)
		} else if input.expectedError == nil && err != nil {
			t.Errorf("Got error message: %d,  wanted none", err)
		}
	}
}

func TestCreateGameBoard(t *testing.T) {
	// Tests for createGameBoard
	expectedOutput := [25]CoordinatePoint{{"A1", false, false, false, false}, {"A2", false, false, false, false}, {"A3", false, false, false, false}, {"A4", false, false, false, false}, {"A5", false, false, false, false}, {"B1", false, false, false, false}, {"B2", false, false, false, false}, {"B3", false, false, false, false}, {"B4", false, false, false, false}, {"B5", false, false, false, false}, {"C1", false, false, false, false}, {"C2", false, false, false, false}, {"C3", false, false, false, false}, {"C4", false, false, false, false}, {"C5", false, false, false, false}, {"D1", false, false, false, false}, {"D2", false, false, false, false}, {"D3", false, false, false, false}, {"D4", false, false, false, false}, {"D5", false, false, false, false}, {"E1", false, false, false, false}, {"E2", false, false, false, false}, {"E3", false, false, false, false}, {"E4", false, false, false, false}, {"E5", false, false, false, false}}
	actualOutput := createGameBoard()
	if expectedOutput != actualOutput {
		t.Errorf("Expected %v,\ngot %v", expectedOutput, actualOutput)
	}
}

func TestCheckOccupation(t *testing.T) {
	// Tests for checkOccupation
	bootStrapTest()
	testInput := []checkOccupData{{userInt: 1, testInputs: []int{13, 15, 22}, expectedOutputs: []bool{false, true, false}}, {userInt: 2, testInputs: []int{13, 15, 22}, expectedOutputs: []bool{true, false, false}}, {userInt: 0, testInputs: []int{13, 15, 22}, expectedOutputs: []bool{true, true, false}}}
	for _, testData := range testInput {
		for i, testIns := range testData.testInputs {
			check, err := checkOccupation(testData.userInt, testIns)
			if check != testData.expectedOutputs[i] || err != nil {
				t.Errorf("checkOccupation(%d, %d) = %t, expected %t", testData.userInt, testIns, check, testData.expectedOutputs[i])
				if err != nil {
					fmt.Printf("Error occurred: %v", err)
				}
			}
		}
	}
	tearDownTest()
}

func TestPlayerOccupation(t *testing.T) {
	bootStrapTest()
	occupStruct := struct {
		coordString  []string
		trueIndex    []int
		errorPresent []bool
	}{[]string{"C4", "D1", "C2"}, []int{13, 15, 11}, []bool{true, true, false}}
	for i, input := range occupStruct.coordString {
		err := playerOccupation(1, input)
		spaceOccupation := gameBoard[occupStruct.trueIndex[i]].player1Ship
		if !occupStruct.errorPresent[i] && err != nil {
			t.Errorf("checkOccupation(1, %v) => Error \"%v\" occurred", input, err)
		} else if !spaceOccupation && !occupStruct.errorPresent[i] {
			t.Errorf("Space %v was unoccupied after attempting occuppation", input)
		}
	}
	tearDownTest()
}
