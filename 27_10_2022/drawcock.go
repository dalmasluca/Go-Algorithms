package main
import "fmt"

func main(){
  var n int;
  fmt.Print("Inserisci in cm : ");
  fmt.Scan(&n);
  fmt.Print("8");
  for i:=0; i<n; i++{
    fmt.Print("=");
  }
  fmt.Print("D");
}
