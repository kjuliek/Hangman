package printhangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintHangman(nbattempts int) {
	if nbattempts != 10 {
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
