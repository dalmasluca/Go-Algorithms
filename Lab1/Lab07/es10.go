package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var maxLen int
	slice := new([]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if strings.Count(scanner.Text(), "") > maxLen {
			maxLen = strings.Count(scanner.Text(), "")
		}
		*slice = append(*slice, scanner.Text())
	}

	for _, word := range *slice {
		if strings.Count(word,"") == maxLen {
			fmt.Printf("%s\n", word)
		}
	}

}
