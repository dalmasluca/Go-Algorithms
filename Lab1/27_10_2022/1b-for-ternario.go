package main

import "fmt"

func main() {
    /*
    scrivere un programma che utilizzando una costante K  richieda K volte
    un numero e ne faccia la sommatoria
    */
  const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
    fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma = somma + n
	}
	fmt.Println(somma)
}
