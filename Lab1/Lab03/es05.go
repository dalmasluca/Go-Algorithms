package main
import "fmt"

func main(){
  var n int;
  fmt.Print("numero : ");
  fmt.Scan(&n);
  for i:=0; i<n; i++{
    fmt.Println(i+1);
  }
}
