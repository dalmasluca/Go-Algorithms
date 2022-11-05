package main
import "fmt"

func main(){
  var n int;
  fmt.Print("inserisci numero : ");
  fmt.Scan(&n);
  for i:=2; i<=n/2; i++ {
    if(n%i==0){ fmt.Print("non "); break;}
  }
  fmt.Print("é primo");
}
