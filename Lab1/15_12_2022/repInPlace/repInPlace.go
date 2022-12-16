package main

import (
	"os"
)

func repInPlace(stringa []rune, old rune, newr rune) {
	for i := 0; i < len(stringa); i++ {
        if stringa[i] == old {
            stringa[i] = newr
        }
	}
}

func main() {
    if len(os.Args) != 4 {
        os.Exit(1)
    }
    repInPlace([]rune(os.Args[1]), rune(os.Args[2][0]), rune(os.Args[3][0]))
}
