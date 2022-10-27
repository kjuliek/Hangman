package printword

import (
	"fmt"
	"hangman/asciiart"
)

func PrintWord(word string) { //choix du mot et affichage de celui-ci
	for i := 0; i < len(word); i++ {
		fmt.Print(string(word[i]), " ")
	}
	fmt.Println()
}

func PrintAsciArt(starting_word, shown_word string, letter_ascii_art [][]string) { //incrémentation de asciiart dans l'affichage du mot
	sort_word := "_" + asciiart.SortWord(starting_word)
	for k := 0; k < 12; k++ {
		for _, letter := range shown_word {
			for i, l := range sort_word {
				if letter == l {
					fmt.Print(letter_ascii_art[i][k])
				}
			}
		}
		fmt.Println()
	}
}
