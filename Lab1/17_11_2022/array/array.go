package main

import "fmt"

const DIM = 5

func reverse(a *[DIM]int) {
	for i := 0; i < DIM/2; i++ {
		a[i], a[DIM-i-1] = a[DIM-i-1], a[i]
	}
}

func switchFirstLast(a *[DIM]int) {
	a[0], a[DIM-1] = a[DIM-1], a[0]
}

func main() {
	var a [DIM]int

	for pos := range a {
		a[pos] = pos
	}

	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
  switchFirstLast(&a)
  fmt.Println(a)

}
