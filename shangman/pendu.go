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
	fmt.Println("Lettres tenté :", Hangman.Lettres_Tentees)
	fmt.Print("Choisissez une lettre qui n'as pas deja étais tenté : ")
	fmt.Scan(&lettre)
	fmt.Print("\033[H\033[2J")

	if lettre == "stop" {
		file, _ := json.Marshal(Hangman)
		_ = ioutil.WriteFile("Save.json", file, 0644)
		Hangman.InGame = false
	} else if lettre == Hangman.Mot_de_Depart {
		fmt.Println("Bravo , vous avez gagnez")
		Hangman.InGame = false
	} else if len(lettre) > 1 {
		Hangman.Attempts = Hangman.Attempts - 1
	} else if len(lettre) == 1 {
		Found := false

		for i := 0; i < len(Hangman.Lettres_Tentees); i++ {
			if lettre == string(Hangman.Lettres_Tentees[i]) {
				Found = true
			}
		}

		if Found {
			fmt.Println("La lettre a deja été tenté, veuillez réessayer")
			Hangman.Attempts++
		} else {
			Hangman.Lettres_Tentees = Hangman.Lettres_Tentees + lettre + " "
		}
	}

	for i := 0; i < len(Hangman.Mot_de_Depart) && !lettreAjoutee && Hangman.InGame; i++ {
		if lettre == string(Hangman.Mot_de_Depart[i]) {
			lettreDansLeMot = true

			if Deja_Affichee(Hangman, lettre) {
				printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_Afficher, Hangman.Lettres_Ascii_Art)
				Hangman.Attempts--
				fmt.Println("Oh non ! La lettre a déjà été affiché, il te reste ", Hangman.Attempts, " essais.")
				printhangman.PrintHangman(Hangman.Attempts)
			} else {
				Hangman.Mot_Afficher = revealLetters.AddLetter(lettre, Hangman.Mot_Afficher, Hangman.Mot_de_Depart)
				printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_Afficher, Hangman.Lettres_Ascii_Art)
				printhangman.PrintHangman(Hangman.Attempts)
				lettreAjoutee = true
			}
		}
	}

	if !lettreDansLeMot && Hangman.InGame {
		Hangman.Attempts--
		printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_Afficher, Hangman.Lettres_Ascii_Art)
		fmt.Println("Il te reste ", Hangman.Attempts, " essais.")
		printhangman.PrintHangman(Hangman.Attempts)
	}
}
