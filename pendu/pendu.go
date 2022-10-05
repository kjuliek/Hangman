package pendu

import "fmt"

func Pendu(mot string) string {
	var lettre string

	fmt.Print("choisissez une lettre : ")
	fmt.Scan(&lettre)
	for i := 0; i < len(mot); i++ {
		if lettre == string(mot[i]) {
			return "bonne réponse"
		}

	}
	return "mauvaise réponse"
}
