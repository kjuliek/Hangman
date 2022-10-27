package printword

import (
	"fmt"
	"hangman/asciiart"
)

func PrintWord(word string) { //choice of the word and display of it
	for i := 0; i < len(word); i++ {
		fmt.Print(string(word[i]), " ")
	}
	fmt.Println()
}

func PrintAsciArt(starting_word, shown_word string, letter_ascii_art [][]string) { //incrementation of the asciiart when the word is showed
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
