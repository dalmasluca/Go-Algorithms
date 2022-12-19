package main

import (
	"fmt"
	"strings"
)

type Segmento struct {
    estremi, interno byte
    orizzontale bool
    lunghezza int
}

func (s *Segmento) String() (string){
    if (*s).orizzontale {
        return string((*s).estremi) + strings.Repeat(string((*s).interno), (*s).lunghezza) + string((*s).estremi)
    }else{
        return string((*s).estremi) + "\n" + strings.Repeat(string((*s).interno) + "\n", (*s).lunghezza) + string((*s).estremi)
    }
}

func (s *Segmento) allunga(n int) {
    (*s).lunghezza += n
}

func main () {
    segmento := new(Segmento)
    fmt.Scan((*segmento).estremi, (*segmento).interno, (*segmento).orizzontale, (*segmento).lunghezza)
    fmt.Println((*segmento).String())
    (*segmento).allunga(5)
    fmt.Println((*segmento).String())
}
