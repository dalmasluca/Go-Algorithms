package main
import "fmt"

func main(){
  var n int;
  
  fmt.Print("numero : ");
  fmt.Scan(&n);

  for i:=1; i<=n; i++{
    
    for j:=1; j<=i; j++ {
      if(j*j == i){
        fmt.Println(i);
      }
    }
  
  }
}
