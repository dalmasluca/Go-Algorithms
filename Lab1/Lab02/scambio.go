package main
import "fmt"

func main(){
  var numero1,numero2 int;
  fmt.Print("val1 e val2 (int): ");
  numero,errore := fmt.Scan(&numero1,&numero2);
  _ = numero
  fmt.Println("errore: ", errore);
  fmt.Println("prima dello scmabio: ",numero1,numero2);
  numero1,numero2 = numero2,numero1;
  fmt.Println("Dopo lo scmabio: ",numero1,numero2);
}
