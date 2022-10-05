package chooseword

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func ChooseWord() string {
	var words []string
	rand.Seed(time.Now().UTC().UnixNano())
	mots, err := os.Open("Mots.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(mots)

	nbwords := 0

	for scanner.Scan() {
		nbwords++
		words = append(words, scanner.Text())
	}
	indexword := rand.Intn(nbwords)

	return words[indexword]
}
