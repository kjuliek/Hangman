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

	for Hangman.Mot_Afficher != Hangman.Mot_de_Depart && Hangman.Attempts > 0 && Hangman.InGame {
		shangman.Pendu(Hangman)
		// fmt.Print("\033[H\033[2J")
	}
	if Hangman.Mot_Afficher == Hangman.Mot_de_Depart {
		fmt.Println("Félicitation ! Tu as gagné !")
	} else if !Hangman.InGame {
		fmt.Println("La partie a éte sauvegardée")
	} else {
		fmt.Println("Oh non ! Tu es dcd !")
		fmt.Println("Le mot était :", Hangman.Mot_de_Depart)
	}
}
