 package main
 import "fmt"

 func main(){
  var n1,n2 float64;
  fmt.Print("Inserisci due numero : ")
  fmt.Scan(&n1);
  fmt.Scan(&n2);
  for n2!=0{
    fmt.Println(n2-n1);
    n1 = n2;
    fmt.Scan(&n2); 
  }
 }
