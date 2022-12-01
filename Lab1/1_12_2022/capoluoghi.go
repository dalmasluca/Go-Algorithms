package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Capoluogo struct {
    Nome, Sigla, Regione  string
    Popolazione, Superficie, Densita, Altitudine int
}

func popolamappa(m *map[string]Capoluogo, scanner *bufio.Scanner){
    sliceofString := strings.Split(scanner.Text(),",")
    capoluogo := new(Capoluogo)

    (*capoluogo).Nome = sliceofString[0]
    (*capoluogo).Sigla = sliceofString[1]
    (*capoluogo).Regione = sliceofString[2]
    num, _ := strconv.Atoi(sliceofString[3])
    (*capoluogo).Popolazione = num
    num, _ = strconv.Atoi(sliceofString[4])
    (*capoluogo).Superficie = num 
    num, _ = strconv.Atoi(sliceofString[5])
    (*capoluogo).Densita = num
    num, _ = strconv.Atoi(sliceofString[6])
    (*capoluogo).Altitudine = num
   
    (*m)[(*capoluogo).Sigla] = *capoluogo

}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    m := make(map[string]Capoluogo)

    for scanner.Scan() {
        popolamappa(&m,scanner)
    }
    
    for i:=1; i<len(os.Args); i++ {
        fmt.Println(m[os.Args[i]])
    }
}
