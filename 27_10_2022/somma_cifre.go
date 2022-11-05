package main
import "fmt"

func main(){
  var n,somma int;
  fmt.Print("Inserisci un numero : ")
  fmt.Scan(&n);
  for n>0 {
    somma += n%10
    n = n/10
  }
  fmt.Print(somma)
}
