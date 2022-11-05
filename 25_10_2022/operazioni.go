package main
import "fmt"

func main(){
    var numero1,numero2 float64;
    fmt.Print("numero1 = ");
    fmt.Scan(&numero1);
    fmt.Print("numero2 = ");
    fmt.Scan(&numero2);
    fmt.Println("Somma =",numero1+numero2);
    fmt.Println("Sottrazione =",numero1-numero2);
    fmt.Println("Prodotto =",numero1*numero2);
    fmt.Println("Quoziente =",numero1/numero2);
}
