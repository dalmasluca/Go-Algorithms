package main
import "fmt"

func main(){
  var ca,in,ob float64;
  var cAnnuo int;
  fmt.Printf("Inserisci capitale iniziale, interesse annuo e obiettivo : ");
  fmt.Scan(&ca,&in,&ob)
  for ca<=ob {
    ca += (ca*in)/100;
    cAnnuo++;
  }
  fmt.Print(cAnnuo);
}
