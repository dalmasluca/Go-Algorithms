package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
)

func righello(n int) {
    if n == 0 {
        return
    }else{
        righello(n-1)
        fmt.Println(strings.Repeat("-", n))
        righello(n-1)
    }
}

func main(){
    n, _ := strconv.Atoi(os.Args[1])
    righello(n)
}
