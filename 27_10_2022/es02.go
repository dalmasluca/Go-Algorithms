package main
import "fmt"

func main(){
  const K = 5;
  var n int;

  for i:=0; i<K; i++{
    fmt.Print("Inserisci un nunmero:\t");
    fmt.Scan(&n);
    fmt.Print(2*n,"\n");
  }
}
