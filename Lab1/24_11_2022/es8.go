package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
  var contaParole,sommaLunghezza int
  slice := new([]string)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  for scanner.Scan() {
    sommaLunghezza += strings.Count(scanner.Text(),"")
    contaParole++
    *slice = append(*slice, scanner.Text())
  }
  
  media := sommaLunghezza/contaParole

  for _,word := range *slice {
    if strings.Count(word,"") > media {
      fmt.Printf("%s\n",word)
    }
  }

}
