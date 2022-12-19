package main

import (
	"fmt"
)

func recursiveSum(list []int) int {
    if len(list) == 0 {
        return 0
    }else{
        return list[0] + recursiveSum(list[1:])
    }
}

func main() {
    var sl []int

    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        sl = append(sl, n)
    }

    fmt.Println(recursiveSum(sl))

}
