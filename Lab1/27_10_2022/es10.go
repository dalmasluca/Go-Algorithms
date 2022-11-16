package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    
    const TARGET = 50
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
