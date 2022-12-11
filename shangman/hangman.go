package shangman

import (
	"encoding/json"
	"fmt"
	"hangman/Hangman-Classic/printhangman"
	"hangman/Hangman-Classic/printword"
	"hangman/Hangman-Classic/revealLetters"
	"io/ioutil"
)

func Allready_Displayed(Hangman *Hangman, letter string) bool {
	for i := 0; i < len(Hangman.Displayed_Word); i++ {
		if string(Hangman.Displayed_Word[i]) == letter {
			return true
		}
	}
	return false
}

func Allready_Attempted(Hangman *Hangman) bool {
	for i := 0; i < len(Hangman.Letters_Attempted); i++ { //chosen letter or already used
		if Hangman.Try == string(Hangman.Letters_Attempted[i]) {
			return true
		}
	}
	return false
}

func Pendu(Hangman *Hangman) { //affichage des lettres, du nombres essaies restant

	fmt.Println("Write 'stop' if you want to exit the game")
	fmt.Println("Attempted letters :", Hangman.Letters_Attempted)
	fmt.Print("Choose a letter which was not attempted : ")
	fmt.Scan(Hangman.Try)
	fmt.Print("\033[H\033[2J")

	if Hangman.Try == "stop" { //save of the gameplay
		file, _ := json.Marshal(Hangman)
		_ = ioutil.WriteFile("Save.json", file, 0644)
		Hangman.InGame = false
	} else if Hangman.Try == Hangman.Starting_Word { //end of the game, winning game
		printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Starting_Word, Hangman.Letters_Ascii_Art)
		fmt.Println("Congratulation, you win !")
		Hangman.Win = true
	} else if len(Hangman.Try) > 1 { //wrong letter
		Hangman.Attempts--
	} else if len(Hangman.Try) == 1 { //add of a letter
		if Allready_Attempted(Hangman) {
			fmt.Println("The letter was already attempted, please try one which was not.")
			Hangman.Attempts--
		} else {
			Hangman.Letters_Attempted = Hangman.Letters_Attempted + Hangman.Try + " "
		}
	}

	letterInTheWord := false
	AddingLetter := false
	for i := 0; i < len(Hangman.Starting_Word) && !AddingLetter && Hangman.InGame && !Hangman.Win; i++ {
		if Hangman.Try == string(Hangman.Starting_Word[i]) {
			letterInTheWord = true

			if Allready_Displayed(Hangman, Hangman.Try) { //wrong letter, one less attempt
				printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Displayed_Word, Hangman.Letters_Ascii_Art)
				Hangman.Attempts--
				fmt.Println("Oh no ! The letter was allready displayed, you are ", Hangman.Attempts, " attempts again.")
				printhangman.PrintHangman(Hangman.Attempts)
			} else { // correct letter chosen and adding letters in ASCII
				Hangman.Displayed_Word = revealLetters.AddLetter(Hangman.Try, Hangman.Displayed_Word, Hangman.Starting_Word)
				printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Displayed_Word, Hangman.Letters_Ascii_Art)
				printhangman.PrintHangman(Hangman.Attempts)
				AddingLetter = true
			}
		}
	}

	if !letterInTheWord && Hangman.InGame && !Hangman.Win { //show of the numbers of attempts the user have left
		Hangman.Attempts--
		printword.PrintAsciArt(Hangman.Starting_Word, Hangman.Displayed_Word, Hangman.Letters_Ascii_Art)
		fmt.Println("Il te reste ", Hangman.Attempts, " essais.")
		printhangman.PrintHangman(Hangman.Attempts)
	}
}
