package main

import "fmt"

func main() {
	var prec, curr, c int
	var seq bool
  fmt.Scan(curr)
	for {
		if curr == 2 {
			break
		}
		prec = curr
		fmt.Scan(&curr)
		if prec == curr && !seq {
			c++
      seq = true
		}else{
      seq = false
    }
	}
  fmt.Println(c);
}
