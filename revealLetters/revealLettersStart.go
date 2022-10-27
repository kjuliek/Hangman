package revealLetters

import (
	"math/rand"
	"time"
)

func ReveallettersStart(word string) string { //affichage des lettres
	new_word := ""
	Nb_Letter_To_Reveal := (len(word) / 2) - 1 //nombre de lettres révélé, longueur du mot divisé par 2 moins 1
	Position_reveal := []int{}                 //position de la lettre révélé
	rand.Seed(time.Now().UTC().UnixNano())
	Verification := true

	for i := 0; i < Nb_Letter_To_Reveal; i++ { //incrémentation de la lettre dans le mot
		Pose_Time := rand.Intn(len(word))
		Verification = true
		for j := 0; j < len(Position_reveal); j++ {
			if Pose_Time == Position_reveal[j] {
				Verification = false
				i--
			}
		}
		if Verification {
			Position_reveal = append(Position_reveal, Pose_Time)
		}
	}

	for i := 0; i < len(word); i++ { //prends le mot et ajoute les tirets
		Verification = true
		for j := 0; j < len(Position_reveal); j++ {
			if word[i] == word[Position_reveal[j]] {
				new_word = new_word + string(word[i])
				Verification = false
				break
			}
		}
		if Verification {
			new_word = new_word + "_"
		}
	}

	return new_word
}
