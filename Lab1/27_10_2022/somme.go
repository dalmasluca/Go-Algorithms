package main
import "fmt"
import "math/rand"

func main(){
  const MAX = 10;
  var somma int;
  for i:=0; i<MAX; i++{
    somma += rand.Intn(MAX+1);
  }
  fmt.Println(somma);
}
