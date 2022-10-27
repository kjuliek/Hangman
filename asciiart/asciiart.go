package asciiart

import (
	"bufio"
	"log"
	"os"
)

func SortWord(word string) string { //recupération des lettres alphabétiques appartenants au mot choisi
	var ref []rune
	var newword string
	ref = append(ref, '_')
	for r := 97; r <= 122; r++ { // de a à z
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

func TabAsciiA(starting_word string) [][]string { // changements des lettres en ascii, incrémentations dans un tableau
	newword := "_" + SortWord(starting_word)
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
	for _, letter := range newword { //ligne par ligne on inscrit une lettre ascii
		var lettreascii []string
		cmpt := 1

		firstline := 1
		if string(letter) == "_" {
			firstline = 1
		} else {
			firstline = (int(rune(letter))-96)*12 + 1 //calcul pour afficher une lettre ascii
		}
		for cmpt <= 12 { //toute les 12 lignes une nouvelle lettre apparait
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
