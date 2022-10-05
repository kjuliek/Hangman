package shangman

import (
	"fmt"
	"hangman/printhangman"
)

func Deja_Affichee() {

}

func Pendu(hangman *Hangman) {
	var lettre string
	lettredanslemot := false

	fmt.Print("choisissez une lettre : ")
	fmt.Scan(&lettre)
	for i := 0; i < len(hangman.Mot_de_Depart); i++ {
		if lettre == string(hangman.Mot_de_Depart[i]) {
			lettredanslemot = true
			if Deja_Affichee() {
				hangman.Attempts--
				fmt.Println("Oh non ! La lettre a déjà été affiché, il te reste ", hangman.Attempts, " essais.")
				printhangman.PrintHangman(hangman.Attempts)
			} else {
				printhangman.PrintHangman(hangman.Attempts)
			}
		}

	}
	if !lettredanslemot {
		hangman.Attempts--
		fmt.Println("Il te reste ", hangman.Attempts, " essais.")
		printhangman.PrintHangman(hangman.Attempts)
	}

}
