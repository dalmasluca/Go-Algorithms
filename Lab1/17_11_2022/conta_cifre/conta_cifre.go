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

  fmt.Scan(&stringaInput)

  for strings.Compare(stringaInput,"stop")!=0{
    if contaCifre(stringaInput,numCifre) {
      contaStringhe+=1
    }
  }
  fmt.Scan(stringaInput)
  fmt.Println(contaStringhe,"stringhe con cifre",numCifre)
}
