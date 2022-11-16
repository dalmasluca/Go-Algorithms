package main
import "fmt"

func main(){
  const K = 10;
  var n int;
  fmt.Print("numero : ");
  fmt.Scan(&n);
  for i:=0; i<K; i++{
    fmt.Print(n*(i+1),"\n");
  }
}
