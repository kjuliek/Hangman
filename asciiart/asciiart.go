package asciiart

func TriMot(word string) string {
	var ref []rune
	var newword string
	ref = append(ref, '_')
	for r := 97; r <= 122; r++ {
		ref = append(ref, rune(r))
	}
	for _, letter := range word {
		for _, l := range ref {
			if letter == l {
				newword += string(letter)
			}
		}
	}
	return newword
}

func AsciiArt(mot string) {
	newword := TriMot(mot)

}
