package main
import "fmt"
import "math/rand"
import "time"

func main(){
  const K = 10;
  rand.Seed(time.Now().Unix());
  for i:=0; i<K; i++{
    fmt.Println(rand.Intn(11));
  }
}
