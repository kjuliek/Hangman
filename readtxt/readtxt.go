package readtxt

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadTxt(firstline, lastline int) {
	//voir si obligé de mettre lastline aussi vu qu'ils feront tous la même taille
	text, err := os.Open("Pendus.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(text)

	line := 1

	for scanner.Scan() {
		if line >= firstline && line <= lastline {
			fmt.Println(scanner.Text())
		} else {
			scanner.Text()
		}
		line++
	}
}
