package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	righePari, righeDispari := []string{}, []string{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	pari := false

	for scanner.Scan() {
		if pari {
			righePari = append(righePari, scanner.Text())
		} else {
			righeDispari = append(righeDispari, scanner.Text())
		}
		pari = !pari
	}

	for i := range righePari {
		fmt.Printf("%s\n",righePari[i])
	}

	for i:= range righeDispari {
		fmt.Printf("%s\n",righeDispari[i])
	}
}
