package main

import (
	"fmt"
)

func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}

func recursiveMax(list []int) int {
    if len(list) == 1 {
        return list[0]
    }else{
        return max(list[1],recursiveMax(list[1:]))
    }
}

func main() {
    var sint []int
    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        sint = append(sint, n)
    }
    fmt.Println(recursiveMax(sint))
}
