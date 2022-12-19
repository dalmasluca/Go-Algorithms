package main
import "fmt"
func main() {
    /*
      stampare i primi numeri pari inferiori ad n letto da standard input
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
