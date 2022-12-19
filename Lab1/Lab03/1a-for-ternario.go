package main
import "fmt"
func main() {
    /*
      creare un prgramma che stampi n * con n preso da standard input
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
