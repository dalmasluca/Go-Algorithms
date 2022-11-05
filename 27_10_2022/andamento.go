package main
import "fmt"

func main(){
  var n1,n2,somma int;
  fmt.Scan(&n1);
  for{
    if(n1==-1){break;}
    somma += n1;
    n2 = n1;
    fmt.Scan(&n1);
    if(n1==-1){break;}
    if(n1>n2){
      fmt.Print("+")
    }else{
      fmt.Print("-")
    }
  }
  fmt.Print(somma)
}
