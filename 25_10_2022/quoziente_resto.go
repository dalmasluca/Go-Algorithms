package main
import "fmt"

func main(){
  var divisore,dividendo int;
  fmt.Print("dividendo: ");
  fmt.Scan(&dividendo);
  fmt.Print("divisore: ");
  fmt.Scan(&divisore);
  fmt.Println("quoziente = ",dividendo/divisore);
  fmt.Println("resto = ",dividendo%divisore);
}
