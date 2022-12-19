package main
import "fmt"

func main(){
  var n,i int;
  fmt.Print("numero : ");fmt.Scan(&n);
  
  for i=1; (i*i)<=n; i++{
  }
  fmt.Println((i-1)*(i-1))
}
