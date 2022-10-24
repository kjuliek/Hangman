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
	word, err := os.Open(string(os.Args[1]))

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(word)
	nbwords := 0

	for scanner.Scan() {
		nbwords++
		words = append(words, scanner.Text())
	}
	indexword := rand.Intn(nbwords)

	return words[indexword]
}
