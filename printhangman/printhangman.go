package printhangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintHangman(nbattempts int) { //affichage du nombre de tentative du joueur avant de mourir
	if nbattempts != 10 { // nombre de tentative
		firstline := (9-nbattempts)*7 + 1
		text, err := os.Open("Hangman.txt")

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(text)

		line := 1

		for scanner.Scan() {
			if line >= firstline && line <= firstline+6 {
				fmt.Println(scanner.Text())
			} else {
				scanner.Text()
			}
			line++
		}
	}
}
