package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func stampainversa(s []int){
    if len(s) != 0 { 
        stampainversa(s[1:])
        fmt.Print(s[0], " ")
    }
}


func main () {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    var sl []int

    for scanner.Scan() {
        n, _ := strconv.Atoi(scanner.Text())
        sl = append(sl, n)
    }
    
    stampainversa(sl)

}
