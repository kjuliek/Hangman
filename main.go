package main

import (
	"encoding/json"
	"fmt"
	"hangman/printmot"
	"hangman/shangman"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("\033[H\033[2J")
	if len(os.Args) > 2 {
		fmt.Println("Trop d'information en parammettre")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("Il manque le nom du fichier txt pour recuperer les mots")
		return
	}

	var Hangman *shangman.Hangman = new(shangman.Hangman)

	if string(os.Args[1]) == "Save.json" || string(os.Args[1]) == ".\\Save.json" {
		data, _ := ioutil.ReadFile("Save.json")
		json.Unmarshal(data, Hangman)
		printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_Afficher, Hangman.Lettres_Ascii_Art)
	} else {
		Hangman = shangman.InitHangman()
		printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_Afficher, Hangman.Lettres_Ascii_Art)
	}

	for Hangman.Mot_Afficher != Hangman.Mot_de_Depart && Hangman.Attempts > 0 && Hangman.InGame && !Hangman.Win {
		shangman.Pendu(Hangman)
	}

	if Hangman.Mot_Afficher == Hangman.Mot_de_Depart || Hangman.Win {
		fmt.Println("Félicitation ! Tu as gagné !")
	} else if !Hangman.InGame {
		fmt.Println("La partie a éte sauvegardée")
		fmt.Println("Pour recharger la partie faites : go run .\\main.go .\\Save.json")
	} else {
		fmt.Println("Oh non ! Tu es dcd !")
		fmt.Println("Le mot était :")
		printmot.PrintAsciiArt(Hangman.Mot_de_Depart, Hangman.Mot_de_Depart, Hangman.Lettres_Ascii_Art)
	}
}
