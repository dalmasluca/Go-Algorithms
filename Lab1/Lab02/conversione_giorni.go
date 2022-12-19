package main
import "fmt"

func main(){
  const anno int = 365;
  const settimana int = 7;
  var giorni,anni,settimane int;
  fmt.Print("numero di giorni da convertire?");
  fmt.Scan(&giorni);
  anni = giorni/anno;
  settimane = (giorni%anno)/settimana;
  fmt.Println(giorni,"giorni =", anni,"anni",settimane,"giorni",(giorni%anno)%settimana);
}
