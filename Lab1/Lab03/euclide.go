package main
import "fmt"

func main(){
  var a,b,r int;
  
  fmt.Print("Inserirei due numeri : ");
  fmt.Scan(&a,&b);
  
  if(a<b){
    a,b = b,a
  }

  for b!=0 {
      r = a%b;
      a,b = b,r;
  }
  fmt.Print("MCD : ",b);
}
