package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func rigaClessidra(length int, full bool, char byte, shift int) (riga string) {

	if full {
		riga = strings.Repeat(" ", shift)
		riga += strings.Repeat(string(char), length-shift-shift)
		riga += strings.Repeat(" ", shift)
	} else {
		riga = strings.Repeat(" ", shift)
		riga += string(char) + strings.Repeat(" ", length-shift-shift-2)
		riga += string(char) + strings.Repeat(" ", shift)
	}

	return riga
}

func disegnaClessidra(height int, time int) {
	length := height*2 + 1
	shift := 0
	for i := 0; i < height; i++ {
		fmt.Println(rigaClessidra(length, i >= time, '*', shift))
		shift++
	}
	shift--
	for i := height; i > 0; i-- {
		fmt.Println(rigaClessidra(length, time >= i, '*', shift))
		shift--
	}
}

func attendi(n_seconds float64) {
	time.Sleep(time.Duration(n_seconds) * time.Second)
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	var secondi int

	fmt.Print("Inserisci il numero di secondi: ")
	fmt.Scan(&secondi)
	attendi(1)
	cancella()
	for i := 0; i < secondi; i++ {
		disegnaClessidra(secondi, i)
		attendi(1)
		cancella()
	}
}
