package main

import (
	"github.com/kjuliek/Hangman/startingGame"
	"os"
)

func main() { //display of the save,  win or lost game
	startingGame.Start(os.Args)
}
