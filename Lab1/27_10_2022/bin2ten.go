package main

import (
	"fmt"
	"math"
)

func main() {
	var binario, decimale, contatore int
	var errore bool
	errore = false
	fmt.Print("Inserisci il binario : ")
	fmt.Scan(&binario)
	for binario > 0 {
		cifraBinario := binario % 10
		if cifraBinario != 0 && cifraBinario != 1 {
			fmt.Print("Errore inserito un numero sbagliato")
			errore = true
			break
		}
		decimale += cifraBinario * int(math.Pow(2, float64(contatore)))
		contatore++
		binario = binario / 10
	}
	if !errore {
		fmt.Println(decimale)
	}
}
