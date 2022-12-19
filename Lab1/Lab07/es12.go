package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() || scanner.Text()=="stop" {
		for _, runa := range scanner.Text() {
			if !strings.ContainsAny(string(runa), "?!,.:;") {
				fmt.Print(string(runa))
			}
		}
	}

}
