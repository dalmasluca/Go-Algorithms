package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*
    creare un prgrammi che generi numeri casuali da 0 a 21 e lo stampi
    fino a quando non viene geenerato il numero 0, stampi a video il conteggio
    dei numeri generati prima dell'interruzione
    */
    rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    contaNumGenerati := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        contaNumGenerati++
    }
    fmt.Println(contaNumGenerati)
}
