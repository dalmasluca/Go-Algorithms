package main

import (
	"bufio"
	"fmt"
	"os"
)

func stampaContrario( scanner * bufio.Scanner ){
  for scanner.Scan() && scanner.Text()=="stop" {
    parola := (*scanner).Text()
    stampaContrario(scanner)
    fmt.Println(parola)
  }
}


func main(){
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  stampaContrario(scanner)

}
