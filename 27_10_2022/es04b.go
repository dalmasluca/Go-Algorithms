package main
import "fmt"
import "math/rand"
import "time"

func main(){
  const K = 10;
  var contapari,contadispari int;
  rand.Seed(time.Now().Unix());
  for i:=0; i<K; i++{
    if(rand.Intn(11)%2 == 0){
      contapari++;
    }else{
      contadispari++;
    }
  }
  fmt.Println("pari : ",contapari);
  fmt.Println("disapri : ",contadispari);
}
