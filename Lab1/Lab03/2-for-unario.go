package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
     crei un programma che generi numeri casuali da 1 a 10 fino a quando
     la somma dei numeri generati è inferiore ad un target
    */
    const TARGET = 20
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    sommaNumeriGenerati:= 0
    for sommaNumeriGenerati < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        sommaNumeriGenerati += n
    }
    fmt.Println("\n",sommaNumeriGenerati)
}
