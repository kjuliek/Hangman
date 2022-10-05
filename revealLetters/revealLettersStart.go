package revealLetters

import (
	"fmt"
	"math/rand"
	"time"
)

func ReveallettersStart(Mot string) {

	Nb_Lettre_To_Reveal := (len(Mot) / 2) - 1
	Position_deja_reveal := []int{}
	rand.Seed(time.Now().UTC().UnixNano())
	Verification := true

	for i := 0; i < Nb_Lettre_To_Reveal; i++ {
		Pose_Tmp := rand.Intn(len(Mot))
		Verification = true

		for j := 0; j < len(Position_deja_reveal); j++ {
			if Pose_Tmp == Position_deja_reveal[j] {
				Verification = false
				i--
			}
		}

		if Verification {
			Position_deja_reveal = append(Position_deja_reveal, Pose_Tmp)
			fmt.Println(string(Mot[Pose_Tmp]))
		}
	}
	fmt.Println(Position_deja_reveal)

	for i := 0; i < len(Mot); i++ {
		Verification = true
		for j := 0; j < len(Position_deja_reveal); j++ {
			if Mot[i] == Mot[Position_deja_reveal[j]] {
				fmt.Print(string(Mot[i]))
				Verification = false
			}
		}
		if Verification {
			fmt.Printf("_")
		}
		fmt.Printf(" ")
	}
}
