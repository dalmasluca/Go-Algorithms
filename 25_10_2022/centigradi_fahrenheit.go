package main
import "fmt"

func main(){
  var centigradi int;
  var fahrenheit float64;
  fmt.Print("temperatura in gradid centigradi? ");
  fmt.Scan(&centigradi);
  fahrenheit = float64( centigradi );
  fahrenheit = ((9./5.) * fahrenheit) + 32;
  fmt.Println("18 C =", fahrenheit,"F");
}
