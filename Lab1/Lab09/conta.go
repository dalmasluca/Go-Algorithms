package main

import "fmt"

func recursiveCount(el int, list []int) int{
    if len(list) == 1 {
        if el == list[0] {
            return 1
        }else{
            return 0
        }
    }
    if el == list[0]{
        return 1+ recursiveCount(el, list[1:])
    }else{
        return recursiveCount(el, list[1:])
    }
}

func main(){
    var l []int
    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        l = append(l, n)
    }
    fmt.Println(recursiveCount(l[0], l[1:]))
}
