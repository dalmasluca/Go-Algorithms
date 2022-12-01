package main

import (
	"bufio"
	"fmt"
	"os"
)

func contaVocali(s string, vocali *map[rune]int){
    for _, runa := range s {
        switch runa{
        case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
            (*vocali)[runa]++
        }
    }
}

func stampaMappaOrdinata(vocali map[rune]int) {
    for len(vocali) != 0 {
        min := 'u'
        for k, _ := range vocali{
            if k < min {
                min = k
            }
        }
        fmt.Print(string(min),":",vocali[min],"\n")
        delete(vocali, min)
    }
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    mappaVocali := new(map[rune]int)
    *mappaVocali = make(map[rune]int)

    for scanner.Scan() {
        contaVocali(scanner.Text(), mappaVocali)
    }

   stampaMappaOrdinata(*mappaVocali) 

}
