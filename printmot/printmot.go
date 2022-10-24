package printword

import (
	"fmt"
	"hangman/asciiart"
)

func Printmot(word string) {
	for i := 0; i < len(word); i++ {
		fmt.Print(string(word[i]), " ")
	}
	fmt.Println()
}

func PrintAsciiArt(starting_word, shown_word string, letter_ascii_art [][]string) {
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
