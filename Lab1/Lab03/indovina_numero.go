package main
import ("fmt";"math/rand";"time")

func main(){
  rand.Seed(time.Now().Unix())
  const MAX = 10;
  var n int;
  fmt.Print("Inserisci numero : ");
  fmt.Scan(&n);
  contTentativi := 1;
  numeroCasuale :=  rand.Intn(MAX+1);
  for n!=numeroCasuale && contTentativi<=(MAX/2){
    if n>MAX || n<0{
      fmt.Println("fuori intervallo!");
      contTentativi-=1;
    }
    fmt.Scan(&n);
    contTentativi++;
  }
  if (n==numeroCasuale){
    fmt.Println("Hai vinto!");
  }else{
    fmt.Println("Hai perso!, il numero ",numeroCasuale);
  }
}
