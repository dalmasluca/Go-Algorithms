package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var VOCALI = "aeiouAEIOU"
	var CIFRE = "0123456789"
	var s1 = "c"
	var s2 = "sc"
	var parola string

	fmt.Printf("Inserisci una stringa : ")
	fmt.Scan(&parola)

	fmt.Println(parola, "contine", s2, ":", strings.Contains(parola, s2))
	fmt.Println(parola, "contine vocali", ":", strings.Contains(parola, VOCALI))
	fmt.Println(parola, "contine occorrenze di", s1, "volte :", strings.Count(parola, s1))

	for pos, runa := range parola {
		if strings.Contains(VOCALI, string(runa)) {
			fmt.Println("La prima vocale si trova in posizione", pos+1)
			break
		}
	}

	for i := len(parola) - 1; i >= 0; i-- {
		if strings.Contains(VOCALI, string(parola[i])) {
			fmt.Println("L'ultima vocale si trova in posizione", i)
			break
		}
	}

	s2 += s2 + s2

	fmt.Println(s2)

	var scifre string

	for _, runa := range parola {
		if strings.Contains(CIFRE, string(runa)) {
			scifre += string(runa)
		}
	}

	n, _ := strconv.Atoi(scifre)
	f, _ := strconv.ParseFloat("0."+scifre, 64)

	fmt.Printf("%d\n", n)
	fmt.Printf("%f\n", f)

}
