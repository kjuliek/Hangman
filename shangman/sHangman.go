package shangman

import (
	"hangman/asciiart"
	"hangman/chooseword"
	"hangman/revealLetters"
)

type Hangman struct {
	Starting_Word     string
	Displayed_Word    string
	Letters_Attempted string
	Attempts          int
	Letters_to_Find   int
	InGame            bool
	Win               bool
	Letters_Ascii_Art [][]string
}

func InitHangman() *Hangman {
	word := chooseword.ChooseWord()
	displayed_Word := revealLetters.ReveallettersStart(word)
	letters_ascii_art := asciiart.TabAsciiA(word)
	return &Hangman{
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
