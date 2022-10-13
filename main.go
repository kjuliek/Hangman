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
	} else {
		Hangman = shangman.InitHangman()
		printmot.Printmot(Hangman.Mot_Afficher)
	}

	for Hangman.Mot_Afficher != Hangman.Mot_de_Depart && Hangman.Attempts > 0 && Hangman.InGame {
		shangman.Pendu(Hangman)
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
