package main

import (
	"fmt"
	"math/rand"
  "time"
)

func main(){
  rand.Seed(time.Now().Unix())
  var max int;
  const MAX = 30;
  const MIN = 20;
  const CICLI = 10
  for i:=0; i<CICLI; i++{
    n := rand.Intn(MAX-MIN+1)+MIN
    if n>max{
      max = n;
    }
    fmt.Println(n);
  }
  fmt.Println(max)
}
