package printmot

import "fmt"

func Printmot(mot string) {
	for i := 0; i < len(mot); i++ {
		fmt.Print(string(mot[i]), " ")
	}
	fmt.Println()
}

func AsciiArt(mot string, lettres_ascii_art [][]string) {

}
