package shangman

import (
	"hangman/asciiart"
	"hangman/chooseword"
	"hangman/revealLetters"
)

type Hangman struct {
	Mot_de_Depart     string
	Mot_Afficher      string
	Lettres_Tentees   string
	Attempts          int
	Lettres_a_Trouver int
	InGame            bool
	Lettres_Ascii_Art [][]string
}

func InitHangman() *Hangman {
	mot := chooseword.ChooseWord()
	mot_Afficher := revealLetters.ReveallettersStart(mot)
	lettres_ascii_art := asciiart.TabAsciiA(mot)
	return &Hangman{
		Mot_de_Depart:     mot,
		Mot_Afficher:      mot_Afficher,
		Lettres_Tentees:   "",
		Attempts:          10,
		Lettres_a_Trouver: len(mot),
		InGame:            true,
		Lettres_Ascii_Art: lettres_ascii_art,
	}
}
