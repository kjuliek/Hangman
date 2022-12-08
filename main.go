package main

import (
	"Hangman-Classic/startingGame"
	"os"
)

func main() { //display of the save,  win or lost game
	startingGame.Start(os.Args)
}
