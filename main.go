package main

import (
	"fmt"
	"hangman/printmot"
	"hangman/shangman"
)

func main() {

	var Hangman *shangman.Hangman = new(shangman.Hangman)

	Hangman = shangman.InitHangman()
	printmot.Printmot(Hangman.Mot_de_Depart)
	printmot.Printmot(Hangman.Mot_Afficher)
	fmt.Println(Hangman.Mot_de_Depart)
	fmt.Println(Hangman.Mot_Afficher)
}
