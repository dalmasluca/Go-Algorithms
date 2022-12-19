package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rettangolo struct {
    base, altezza int
}

func (r *Rettangolo) String() (string){
    s := strings.Repeat(".", (*r).base) + "\n"
    s = strings.Repeat(s, (*r).altezza)
    return s
}

func main () {
    
    if len(os.Args) != 3 {
        os.Exit(1)
    }
    al, _ := strconv.Atoi(os.Args[1])
    base, _ := strconv.Atoi(os.Args[2])
    rettangolo := &Rettangolo{al,base}
    fmt.Print(rettangolo.String())
}
