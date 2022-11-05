package main
import "fmt"

func main(){
  const DIM = 10
  var n, max, ric uint
  
  fmt.Scan(&n)
  max = n
  ric = 1
  for i:=1; i<DIM; i++ {
    fmt.Scan(&n)
    if n > max {
      max = n
      ric = 1
    }else if n == max {
      ric += 1
    }
  }

  fmt.Println("Il max è", max, "ed è comparso", ric,"volte")
}
