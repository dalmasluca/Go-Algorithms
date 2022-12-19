package main

import (
	"fmt"
	"strings"
)

func contaCifre(s string, numCifre *[10] int ) (haCifre bool){
  haCifre = false
  for _,runa := range s{
    if strings.ContainsAny("0123456789",string(runa)){
      haCifre = true
      (*numCifre)[int(runa - '0')]++
    }
  }
  return haCifre
}

func main(){
  var stringaInput string
  var numCifre *[10] int
  var contaStringhe int
  
  numCifre = new([10]int)

  fmt.Scan(&stringaInput)

  for stringaInput!="stop"{
    if contaCifre(stringaInput,numCifre) {
      contaStringhe+=1
    }
  fmt.Scan(&stringaInput)
  }
  fmt.Println(contaStringhe,"stringhe con cifre\n[0 1 2 3 4 5 6 7 8 9]")
  fmt.Println(*numCifre)
}
