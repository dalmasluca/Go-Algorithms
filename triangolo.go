package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < (n*2); i++ {
		for j := 0; j < (n * 2)-1; j++ {
			if (i == j && i < n) || (i == j+1 && i > n) || (i == 0 && j < n) || (i == n*2-1 && j >= n) || (j == n-1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/*
n=2
*
 *

n=4

****
 * *
  **
   *
   **
   * *
   ****
*/
