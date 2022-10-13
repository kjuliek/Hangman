package shangman

import (
	"encoding/json"
	"fmt"
	"hangman/printhangman"
	"hangman/printmot"
	"hangman/revealLetters"
	"io/ioutil"
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
	lettreDansLeMot := false
	lettreAjoutee := false
	fmt.Println("Si vous souhaitez quitter la partie écrivez : stop")
	fmt.Print("Choisissez une lettre : ")
	fmt.Scan(&lettre)

	if lettre == "stop" {
		file, _ := json.Marshal(Hangman)
		_ = ioutil.WriteFile("Save.json", file, 0644)
		Hangman.InGame = false
	}

	for i := 0; i < len(Hangman.Mot_de_Depart) && !lettreAjoutee && Hangman.InGame; i++ {
		if lettre == string(Hangman.Mot_de_Depart[i]) {
			lettreDansLeMot = true
			if Deja_Affichee(Hangman, lettre) {
				printmot.Printmot(Hangman.Mot_Afficher)
				Hangman.Attempts--
				fmt.Println("Oh non ! La lettre a déjà été affiché, il te reste ", Hangman.Attempts, " essais.")
				printhangman.PrintHangman(Hangman.Attempts)
			} else {
				Hangman.Mot_Afficher = revealLetters.AddLetter(lettre, Hangman.Mot_Afficher, Hangman.Mot_de_Depart)
				printmot.Printmot(Hangman.Mot_Afficher)
				printhangman.PrintHangman(Hangman.Attempts)
				lettreAjoutee = true
			}
		}
	}

	if !lettreDansLeMot && Hangman.InGame {
		Hangman.Attempts--
		printmot.Printmot(Hangman.Mot_Afficher)
		fmt.Println("Il te reste ", Hangman.Attempts, " essais.")
		printhangman.PrintHangman(Hangman.Attempts)
	}
}
