package revealLetters

import (
	"math/rand"
	"time"
)

func ReveallettersStart(mot string) string {

	new_mot := ""
	Nb_Lettre_To_Reveal := (len(mot) / 2) - 1
	Position_deja_reveal := []int{}
	rand.Seed(time.Now().UTC().UnixNano())
	Verification := true

	for i := 0; i < Nb_Lettre_To_Reveal; i++ {
		Pose_Tmp := rand.Intn(len(mot))
		Verification = true

		for j := 0; j < len(Position_deja_reveal); j++ {
			if Pose_Tmp == Position_deja_reveal[j] {
				Verification = false
				i--
			}
		}

		if Verification {
			Position_deja_reveal = append(Position_deja_reveal, Pose_Tmp)
			// fmt.Println(string(Mot[Pose_Tmp]))
		}
	}
	// fmt.Println(Position_deja_reveal)

	for i := 0; i < len(mot); i++ {
		Verification = true
		for j := 0; j < len(Position_deja_reveal); j++ {
			if mot[i] == mot[Position_deja_reveal[j]] {
				// fmt.Print(string(Mot[i]))
				new_mot = new_mot + string(mot[i])
				Verification = false
				break
			}
		}
		if Verification {
			// fmt.Printf("_")
			new_mot = new_mot + "_"
		}
		// fmt.Printf(" ")
		new_mot = new_mot + " "
	}
	// fmt.Println()
	return new_mot
}
