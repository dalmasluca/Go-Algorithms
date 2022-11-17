package main

import "fmt"

func main() {
	var mm, c, cp int
	for {
		
    n, err := fmt.Scan(&mm)
    c++

    if (mm!=0 && n!=0){ cp = c }

    if err != nil {
			break
		}
	}
  fmt.Println(cp)
}
