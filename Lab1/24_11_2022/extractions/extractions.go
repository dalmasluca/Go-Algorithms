package main

import (
  "fmt"
)

func estraiPari(in []int) (out []int){
  for i := range in {
    if in[i] % 2 == 0 {
      out = append(out, in[i])
    }
  }
  return out
}
func rimuoviMultipli(m int, in []int) (out []int){
  for i := range in {
    if in[i] % m != 0 {
      out = append(out, in[i])
    }
  }
  return out
}

func main(){
  Pari := []int{2,3,4,5,8,10,20,22,38,40}
  fmt.Println(2%2)
  fmt.Println(Pari)
  fmt.Println(estraiPari(Pari))
  fmt.Println(rimuoviMultipli(5,Pari))
}
