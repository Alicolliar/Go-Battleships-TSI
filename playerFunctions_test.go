package main

import "testing"

func TestRegisterShot(t *testing.T) {
	bootStrapTest()
	registerShot(1, true)
	if users[0].score != 1 && users[1].remainingShips != 4 {
		t.Errorf("registerShot(1, true) did not update values")
	}
	registerShot(1, false)
	if users[0].score != 1 && users[1].remainingShips != 4 {
		t.Errorf("registerShot(1, false) did not update values")
	}
	tearDownTest()
}
