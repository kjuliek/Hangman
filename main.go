package main

import (
	"encoding/json"
	"fmt"
	"hangman/printword"
	"hangman/shangman"
	"io/ioutil"
	"os"
)

func main() { //display of the save,  win or lost game
	fmt.Print("\033[H\033[2J")
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments.")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("Text file is missing.")
		return
	}

	var Hangman *shangman.Hangman = new(shangman.Hangman)

	if string(os.Args[1]) == "Save.json" || string(os.Args[1]) == ".\\Save.json" {
		data, _ := ioutil.ReadFile("Save.json")
		json.Unmarshal(data, Hangman)
		printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Displayed_Word, Hangman.Letters_Ascii_Art)
	} else {
		Hangman = shangman.InitHangman()
		printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Displayed_Word, Hangman.Letters_Ascii_Art)
	}

	for Hangman.Displayed_Word != Hangman.Starting_Word && Hangman.Attempts > 0 && Hangman.InGame && !Hangman.Win {
		shangman.Pendu(Hangman)
	}

	if Hangman.Displayed_Word == Hangman.Starting_Word || Hangman.Win { //winned game
		fmt.Println("Congratulation ! You win !")
	} else if !Hangman.InGame { //save
		fmt.Println("The game has been saved.")
		fmt.Println("To resume the game, do : go run .\\main.go .\\Save.json")
	} else { //death of the player
		fmt.Println("Oh no ! You died !")
		fmt.Println("The word was :")
		printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Starting_Word, Hangman.Letters_Ascii_Art)
	}
}
