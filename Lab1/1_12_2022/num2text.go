package main

import (
    "fmt"
    "strconv"
)

func main() {
    m := make(map[int]string)
    var numString string
    var parola string
    fmt.Print("inserirei un numero: ")
    fmt.Scan(&numString)

    for _, runa := range numString {
        num, _ := strconv.Atoi(string(runa))
        if m[num] == "" {
            fmt.Printf("parola per %d ? ", num)
            fmt.Scan(&parola)
            m[num] = parola
        }
    }
    for i, runa := range numString {
        num, _ := strconv.Atoi(string(runa)) 
        fmt.Printf("%s", m[num])
        if i != len(numString)-1 {
            fmt.Printf(" - ")
        }
    }
    fmt.Println()
}
