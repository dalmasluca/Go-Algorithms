package main

import (
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
    if len(s) == 1 || len(s) == 0 {
        return true
    }
    if s[0] == s[len(s)-1] {
        return isPalindrome(s[1:len(s)-2])
    }else{
        return false
    }
}

func main(){
    if isPalindrome(os.Args[1]) {
        fmt.Println(os.Args[1] + " è palindroma")
    }else{
         fmt.Println(os.Args[1] + " non è palindroma")
     }
         
}
