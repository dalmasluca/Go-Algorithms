package main

import "fmt"

func main(){
  var contatore,numero int
  var found bool
  found = false
  fmt.Scan(&numero)
  for numero > 0{
    cifra := numero%10
    contatore++
    if cifra%2==0 {
      fmt.Print("La prima cifra si trova in posizione ",contatore)
      found = true
    }
  numero = numero/10
  }
  if !found{
    fmt.Print("-1")
  }
}
