package revealLetters

func AddLetter(letter, Mot_Afficher, Mot_de_Depart string) string {
	var new_mot string
	for index, l := range Mot_Afficher {
		if letter == string(Mot_de_Depart[index]) {
			new_mot += letter
		} else {
			new_mot += string(l)
		}
	}

	return new_mot
}
