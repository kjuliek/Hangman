package asciiart

import (
	"bufio"
	"log"
	"os"
)

func TriMot(word string) string {
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

func TabAsciiA(Mot_de_Depart string) [][]string {
	newword := "_" + TriMot(Mot_de_Depart)
	var tab [][]string

	text, err := os.Open("alphabet.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(text)

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
		for scanner.Scan() && cmpt < 13 {

			if line >= firstline && line <= firstline+12 {
				lettreascii = append(lettreascii, scanner.Text())
				cmpt++
			} else {
				scanner.Text()
			}
			line++
		}
		tab = append(tab, lettreascii)
	}
	return tab
}

func AsciiArt(mot string) {

}
