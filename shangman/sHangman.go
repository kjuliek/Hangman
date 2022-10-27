package shangman

import (
	"hangman/asciiart"
	"hangman/chooseword"
	"hangman/revealLetters"
)

type Hangman struct { // tout les types de variable creer pour le code
	Starting_Word     string //mot de départ
	Displayed_Word    string //mot affiché
	Letters_Attempted string //lettres tentées
	Attempts          int    //essaie
	Letters_to_Find   int    //lettres manquantes
	InGame            bool
	Win               bool
	Letters_Ascii_Art [][]string
}

func InitHangman() *Hangman { //mot choisi aléatoirement dans Words.txt
	word := chooseword.ChooseWord()
	displayed_Word := revealLetters.ReveallettersStart(word) //lettres choisi aléatoirement
	letters_ascii_art := asciiart.TabAsciiA(word)            //implantation des lettres ascii
	return &Hangman{                                         //affichage du mot
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
