package main

import (
	"fmt"
	"os"
	"strconv"
)

func isSquare(n int) bool {
	for i:=0; i<= 1 + n/2; i++ {
		if i*i == n{
			return true
		}
	}
	return false
}

func main(){
	for i:= 1; i<len(os.Args); i++{
		number, _ := strconv.Atoi(os.Args[i])
		if isSquare(number){
			fmt.Println(os.Args[i])
		}
	}
}
