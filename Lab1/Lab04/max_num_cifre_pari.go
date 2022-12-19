package main
import "fmt"

func main(){
  var s string
  var ric, maxric int

  for !(s=="000") {
    fmt.Scan(&s)
    ric = 0
    for _, bt := range s {
      if bt%2 == 0 {
        ric += 1 
      }
    }
    if ric > maxric { maxric = ric }
  }

  fmt.Println("Massima riccorrenza di cifre pari è", maxric)

}
