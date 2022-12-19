package main
import "fmt"
func main() {
    /*
    stampare il conto alla rovescia a partire da un numero letto da standard
        input
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
