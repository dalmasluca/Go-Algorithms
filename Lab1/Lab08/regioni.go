package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func popolamappa(m *map[string]([]string), scanner *bufio.Scanner){
    sliceofString := strings.Split(scanner.Text(),",")
    
    (*m)[sliceofString[2]] = append((*m)[sliceofString[2]], sliceofString[1])
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    m := make(map[string]([]string))

    for scanner.Scan() {
        popolamappa(&m,scanner)
    }
    
    for i:=1; i<len(os.Args); i++ {
        fmt.Println(m[os.Args[i]])
    }
}
