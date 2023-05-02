package main

import (
	"errors"
	"fmt"
	"testing"
)

type testData struct {
	inputText      string
	expectedOutput int
	expectedError  error
}

// Tests for decodeCoordString
func TestDecodeCoordString(t *testing.T) {
	testData := []testData{{"B2", 6, nil}, {"Z4", 0, errors.New(`incorrect X-Coord`)}, {"B19", 0, errors.New(`string too long`)}}
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
