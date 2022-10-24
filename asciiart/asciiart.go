package asciiart

import (
	"bufio"
	"log"
	"os"
)

func sort_word(word string) string {
	var ref []rune
	var newword string
	ref = append(ref, '_')
	for r := 97; r <= 122; r++ {
		ref = append(ref, rune(r))
	}
	for _, l := range ref {
		for _, letter := range word {
			if letter == l {
				if len(newword) == 0 || rune(newword[len(newword)-1]) != letter {
					newword += string(letter)
				}
			}
		}
	}
	return newword
}

func TabAsciiA(starting_word string) [][]string {
	newword := "_" + sort_word(starting_word)
	var tab [][]string

	text, err := os.Open("alphabet.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(text)

	var ascii []string
	for scanner.Scan() {
		ascii = append(ascii, scanner.Text())
	}

	line := 1
	for _, letter := range newword {
		var lettreascii []string
		cmpt := 1

		firstline := 1
		if string(letter) == "_" {
			firstline = 1
		} else {
			firstline = (int(rune(letter))-96)*12 + 1
		}
		for cmpt <= 12 {
			if line >= firstline && line <= firstline+12 {
				lettreascii = append(lettreascii, ascii[line])
				cmpt++
			}
			line++
		}
		tab = append(tab, lettreascii)

	}
	return tab
}
