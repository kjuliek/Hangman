package shangman

import "fmt"

func Pendu(hangman *Hangman) {
	var lettre string

	fmt.Print("choisissez une lettre : ")
	fmt.Scan(&lettre)
	for i := 0; i < len(hangman.Mot_de_Depart); i++ {
		if lettre == string(hangman.Mot_de_Depart[i]) {
			if Deja_Affichee() {
				hangman.Attempts--
				printhangman.PrintHangman(hangman.Attempts)
			}
		}

	}
	return "mauvaise réponse"
}
