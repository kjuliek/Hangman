package shangman

import (
	"hangman/Hangman-Classic/asciiart"
	"hangman/Hangman-Classic/chooseword"
	"hangman/Hangman-Classic/revealLetters"
)

type Hangman struct { // every type of variables that are created for the code
	Starting_Word     string
	Displayed_Word    string
	Letters_Attempted string
	Attempts          int
	Letters_to_Find   int
	InGame            bool
	Win               bool
	Letters_Ascii_Art [][]string
}

func InitHangman() *Hangman { //word chosen randomly in words.txt
	word := chooseword.ChooseWord()
	displayed_Word := revealLetters.ReveallettersStart(word) //letters chosen randomly
	letters_ascii_art := asciiart.TabAsciiA(word)            //implamantation of the ASCII letters
	return &Hangman{                                         //display of the word
		Starting_Word:     word,
		Displayed_Word:    displayed_Word,
		Letters_Attempted: "",
		Attempts:          10,
		Letters_to_Find:   len(word),
		InGame:            true,
		Win:               false,
		Letters_Ascii_Art: letters_ascii_art,
	}
}
