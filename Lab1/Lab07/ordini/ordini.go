package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totale float64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '#' || data[i] == '\n'{
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	})

	for scanner.Scan() {
		prezzo, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		quantita, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		sconto, _ := strconv.ParseFloat(scanner.Text(), 64)
		
		totale += prezzo * ( 1 - sconto) * quantita

	}

	fmt.Println(totale)
}
