package main
import "fmt"
import "time"
import "math/rand"

func main(){
  rand.Seed(time.Now().Unix());
  fmt.Println(rand.Intn(101));
}
