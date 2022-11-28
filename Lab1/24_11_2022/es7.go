package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

func ordinaSlice(s *[]string) {
	max := len(*s)
	for i := 0; i < max; i++ {
		for j := 0; j < max-1-i; j++ {
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	testo := new([]string)

	for scanner.Scan() && scanner.Text() != "stop" {
		*testo = append(*testo, strings.ToLower(scanner.Text()))
	}

	ordinaSlice(testo)
  fmt.Println(*testo)

}
