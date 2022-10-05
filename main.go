package main

import (
	"fmt"
	"hangman/printmot"
	"hangman/shangman"
)

func main() {

	var Hangman *shangman.Hangman = new(shangman.Hangman)

	Hangman = shangman.InitHangman()
	printmot.Printmot(Hangman.Mot_Afficher)
	for Hangman.Mot_Afficher != Hangman.Mot_de_Depart && Hangman.Attempts > 0 {
		shangman.Pendu(Hangman)
	}
	if Hangman.Mot_Afficher == Hangman.Mot_de_Depart {
		fmt.Println("Félicitation ! Tu as gagné !")
	} else {
		fmt.Println("Oh non ! Tu es dcd !")
	}
}
