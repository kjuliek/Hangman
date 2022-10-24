package printmot

import (
	"fmt"
	"hangman/asciiart"
)

func Printmot(mot string) {
	for i := 0; i < len(mot); i++ {
		fmt.Print(string(mot[i]), " ")
	}
	fmt.Println()
}

func PrintAsciiArt(mot_de_depart, mot_afficher string, lettres_ascii_art [][]string) {
	mot_trie := "_" + asciiart.TriMot(mot_de_depart)
	for k := 0; k < 11; k++ {
		for _, letter := range mot_afficher {
			for i, l := range mot_trie {
				if letter == l {
					fmt.Print(lettres_ascii_art[i][k])
				}
			}
		}
		fmt.Println()
	}
}
