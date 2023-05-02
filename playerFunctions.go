package main

import (
	"bufio"
	"os"
)

type userObject struct {
	userNum  int
	username string
}

func userInitialisation(userNum int) {
	useNumStr := string(userNum)
	inputReader := bufio.NewReader(os.Stdin)

}
