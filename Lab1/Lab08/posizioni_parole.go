package main

import (
	"bufio"
	"fmt"
    "os"
)

func main(){
    m := make(map[string][]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    var c int

    for scanner.Scan() {
        m[scanner.Text()] = append(m[scanner.Text()], c)
        c++
    }

    fmt.Println(m)
}
