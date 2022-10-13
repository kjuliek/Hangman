package shangman

import (
	"fmt"
	"hangman/printhangman"
	"hangman/printmot"
	"hangman/revealLetters"
)

func Deja_Affichee(Hangman *Hangman, lettre string) bool {
	for i := 0; i < len(Hangman.Mot_Afficher); i++ {
		if string(Hangman.Mot_Afficher[i]) == lettre {
			return true
		}
	}
	return false
}

func Pendu(Hangman *Hangman) {
	var lettre string
	lettredanslemot := false
	lettreajoutee := false

	fmt.Print("choisissez une lettre : ")
	fmt.Scan(&lettre)

	if lettre == "stop" {
		Hangman.InGame = false
	}

	for i := 0; i < len(Hangman.Mot_de_Depart) && !lettreajoutee && Hangman.InGame; i++ {
		if lettre == string(Hangman.Mot_de_Depart[i]) {
			lettredanslemot = true
			if Deja_Affichee(Hangman, lettre) {
				printmot.Printmot(Hangman.Mot_Afficher)
				Hangman.Attempts--
				fmt.Println("Oh non ! La lettre a déjà été affiché, il te reste ", Hangman.Attempts, " essais.")
				printhangman.PrintHangman(Hangman.Attempts)
			} else {
				Hangman.Mot_Afficher = revealLetters.AddLetter(lettre, Hangman.Mot_Afficher, Hangman.Mot_de_Depart)
				printmot.Printmot(Hangman.Mot_Afficher)
				printhangman.PrintHangman(Hangman.Attempts)
				lettreajoutee = true
			}
		}
	}

	if !lettredanslemot && Hangman.InGame {
		Hangman.Attempts--
		printmot.Printmot(Hangman.Mot_Afficher)
		fmt.Println("Il te reste ", Hangman.Attempts, " essais.")
		printhangman.PrintHangman(Hangman.Attempts)
	}
}
