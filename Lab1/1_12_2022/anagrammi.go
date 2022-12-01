package main

import (
	"fmt"
	"os"
)

func makeMapFromString(p string) *map[rune]int {
    m := make(map[rune]int)
    for _, runa := range p {
        m[runa] ++
    }
    return &m
}

func isAnagram(p1 string, p2 string) bool{
    m1 := makeMapFromString(p1)
    m2 := makeMapFromString(p2)
    
    if len(*m1) < len(*m2) {
        m1, m2 = m2, m1
    }

    for k, _ := range *m1 {
        if (*m1)[k] != (*m2)[k] {
            return false
        }
    }
    return true
}

func main(){
    if len(os.Args) != 3 {
        fmt.Println("input errato")
        os.Exit(1)
    }
    fmt.Print(os.Args[1], " e ", os.Args[2])
    if !isAnagram(os.Args[1],os.Args[2]) {
        fmt.Print(" non")
    }
    fmt.Print(" sono anagrammi")
}
