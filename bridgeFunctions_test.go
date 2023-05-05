package main

import "testing"

func TestStringInSlice(t *testing.T) {
	alphabet := []string{"A", "B", "C", "D", "E"}
	test1R := stringInSlice("B", alphabet)
	test2R := stringInSlice("Q", alphabet)
	if !test1R {
		t.Errorf("\"B\" not found where it obviously is")
	} else if test2R {
		t.Errorf("\"Q\" found where it obviously is not")
	}
}

func TestPlayerShoot(t *testing.T) {
	// Uses bootStrapTest() and testTearDown() from gameBoardFunctions tests
	bootStrapTest()
	posStrings := []string{"D1", "B4"}
	posExpect := []bool{true, false}
	for i, input := range posStrings {
		output, err := playerShoot(1, input)
		if output != posExpect[i] || err != nil {
			if err != nil {
				t.Errorf("Error \"%v\" encountered", err)
			}
			t.Errorf("playerShoot(1, %v) returned %t", input, output)
		}
	}
	tearDownTest()
}
