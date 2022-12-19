package main

import (
	"fmt"
	"math"
)

func main(){
  const TARGET = 10;
  const CICLI = 5;

  var n,vicino int;

  for i:=0; i<CICLI; i++{
    fmt.Scan(&n);
    if math.Abs(float64(TARGET-n)) < math.Abs(float64(TARGET-vicino)){
      vicino = n;
    }
  }
  fmt.Println(vicino)
}
