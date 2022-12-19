package main

import (
	"bufio"
	"fmt"
	"os"
)

func aggiornaConteggio(m *map[int]int, riga string){
    (*m)[len(riga)] ++ 
}

func main(){
    mappaParole := make(map[int]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        aggiornaConteggio(&mappaParole, scanner.Text())
    }

    fmt.Println(mappaParole)
}
