package main
import "fmt"

func main(){
  var distanza,litri float64;
  fmt.Print("distanza percorsa (in km): ");
  fmt.Scan(&distanza);
  fmt.Print("quantita di carburante utilizzata (in l):");
  fmt.Scan(&litri);
  fmt.Println("consumo medio:",litri/distanza,"l/km");
  fmt.Println("resa media:",distanza/litri,"km/l");
}
