package shangman

import (
	"fmt"
	"hangman/printhangman"
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

	fmt.Print("choisissez une lettre : ")
	fmt.Scan(&lettre)
	for i := 0; i < len(Hangman.Mot_de_Depart); i++ {
		if lettre == string(Hangman.Mot_de_Depart[i]) {
			lettredanslemot = true
			if Deja_Affichee(Hangman, lettre) {
				Hangman.Attempts--
				fmt.Println("Oh non ! La lettre a déjà été affiché, il te reste ", Hangman.Attempts, " essais.")
				printhangman.PrintHangman(Hangman.Attempts)
			} else {
				printhangman.PrintHangman(Hangman.Attempts)
			}
		}

	}
	if !lettredanslemot {
		Hangman.Attempts--
		fmt.Println("Il te reste ", Hangman.Attempts, " essais.")
		printhangman.PrintHangman(Hangman.Attempts)
	}

}
