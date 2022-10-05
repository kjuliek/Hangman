package main

import (
	"fmt"
	"hangman/shangman"
)

func main() {

	var Hangman *shangman.Hangman = new(shangman.Hangman)

	Hangman = shangman.InitHangman()
	fmt.Println(Hangman.Mot_de_Depart)
	fmt.Println(Hangman.Mot_Afficher)
}
