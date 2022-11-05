package main

import "fmt"

func main() {
	var prec, curr, c, maxc int
	var seq, gradino bool
	_ = gradino

	fmt.Scan(curr)

	for {

		prec = curr
		fmt.Scan(&curr)
		if curr < prec {
			break
		}



	}
	fmt.Println(c)
}
