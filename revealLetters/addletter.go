package revealLetters

func AddLetter(letter, shown_word, starting_word string) string { //ajout des lettres du mots affiché à trou
	var new_word string

	for index, l := range shown_word {
		if letter == string(starting_word[index]) {
			new_word += letter
		} else {
			new_word += string(l)
		}
	}

	return new_word
}
