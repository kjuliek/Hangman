package readtxt

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintHangman(nbattempts int) {
	//voir si obligé de mettre lastline aussi vu qu'ils feront tous la même taille
	if nbattempts != 10 {
		firstline := (9-nbattempts)*7 + 1
		text, err := os.Open("Pendus.txt")

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
