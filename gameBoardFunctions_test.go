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

func TestDecodeCoordString(t *testing.T) {
	// Tests for decodeCoordString
	testData := []testData{{"B2", 6, nil}, {"Z4", 0, errors.New(`incorrect X-Coord`)}, {"B19", 0, errors.New(`string too long`)}, {"E7", 0, errors.New(`incorrect Y-Coord`)}}
	for _, input := range testData {
		testOutput, err := decodeCoordString(input.inputText)
		fmt.Printf("decodeCoordString(\"%v\") = %d\n", input.inputText, testOutput)
		if testOutput != input.expectedOutput {
			t.Errorf("want %d\n", input.expectedOutput)
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
		t.Errorf("Expected %v,\n\n got %v", expectedOutput, actualOutput)
	}
}
