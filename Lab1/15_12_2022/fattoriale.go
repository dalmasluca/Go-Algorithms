package main

import (
	"fmt"
	"os"
	"strconv"
)

func fattoriale(n int) int {
    if n == 0 {
        return 1
    }else{
        return n * fattoriale(n-1)
    }
}

func main() {
    n, _ := strconv.Atoi(os.Args[1])
    fmt.Println(fattoriale(n))
}
