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

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    mappaVocali := new(map[rune]int)
    *mappaVocali = make(map[rune]int)

    for scanner.Scan() {
        contaVocali(scanner.Text(), mappaVocali)
    }
    for k, v := range *mappaVocali {
        fmt.Print(string(k),": ",v,"\n")
    }
}
