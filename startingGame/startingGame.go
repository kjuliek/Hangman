package startingGame

import (
	"fmt"
	"os"
)

func start() {
	fmt.Print("\033[H\033[2J")
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments.")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("Text file is missing.")
		return
	}

}
