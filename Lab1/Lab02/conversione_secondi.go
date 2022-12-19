package main
import "fmt"

func main(){
  var secondi int;
  const giorno,ora,minuto int = 86400,3600,60;
  fmt.Print("numero di secondi: ");
  fmt.Scan(&secondi);
  fmt.Print("g:h:m:s = ",(secondi/giorno),":",(secondi%giorno)/ora,":", ((secondi%giorno)%ora)/minuto,":",((secondi%giorno)%ora)%minuto ,"\n");

}
